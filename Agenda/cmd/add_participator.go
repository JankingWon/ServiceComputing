// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
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

var name, title string

// add_participatorCmd represents the add_participator command
var add_participatorCmd = &cobra.Command{
	Use:   "add_participator -n [participator name] -t [meeting title]",
	Short: "add a participator to a meeting with title",
	Long: `Input command model like : add_participator -n name -t title`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		title, _ := cmd.Flags().GetString("title")
		name, _ := cmd.Flags().GetString("name")
		if service.GetFlag() == true {
			service.Add_participator(name, title)
		} else {
			fmt.Println("You don't log in!")
		}
		
	},
}

func init() {
	RootCmd.AddCommand(add_participatorCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// add_participatorCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// add_participatorCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	add_participatorCmd.Flags().StringP("name", "n", "", "participator name")
	add_participatorCmd.Flags().StringP("title", "t", "", "meeting title")
}
