package cmd

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"strings"
	"text/template"
	"time"

	"github.com/speckJ8/cpr/langs"
	"github.com/speckJ8/cpr/licenses"
	"github.com/speckJ8/cpr/utils"
	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command {
    Use: "new",
    Aliases: []string{ "n" },
    Short: "Create a new file with a copyright notice.",
    Long: "Create a new file with a copyright notice.",
    Args: cobra.MinimumNArgs(1),
    Run: newFile,
}

func init() {
    newCmd.Flags().StringVarP(&author, "author", "a", "", "Author name")
    newCmd.Flags().StringVarP(&email, "email", "e", "", "Author email")
    newCmd.Flags().StringVarP(&license, "license", "l", "", "The license to use")
    newCmd.Flags().BoolVarP(&short, "short", "s", false, "Use a short license notice")
}

func newFile(cmd *cobra.Command, args []string) {
    if author == "" {
        author = utils.DetectAuthor()
        if author == "" {
            fmt.Println("author name not specified")
            os.Exit(1)
        }
    }

    if email == "" {
        email = utils.DetectEmail()
        if email == "" {
            fmt.Println("email not specified")
            os.Exit(1)
        }
    }

    var li *licenses.License
    if license == "" {
        li = utils.DetectLicense()
        if li == nil {
            fmt.Println("license not specified")
            os.Exit(1)
        }
    } else {
        li = licenses.Licenses[license]
        if li == nil {
            fmt.Printf("unknown license `%s`\n", license)
            os.Exit(1)
        }
    }

    var licenseNotice string
    if short {
        licenseNotice = li.Simple
    } else {
        licenseNotice = li.Verbose
    }
    tmp, err := template.New("notice").Parse(licenseNotice)
    if err != nil {
        fmt.Printf("failed to load license `%s`: %s\n", license, err.Error())
        os.Exit(1)
    }

    var notice bytes.Buffer
    var templateArgs = licenses.LicenseNoticeTemplateArgs {
        CurrentYear: fmt.Sprintf("%d", time.Now().UTC().Year()),
        Author: author,
        Email: email,
    }
    tmp.Execute(&notice, templateArgs)

    for _, file := range args {
        if _, err := os.Stat(file); errors.Is(err, os.ErrNotExist) {
            continue
        } else {
            fmt.Printf("%s already exists\n", file)
            os.Exit(1)
        }
    }

    var noticeString = notice.String()
    for _, file := range args {
        var nameParts = strings.Split(file, ".")
        var extension = nameParts[len(nameParts) - 1]
        lang, ok := langs.Langs[extension]
        if !ok {
            fmt.Printf("unknown language extension .%s\n", extension)
            os.Exit(1)
        }

        fil, err := os.Create(file)
        defer fil.Close()
        if err != nil {
            fmt.Println(err.Error())
            os.Exit(1)
        }

        _, err = fil.Write([]byte(lang.WrapInComments(noticeString)))
        if err != nil {
            fmt.Printf("failed to write to %s: %s\n", file, err.Error())
            os.Exit(1)
        }
    }
}
