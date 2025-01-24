package cmd

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/cobra"
)

// onCmd represents the on command
var onCmd = &cobra.Command{
	Use:   "on",
	Short: "Power ON the display",
	Long:  `This powers on the display`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("on called")
		client := &http.Client{}

		body := []byte("{\n  \"frequency\": \"38000\",\n  \"preamble\": \"343,172,21,22,21,22,21,65,21,22,21,22,21,22,21,22,21,22,21,65,21,65,21,22,21,65,21,65,21,65,21,65,21,65,21,22,21,65,21,22,21,65,21,22,21,65,21,22,21,22,21,65,21,22,21,65,21,22,21,65,21,22,21,65,21,65,21,1673\",\n  \"irCode\": \"20DF23DC\",\n  \"repeat\": \"1\"\n}")

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
	rootCmd.AddCommand(onCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// onCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// onCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
