package cmd

import (
	"agenda/service"
	"github.com/spf13/cobra"
)



// log_inCmd represents the log_in command
var log_inCmd = &cobra.Command{
	Use:   "login -n [username] -p [password]",
	Short: "log in",
	Long: `Input command mode like : login -n user -p 123456`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		tmp_n, _ := cmd.Flags().GetString("name")
		tmp_p, _ := cmd.Flags().GetString("password")
		service.Log_in(tmp_n, tmp_p)
	},
}

func init() {
	RootCmd.AddCommand(log_inCmd)
	log_inCmd.Flags().StringP("name", "n", "", "user name")
	log_inCmd.Flags().StringP("password", "p", "", "user password")

}
