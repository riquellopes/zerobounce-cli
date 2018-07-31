package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var getcredits = &cobra.Command{
	Use:   "getcredits",
	Short: "Get the number of credits available in your ZeroBounce account",
	Long:  `Get the number of credits available in your ZeroBounce account`,
	Run: func(cmd *cobra.Command, args []string) {
		credit, _ := Zero.GetCredits()

		fmt.Println(credit)
	},
}

func init() {
	RootCmd.AddCommand(getcredits)

}
