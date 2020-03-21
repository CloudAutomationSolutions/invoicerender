package models

// Client - The details of the party receiving the invoice.
type Client struct {
	ID        string   `json:"id" yaml:"id" dynamodbav:"id"`
	Name      string   `json:"name" yaml:"name" dynamodbav:"name"`
	Address   *Address `json:"address" yaml:"address" dynamodbav:"address"`
	VATNumber string   `json:"vat-number" yaml:"vat-number" dynamodbav:"vat-number"`
}

// Issuer - The details of the party generating the invoice.
type Issuer struct {
	Name                 string   `json:"name" yaml:"name" dynamodbav:"name"`
	Address              *Address `json:"address" yaml:"address" dynamodbav:"address"`
	VATNumber            string   `json:"vat-number" yaml:"vat-number" dynamodbav:"vat-number"`
	IncomeSalesTaxNumber string   `json:"income-sales-tax-number" yaml:"income-sales-tax-number" dynamodbav:"income-sales-tax-number"`
	Bank                 *Bank    `json:"bank" yaml:"bank" dynamodbav:"bank"`
}
