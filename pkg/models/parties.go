package models

// Client - The details of the party receiving the invoice.
type Client struct {
	ID        int      `yaml:"id"`
	Name      string   `yaml:"name"`
	Address   *Address `yaml:"address"`
	VATNumber string   `yaml:"vat-number"`
}

// Issuer - The details of the party generating the invoice.
type Issuer struct {
	Name                 string   `yaml:"name"`
	Address              *Address `yaml:"address"`
	VATNumber            string   `yaml:"vat-number"`
	IncomeSalesTaxNumber string   `yaml:"income-sales-tax-number"`
	Bank                 *Bank    `yaml:"bank"`
}
