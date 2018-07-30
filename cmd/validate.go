package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var validate = &cobra.Command{
	Use:   "validate [string email]",
	Short: "Get the validation result for one email address.",
	Long: `Get the validation result for one email address
        Get information about the email address, such as
        bounce status, First Name, Last Name`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("bye called")
	},
}

func init() {
	RootCmd.AddCommand(validate)

}
