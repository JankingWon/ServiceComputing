
package cmd

import (
	"fmt"
	"agenda/service"
	"github.com/spf13/cobra"
)


// 创建会议命令
var create_meetingCmd = &cobra.Command{
	Use:   "create_meeting -t [meeting title] -s [start time] -e [end time]",
	Short: "create a meeting",
	Long: `Input command model like : create_meeting -t Golang -s 2018-11-01/14:00 -e 2017-11-01/16:00

	in order to avoid longer command , we will call a fuction to input paticilator after you input command.`,
	Run: func(cmd *cobra.Command, args []string) {
		//定义输入计数器 
		tmp_t, _ := cmd.Flags().GetString("title")
		tmp_s, _ := cmd.Flags().GetString("start")
		tmp_e, _ := cmd.Flags().GetString("end")
		if service.GetFlag() == true {
			service.Create_meeting(tmp_t, tmp_s, tmp_e)
		} else {
			fmt.Println("You have not sign in!")
		}
		
	},
}

func init() {
	RootCmd.AddCommand(create_meetingCmd)
	create_meetingCmd.Flags().StringP("title", "t", "", "meeting title")
	create_meetingCmd.Flags().StringP("start",  "s", "", "start time")
	create_meetingCmd.Flags().StringP("end",  "e", "", "end time")

}
