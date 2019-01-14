package cmd

import (
	"fmt"
	"agenda/service"
	"github.com/spf13/cobra"
)


var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete your account",
	Long: `Input command model like: delete`,
	Run: func(cmd *cobra.Command, args []string) {
		if service.GetFlag() == true {
			service.Delete_user()
		} else {
			fmt.Println("You don't log in!")
		}
		//
		
	},
}

func init() {
	RootCmd.AddCommand(deleteCmd)
}
