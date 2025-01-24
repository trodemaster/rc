package cmd

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/cobra"
)

// hdmi2Cmd represents the hdmi2 command
var hdmi2Cmd = &cobra.Command{
	Use:   "hdmi2",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hdmi2 called")
		client := &http.Client{}

		body := []byte("{\n  \"frequency\": \"38028\",\n  \"preamble\": \"\",\n  \"irCode\": \"343,172,21,22,21,22,21,65,21,22,21,22,21,22,21,22,21,22,21,65,21,65,21,22,21,65,21,65,21,65,21,65,21,65,21,22,21,65,21,22,21,22,21,22,21,22,21,22,21,65,21,65,21,22,21,65,21,65,21,65,21,65,21,65,21,22,21,1673,343,86,21,3732\",\n  \"repeat\": \"1\"\n}")

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
	rootCmd.AddCommand(hdmi2Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// hdmi2Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// hdmi2Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
