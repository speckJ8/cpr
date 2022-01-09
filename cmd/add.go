package cmd

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"text/template"
	"time"

	"github.com/speckJ8/cpr/langs"
	"github.com/speckJ8/cpr/licenses"
	"github.com/speckJ8/cpr/utils"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command {
    Use: "add",
    Aliases: []string{ "a" },
    Short: "Add a copyright notice to existing files.",
    Long: "Add a copyright notice to existing files.",
    Args: cobra.MinimumNArgs(1),
    Run: addToFile,
}

func init() {
    addCmd.Flags().StringVarP(&author, "author", "u", "", "Author name")
    addCmd.Flags().StringVarP(&email, "email", "e", "", "Author email")
    addCmd.Flags().StringVarP(&license, "license", "l", "", "The license to use")
    addCmd.Flags().BoolVarP(&short, "short", "s", false, "Use a short license notice")
    addCmd.PersistentFlags().BoolVarP(&all, "all", "a", false,
        "Add the notice to all the files in the project.")
    addCmd.PersistentFlags().StringSliceVarP(&extensions, "extensions", "x", []string{},
        "Add the notice only to files with these extensions.")
}

func addToFile(cmd *cobra.Command, args []string) {
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

    var noticeString = notice.String()
    for _, file := range args {
        var nameParts = strings.Split(file, ".")
        var extension = nameParts[len(nameParts) - 1]
        lang, ok := langs.Langs[extension]
        if !ok {
            fmt.Printf("unknown language extension .%s\n", extension)
            os.Exit(1)
        }

        sourceFile, err := os.OpenFile(file, os.O_CREATE|os.O_RDWR, os.ModePerm)
        if err != nil {
            fmt.Println(err.Error())
            os.Exit(1)
        }
        defer sourceFile.Close()

        contents, err := io.ReadAll(sourceFile)
        if err != nil {
            fmt.Println(err.Error())
            os.Exit(1)
        }

        contents = append([]byte(lang.WrapInComments(noticeString)), contents...)
        _, err = sourceFile.Write(contents)
        if err != nil {
            fmt.Printf("failed to write to %s: %s\n", file, err.Error())
            os.Exit(1)
        }
        sourceFile.Close()
    }
}
