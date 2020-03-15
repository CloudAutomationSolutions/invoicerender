package models

// Client - The details of the party receiving the invoice.
type Client struct {
	ID        int      `yaml:"id" dynamodbav:"id"`
	Name      string   `yaml:"name" dynamodbav:"name"`
	Address   *Address `yaml:"address" dynamodbav:"address"`
	VATNumber string   `yaml:"vat-number" dynamodbav:"vat-number"`
}

// Issuer - The details of the party generating the invoice.
type Issuer struct {
	Name                 string   `yaml:"name" dynamodbav:"name"`
	Address              *Address `yaml:"address" dynamodbav:"address"`
	VATNumber            string   `yaml:"vat-number" dynamodbav:"vat-number"`
	IncomeSalesTaxNumber string   `yaml:"income-sales-tax-number" dynamodbav:"income-sales-tax-number"`
	Bank                 *Bank    `yaml:"bank" dynamodbav:"bank"`
}
