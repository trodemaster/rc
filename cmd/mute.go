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
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/cobra"
)

// muteCmd represents the mute command
var muteCmd = &cobra.Command{
	Use:   "mute",
	Short: "Mutes the speakers",
	Long:  `Pushing this button toggles the muted state of the speaker`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("mute called")
		client := &http.Client{}

		body := []byte("{\n  \"frequency\": \"38000\",\n  \"preamble\": \"347,174,21,21,21,21,21,21,21,21,21,21,21,21,21,21,21,21,21,65,21,65,21,65,21,65,21,65,21,65,21,65,21,65,21,21,21,21,21,21,21,65,21,21,21,21,21,65,21,21,21,65,21,65,21,65,21,21,21,65,21,65,21,21,21,65,21,1572\",\n  \"irCode\": \"347,87,21,3708\",\n  \"repeat\": \"1\"\n}")

		req, _ := http.NewRequest("POST", "http://itach/api/v1/irports/3/sendir", bytes.NewBuffer(body))

		req.Header.Add("Content-Length", "171")
		req.Header.Add("Content-Type", "application/json")

		resp, err := client.Do(req)

		if err != nil {
			fmt.Println("Errored when sending request to the server")
			return
		}

		defer resp.Body.Close()
		respBody, _ := ioutil.ReadAll(resp.Body)

		fmt.Println(resp.Status)
		fmt.Println(string(respBody))
	},
}

func init() {
	rootCmd.AddCommand(muteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// muteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// muteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
