package cmd

import (
	"fmt"
	"os"

	Z "github.com/riquellopes/zerobounce-go-api/zerobounce"
	"github.com/spf13/cobra"
)

// RootCmd -
var RootCmd = &cobra.Command{
	Use:   "zerobouncecli",
	Short: "cli version of zerobounce.",
	Long:  `cli version of zerobounce for quick reference.`,
}

// Zero -
var Zero = &Z.ZeroBounce{Apikey: os.Getenv("API_KEY_ZERO")}

// Execute -
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
