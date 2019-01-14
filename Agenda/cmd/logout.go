package cmd

import (
	"fmt"
	"agenda/service"
	"github.com/spf13/cobra"
)

// log_outCmd represents the log_out command
var log_outCmd = &cobra.Command{
	Use:   "logout",
	Short: "log out",
	Long: `Input command mode like : logout`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		if service.GetFlag() == true {
			service.Log_out()
		} else {
			fmt.Println("You don't log in!")
		}
		
	},
}

func init() {
	RootCmd.AddCommand(log_outCmd)
}
