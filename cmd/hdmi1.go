package cmd

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/cobra"
)

// hdmi1Cmd represents the hdmi1 command
var hdmi1Cmd = &cobra.Command{
	Use:   "hdmi1",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hdmi1 called")
		client := &http.Client{}

		body := []byte("{\n  \"frequency\": \"38028\",\n  \"preamble\": \"\",\n  \"irCode\": \"20DF738C\",\n  \"repeat\": \"1\"\n}")

		req, _ := http.NewRequest("POST", "http://itach/api/v1/irports/3/sendir", bytes.NewBuffer(body))

		req.Header.Add("Content-Length", "123")
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
	rootCmd.AddCommand(hdmi1Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// hdmi1Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// hdmi1Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
