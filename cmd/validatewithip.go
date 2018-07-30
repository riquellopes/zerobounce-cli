package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var validatewithip = &cobra.Command{
	Use:   "validatewithip [string email][string ipaddress]",
	Short: "Get the validation result for one email address, using an IP",
	Long: ` Get information about the email address, such as
        bounce status, First Name, Last Name, using an IP Address
        for live validations.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("bye called")
	},
}

func init() {
	RootCmd.AddCommand(validatewithip)

}
