package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var validatewithip = &cobra.Command{
	Use:   "validatewithip [string email][string ipaddress]",
	Short: "Get the validation result for one email address, using an IP",
	Long: ` Get information about the email address, such as
        bounce status, First Name, Last Name, using an IP Address
        for live validations.`,
	Args: cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
		ip := ""
		if len(args) == 2 {
			ip = args[1]
		}

		validate, _ := Zero.ValidateWithip(args[0], ip)

		response, _ := json.MarshalIndent(validate, "", " ")

		fmt.Println(string(response))
	},
}

func init() {
	RootCmd.AddCommand(validatewithip)

}
