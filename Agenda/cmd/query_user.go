// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"agenda/service"
	"github.com/spf13/cobra"
)



// query_userCmd represents the query_user command
var query_userCmd = &cobra.Command{
	Use:   "query_user -n [username]",
	Short: "query a user through username",
	Long: `Input command  like : query_user -n wang

	then if the user exist , we will show you its information`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		name, _ := cmd.Flags().GetString("name")
		if service.GetFlag() == true {
			service.Query_user(name)
		} else {
			fmt.Println("You havn't log in!")
		}
		
	},
}

func init() {
	RootCmd.AddCommand(query_userCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// query_userCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// query_userCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	query_userCmd.Flags().StringP("name", "n", "", "user name")

}
