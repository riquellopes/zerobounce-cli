package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// RootCmd -
var RootCmd = &cobra.Command{
	Use:   "zerobouncecli",
	Short: "cli version of zerobounce.",
	Long:  `cli version of zerobounce for quick reference.`,
}

// Execute -
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
