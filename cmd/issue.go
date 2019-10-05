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
	"io/ioutil"
	"path/filepath"
	"strings"
	"time"

	"github.com/spf13/cobra"
	yaml "gopkg.in/yaml.v2"

	"github.com/CloudAutomationSolutions/invoicerender/pkg/models"
	"github.com/CloudAutomationSolutions/invoicerender/pkg/pdfwriter"
)

const (
	layoutEU = "02.01.2006"
)

// issueCmd represents the issue command
var issueCmd = &cobra.Command{
	Use:   "issue",
	Short: "Issue an invoice for a saved client",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		clientID, err := cmd.Flags().GetInt("client-id")
		if err != nil {
			return err
		}
		if clientID == 0 {
			return errors.New("please provide mandatory flag --client-id")
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

		invoiceSubject, err := cmd.Flags().GetString("subject")
		if err != nil {
			return err
		}
		if len(invoiceSubject) == 0 && interactive == false {
			return errors.New("please provide mandatory flag --subject")
		}

		return nil
	},

	RunE: func(cmd *cobra.Command, args []string) error {

		var config models.Configuration
		var invoice models.Invoice
		clientID, err := cmd.Flags().GetInt("client-id")

		itemQuantity, err := cmd.Flags().GetInt("item-quantity")
		if err != nil {
			return err
		}
		itemVATPercentage, err := cmd.Flags().GetInt("item-vat-percentage")
		if err != nil {
			return err
		}
		itemDescription, err := cmd.Flags().GetString("item-description")
		if err != nil {
			return err
		}

		itemUnitPrice, err := cmd.Flags().GetInt("item-unit-price")
		if err != nil {
			return err
		}

		invoiceSubject, err := cmd.Flags().GetString("subject")
		if err != nil {
			return err
		}

		configBytes, err := ioutil.ReadFile(cfgFile)
		if err != nil {
			return err
		}
		yaml.Unmarshal(configBytes, &config)

		invoice.HeaderText = config.HeaderText
		invoice.LogoPath = config.LogoPath

		invoice.ID = fmt.Sprintf("%d/%d", config.LastUsedID+1, config.YearForLastUsedID)
		invoice.IssueDate = time.Now().Format(layoutEU)
		invoice.DueDate = time.Now().AddDate(0, 0, 25).Format(layoutEU) // We always add 25 days until the payment day

		invoice.Issuer = config.Issuer

		for _, client := range config.Clients {
			if client.ID == clientID {
				invoice.Client = &client
				break
			}
		}

		invoice.ProvidedServices = []models.ProvidedService{
			{
				Name:          itemDescription,
				Quantity:      itemQuantity,
				VATPercentage: float64(itemVATPercentage),
				UnitPrice:     float64(itemUnitPrice),
				TotalNetPrice: float64(itemUnitPrice * itemQuantity),
				VATAmount:     float64(itemVATPercentage) / 100.00 * float64(itemQuantity*itemUnitPrice),
				TotalGross:    float64(itemVATPercentage)/100.00*float64(itemQuantity*itemUnitPrice) + float64(itemQuantity*itemUnitPrice),
			},
		}
		invoice.Subject = strings.Split(invoiceSubject, "|") //TODO: Make default

		for _, service := range invoice.ProvidedServices {
			invoice.TotalNetPrice += service.TotalNetPrice
			invoice.TotalGrossPrice += service.TotalGross
			invoice.TotalVATAmount += service.VATAmount
		}

		// TODO: Don't forget to update the invoice ID in the config if the file was written with success
		outputPath := filepath.Join(configDir, fmt.Sprintf("%s-%s.pdf", invoice.Client.Name, invoice.IssueDate))

		err = pdfwriter.WriteInvoicePDF(&invoice, outputPath)
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(issueCmd)

	issueCmd.PersistentFlags().BoolP("interactive", "i", false, "Use this mode when multiple items need to be invoiced. An interactive prompt will ask the details for each")
	issueCmd.PersistentFlags().Int("item-count", 0, "Use this in interactive mode when multiple items need to be invoiced. This will set the number of items to ask details for")

	issueCmd.Flags().IntP("client-id", "c", 0, "The client name to generate the invoice for. Example --client-id=1")
	issueCmd.Flags().String("subject", "", "The invoice subect. Example --subject=\"Services for the SysOps role - 123 Project\"")
	issueCmd.Flags().String("item-description", "", "Item or services to add to the invoice. Example --item-description=\"Services provided for the last month\"")
	issueCmd.Flags().Int("item-unit-price", 0, "The base price for each item. Example --item-unit-price=500")
	issueCmd.Flags().Int("item-quantity", 20, "The amount of units to be invoiced. Default: 20")
	issueCmd.Flags().Int("item-vat-percentage", 19, "The amount of units to be invoiced. Default: 19")
}
