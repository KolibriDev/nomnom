// Copyright ©2015 NAME HERE <EMAIL ADDRESS>
//
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
//
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// webCmd respresents the web command
var webCmd = &cobra.Command{
	Use:   "web",
	Short: "Clones a web application project and performs some basic replacements",
	Long:  `Clones a given web application project and performs basic name replacements.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("web called")
	},
}

func init() {
	RootCmd.AddCommand(webCmd)

	// Here you will define your flags and configuration settings

	// Cobra supports Persistent Flags which will work for this command and all subcommands
	// webCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command is called directly
	// webCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle" )

}
