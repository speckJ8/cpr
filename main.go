package main

import (
        "os"
        "github.com/speckJ8/cpr/cmd"
)

func main() {
    if err := cmd.Execute(); err != nil {
        os.Exit(1)
    }
}
