/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
	yaml "gopkg.in/yaml.v2"

	"github.com/CloudAutomationSolutions/invoicerender/pkg/models"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Use to list items that have been provided or generated",
	Long: `Use this to list clients, invoices etc.
Example for listing clients

invoicerender list clients
	`,

	Args: cobra.ExactArgs(1),

	RunE: func(cmd *cobra.Command, args []string) error {
		var config models.Configuration

		configBytes, err := ioutil.ReadFile(cfgFile)
		if err != nil {
			return err
		}

		yaml.Unmarshal(configBytes, &config)

		w := tabwriter.NewWriter(os.Stdout, 0, 8, 2, ' ', 0)

		switch args[0] {
		case "clients":
			fmt.Fprintln(w, "NAME\tID")
			for _, client := range config.Clients {
				fmt.Fprintf(w, "%s\t%d\n", client.Name, client.ID)
			}
		}
		w.Flush()
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
