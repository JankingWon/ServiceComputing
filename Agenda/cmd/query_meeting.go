
package cmd

import (
	"fmt"
	"agenda/service"
	"github.com/spf13/cobra"
)


// 查询会议命令
var query_meetingCmd = &cobra.Command{
	Use:   "query_meeting -s [start time] -e [end time]",
	Short: "query meeting",
	Long: `Input command model like : query_meeting -s 2018-10-01/12:00 -e 2018-11-01/12:00

	then we will query the meeting which is taken place between 2018-10-01/12:00 and 2014-11-01/12:00`,
	Run: func(cmd *cobra.Command, args []string) {

		tmp_s, _ := cmd.Flags().GetString("start")
		tmp_e, _ := cmd.Flags().GetString("end")
		if service.GetFlag() == true {
			service.Query_meeting(tmp_s, tmp_e)
		} else {
			fmt.Println("You have not sign in!")
		}
		
	},
}

func init() {
	RootCmd.AddCommand(query_meetingCmd)

	query_meetingCmd.Flags().StringP("start", "s", "", "start time")
	query_meetingCmd.Flags().StringP("end", "e", "", "end time")

}
