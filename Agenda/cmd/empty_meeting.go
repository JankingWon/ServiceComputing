
package cmd

import (
	"fmt"
	"agenda/service"
	"github.com/spf13/cobra"
)

// 清空所有会议
var empty_meetingCmd = &cobra.Command{
	Use:   "empty_meeting",
	Short: "empty all meeting you call",
	Long: `Input command model like : empty_meeting`,
	Run: func(cmd *cobra.Command, args []string) {

		if service.GetFlag() == true {
			service.Empty_meeting()
		} else {
			fmt.Println("You have not sign in!")
		}
		
	},
}

func init() {
	RootCmd.AddCommand(empty_meetingCmd)


}
