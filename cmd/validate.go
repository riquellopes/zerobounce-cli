package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var validate = &cobra.Command{
	Use:   "validate [string email]",
	Short: "Get the validation result for one email address.",
	Long: `Get the validation result for one email address
        Get information about the email address, such as
        bounce status, First Name, Last Name`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		validate, _ := Zero.Validate(args[0])

		response, _ := json.MarshalIndent(validate, "", " ")

		fmt.Println(string(response))
	},
}

func init() {
	RootCmd.AddCommand(validate)

}
