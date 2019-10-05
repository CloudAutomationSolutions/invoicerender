package models

import (
	"fmt"
	"os"
	"path/filepath"

	yaml "gopkg.in/yaml.v2"
)

// Configuration - The information that will be stored in order to fill in the template and generate the invoices on demand.
type Configuration struct {
	HeaderText string `yaml:"header-text"`
	LogoPath   string `yaml:"logo-path"`

	Issuer  *Issuer  `yaml:"issuer"`
	Clients []Client `yaml:"clients"`

	LastUsedID int `yaml:"last-used-id"`
	// The ID of an invoice is "x/y" where "x" is a number that increments with every use and "y" is the year. If a new year has come, we need to start x from 1, not LastUsedID.
	YearForLastUsedID int `yaml:"year-for-last-used-id"`

	SavedNotes [][]string
	Footer     string `yaml:"footer"`

	OutputDirectory string `yaml:"output-directory"`
}

// WriteToDisk - Write the configuration in place, where the specified configuration file is, to save it for later use
func (c *Configuration) WriteToDisk(cfgFile string) error {

	// cfgFile is declared in root.go (same package). We ensure that it's directory exists
	configDir := filepath.Dir(cfgFile)
	err := os.MkdirAll(configDir, os.ModePerm)

	// TODO (trivial): Find a way to see when the directory was created. MkdirAll returns nil if already exists. 
	if err != nil && !os.IsExist(err) {
		return err
	}

	f, err := os.OpenFile(cfgFile, os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()

	d, err := yaml.Marshal(c)
	if err != nil {
		fmt.Println("Error Marshal to file...")
		return err
	}

	_, err = f.Write(d)
	if err != nil {
		fmt.Println("Error while Writing to file")
		return err
	}

	return nil
}
