
package cmd

import (
	"fmt"
	"agenda/service"
	"github.com/spf13/cobra"
)


// 退出会议命令
var exit_meetingCmd = &cobra.Command{
	Use:   "exit_meeting -t [meeting title]",
	Short: "exit a meeting, if the paticipator of the meeting is none, we will delete the meeting",
	Long: `Input command mode like : exit_meeting -t Golang`,
	Run: func(cmd *cobra.Command, args []string) {

		tmp_t, _ := cmd.Flags().GetString("title")
		if service.GetFlag() == true {
			service.Exit_meeting(tmp_t)
		} else {
			fmt.Println("You have not sign in!")
		}
		
	},
}

func init() {
	RootCmd.AddCommand(exit_meetingCmd)

	exit_meetingCmd.Flags().StringP("title", "t", "", "meeting title")

}
