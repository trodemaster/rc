package cmd

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/cobra"
)

// upCmd represents the up command
var upCmd = &cobra.Command{
	Use:   "up",
	Short: "Volume up",
	Long:  `This turns the volume of the speakers up`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("up called")
		client := &http.Client{}

		body := []byte("{\n  \"frequency\": \"38000\",\n  \"preamble\": \"347,174,21,21,21,21,21,21,21,21,21,21,21,21,21,21,21,21,21,65,21,65,21,65,21,65,21,65,21,65,21,65,21,65,21,65,21,21,21,21,21,21,21,21,21,21,21,65,21,21,21,21,21,65,21,65,21,65,21,65,21,65,21,21,21,65,21,1572\",\n  \"irCode\": \"20DF40BF\",\n  \"repeat\": \"1\"\n}")

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
	rootCmd.AddCommand(upCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// upCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// upCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
