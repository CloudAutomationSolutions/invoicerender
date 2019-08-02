package models

// Client - The details of the party receiving the invoice.
type Client struct {
	Name      string   `yaml:"name"`
	Address   *Address `yaml:"address"`
	VATNumber string   `yaml:"address"`
}

// Issuer - The details of the party generating the invoice.
type Issuer struct {
	Name                 string   `yaml:"name"`
	Address              *Address `yaml:"address"`
	VATNumber            string   `yaml:"address"`
	IncomeSalesTaxNumber string   `yaml:"income-sales-tax-number"`
	Bank                 *Bank    `yaml:"bank"`
}
