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
	"errors"
	"fmt"

	// "github.com/CloudAutomationSolutions/invoicerender/pkg/models"
	"github.com/spf13/cobra"
)

// issueCmd represents the issue command
var issueCmd = &cobra.Command{
	Use:   "issue",
	Short: "Issue an invoice for a saved client",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		client, err := cmd.Flags().GetString("client")
		if err != nil {
			return err
		}
		if len(client) == 0 {
			return errors.New("please provide mandatory flag --client")
		}

		interactive, err := cmd.Flags().GetBool("interactive")
		if err != nil {
			return err
		}

		itemDescription, err := cmd.Flags().GetString("item-description")
		if err != nil {
			return err
		}
		if len(itemDescription) == 0 && interactive == false {
			return errors.New("please provide mandatory flag --item-description")
		}

		itemUnitPrice, err := cmd.Flags().GetInt("item-unit-price")
		if err != nil {
			return err
		}
		if itemUnitPrice == 0 && interactive == false {
			return errors.New("please provide mandatory flag --item-unit-price")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		// var invoice models.Invoice
		fmt.Println("issue called")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(issueCmd)

	issueCmd.PersistentFlags().BoolP("interactive", "i", false, "Use this mode when multiple items need to be invoiced. An interactive prompt will ask the details for each")
	issueCmd.PersistentFlags().Int("item-count", 0, "Use this in interactive mode when multiple items need to be invoiced. This will set the number of items to ask details for")

	issueCmd.Flags().String("client", "", "The client name to generate the invoice for. Example --client=acme")
	issueCmd.Flags().String("item-description", "", "Item or services to add to the invoice. Example --item-description=\"Services provided for the last month\"")
	issueCmd.Flags().Int("item-unit-price", 0, "The base price for each item. Example --item-unit-price=500")
	issueCmd.Flags().String("item-currency", "EUR", "The currency to use when writing the invoice. [EUR, USD, RON, GBP]. Defaults to EUR")
	issueCmd.Flags().Int("item-quantity", 20, "The amount of units to be invoiced. Default: 20")
	issueCmd.Flags().Int("item-vat-percentage", 19, "The amount of units to be invoiced. Default: 19")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// issueCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
