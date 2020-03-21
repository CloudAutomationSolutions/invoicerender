package models

// Invoice - The details of the party receiving the invoice.
type Invoice struct {
	HeaderText string `json:"header-text" yaml:"header-text" dynamodbav:"header-text"`
	LogoPath   string `json:"logo-path" yaml:"logo-path" dynamodbav:"logo-path"`

	ID        string `json:"id" yaml:"id" dynamodbav:"id"`
	IssueDate string `json:"issue-date" yaml:"issue-date" dynamodbav:"issue-date"`
	DueDate   string `json:"due-date" yaml:"due-date" dynamodbav:"due-date"`

	Issuer *Issuer `json:"issuer" yaml:"issuer" dynamodbav:"issuer"`
	Client *Client `json:"client" yaml:"client" dynamodbav:"client"`

	Subject          []string          `json:"subject" yaml:"subject" dynamodbav:"subject"`
	ProvidedServices []ProvidedService `json:"provided-services" yaml:"provided-services" dynamodbav:"provided-services"`

	TotalNetPrice   float64 `json:"total-net-price" yaml:"total-net-price" dynamodbav:"total-net-price"`
	TotalVATAmount  float64 `json:"total-vat-amount" yaml:"total-vat-amount" dynamodbav:"total-vat-amount"`
	TotalGrossPrice float64 `json:"total-gross-price" yaml:"total-gross-price" dynamodbav:"total-gross-price"`

	Notes  string `json:"notes" yaml:"notes" dynamodbav:"notes"`
	Footer string `json:"footer" yaml:"footer" dynamodbav:"footer"`
}
