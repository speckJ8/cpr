package utils

import (
	"os"

	git "github.com/libgit2/git2go/v30"
	"github.com/speckJ8/cpr/licenses"
)

// Detect the license used in the project by looking
// at LICENSE or COPYING files
func DetectLicense() *licenses.License {
    licenseContents, err := os.ReadFile("LICENSE")
    if err != nil {
        licenseContents, err = os.ReadFile("COPYING")
        if err != nil {
            return nil
        }
    }

    var license = string(licenseContents)
    if l := licenses.IsMIT(license); l != nil {
        return l
    }
    if l := licenses.IsGPL3(license); l != nil {
        return l
    }

    return nil
}

// Detect the author by looking at git's config
func DetectAuthor() string {
    config, err := git.OpenDefault()
    if err != nil {
        return ""
    }

    author, err := config.LookupString("user.name")
    if err != nil {
        return ""
    }

    return author
}

// Detect the email by looking at git's config
func DetectEmail() string {
    config, err := git.OpenDefault()
    if err != nil {
        return ""
    }

    email, err := config.LookupString("user.email")
    if err != nil {
        return ""
    }

    return email
}
