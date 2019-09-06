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
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/CloudAutomationSolutions/invoicerender/pkg/models"
	"github.com/spf13/cobra"
	yaml "gopkg.in/yaml.v2"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Invoicerender initial configuration",
	RunE: func(cmd *cobra.Command, args []string) error {
		// TODO: Read the config file and see if there is data already there. If there is, leave as default.
		fmt.Println()
		fmt.Println("Please provide the following information for the invoicerender initial configuration:")
		fmt.Println()

		var err error
		var config models.Configuration
		config.LastUsedID = 0

		year, _, _ := time.Now().Date()
		config.YearForLastUsedID = year

		config.HeaderText, err = getUserInput("Header Text")
		if err != nil {
			return err
		}

		config.LogoPath, err = getUserInput("Header Logo (path)")
		if err != nil {
			return err
		}

		config.Footer, err = getUserInput("Footer Text")
		if err != nil {
			return err
		}

		fmt.Println()
		fmt.Println("Please provide the issuer's configuration:")
		fmt.Println()

		var issuer models.Issuer
		var issuerAddress models.Address
		var issuerBank models.Bank

		// Fetch base issuer details
		issuer.Name, err = getUserInput("Issuer name")
		if err != nil {
			return err
		}

		issuer.VATNumber, err = getUserInput("Issuer VAT Number")
		if err != nil {
			return err
		}

		issuer.IncomeSalesTaxNumber, err = getUserInput("Income and sales tax number")
		if err != nil {
			return err
		}

		// Fetch issuer address details
		issuerAddress.StreetAndNumber, err = getUserInput("Street name and number")
		if err != nil {
			return err
		}

		issuerAddress.PostCode, err = getUserInput("Postcode")
		if err != nil {
			return err
		}

		issuerAddress.City, err = getUserInput("City")
		if err != nil {
			return err
		}

		issuerAddress.Country, err = getUserInput("Country")
		if err != nil {
			return err
		}

		// Fetch issuer bank details
		issuerBank.Name, err = getUserInput("Bank name")
		if err != nil {
			return err
		}

		issuerBank.IBAN, err = getUserInput("IBAN")
		if err != nil {
			return err
		}

		issuerBank.Swift, err = getUserInput("SWIFT/BIC")
		if err != nil {
			return err
		}

		// Add the address and banking details we just fetched to the issuer
		issuer.Address = &issuerAddress
		issuer.Bank = &issuerBank
		// Add the issuer details we just fetched to the config
		config.Issuer = &issuer

		fmt.Println()
		fmt.Println("Please provide the clients' configuration:")
		fmt.Println()

		clientCount, err := cmd.Flags().GetInt("client-count")
		if err != nil {
			return err
		}

		for i := 0; i < clientCount; i++ {
			var client models.Client
			fmt.Printf("Configuration for client no. %d:\n\n", i+1)

			client.ID = i + 1
			client.Name, err = getUserInput("Name")
			if err != nil {
				return err
			}
			client.VATNumber, err = getUserInput("VAT Number")
			if err != nil {
				return err
			}

			var clientAddress models.Address

			clientAddress.StreetAndNumber, err = getUserInput("Street name and number")
			if err != nil {
				return err
			}

			clientAddress.PostCode, err = getUserInput("Postcode")
			if err != nil {
				return err
			}

			clientAddress.City, err = getUserInput("City")
			if err != nil {
				return err
			}

			clientAddress.Country, err = getUserInput("Country")
			if err != nil {
				return err
			}

			client.Address = &clientAddress

			config.Clients = append(config.Clients, client)

		}

		err = writeConfig(&config)
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	initCmd.Flags().Int("client-count", 1, "Number of clients to ask details for. Default: 1")
}

func getUserInput(prompt string) (string, error) {
	fmt.Printf("%s: ", prompt)

	reader := bufio.NewReader(os.Stdin)
	// TODO: See if we can remove error and change the function signature. Will help with verbosity
	text, err := reader.ReadString('\n')

	if err != nil {
		return "", err
	}
	return strings.TrimSpace(text), nil
}

func writeConfig(config *models.Configuration) error {

	// cfgFile is declared in root.go (same package). We ensure that it's directory exists
	configDir := filepath.Dir(cfgFile)
	err := os.MkdirAll(configDir, os.ModePerm)

	if err == nil {
		fmt.Printf("The cofiguration directory \"%s\" did not exist and was created\n", configDir)
	} else if err != nil && !os.IsExist(err) {
		return err
	}

	f, err := os.OpenFile(cfgFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()

	d, err := yaml.Marshal(&config)
	if err != nil {
		fmt.Println("Error Marshal to file...")
		return err
	}

	fmt.Println("Writing yaml configuration to file")
	_, err = f.Write(d)
	if err != nil {
		fmt.Println("Error while Writing to file")
		return err
	}

	return nil
}
