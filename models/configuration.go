package models

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
