
package cmd

import (
	"fmt"
	"agenda/service"
	"github.com/spf13/cobra"
)



// 取消会议命令
var cancell_meetingCmd = &cobra.Command{
	Use:   "cancell_meeting -t [meeting title]",
	Short: "cancell meeting through its title",
	Long: `Input command model like: cancell_meeting -t BeautifulGo`,
	Run: func(cmd *cobra.Command, args []string) {
		tmp_t, _ := cmd.Flags().GetString("title")
		if service.GetFlag() == true {
			service.Cancell_meeting(tmp_t)
		} else {
			fmt.Println("You have not sign in!")
		}
		
	},
}

func init() {
	RootCmd.AddCommand(cancell_meetingCmd)
	cancell_meetingCmd.Flags().StringP("title", "t", "", "meeting title")

}
