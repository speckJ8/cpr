package cmd

import (
        "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command {
    Use: "cpr",
    Short: "Manage copyright notices in source files",
    Long: `Manage copyright notices in source files. cpr allows you to
create files with with copyright banners, update copyright dates
on existing files and update authors in existing files.`,
}

var (
    license string
    author string
    email string
    short bool
    all bool
    extensions []string
)

func Execute() error {
    rootCmd.AddCommand(newCmd)

    return rootCmd.Execute()
}
