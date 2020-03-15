package models

// Invoice - The details of the party receiving the invoice.
type Invoice struct {
	HeaderText string `yaml:"header-text" dynamodbav:"header-text"`
	LogoPath   string `yaml:"logo-path" dynamodbav:"logo-path"`

	ID        string `yaml:"id" dynamodbav:"id"`
	IssueDate string `yaml:"issue-date" dynamodbav:"issue-date"`
	DueDate   string `yaml:"due-date" dynamodbav:"due-date"`

	Issuer *Issuer `yaml:"issuer" dynamodbav:"issuer"`
	Client *Client `yaml:"client" dynamodbav:"client"`

	Subject          []string          `yaml:"subject" dynamodbav:"subject"`
	ProvidedServices []ProvidedService `yaml:"provided-services" dynamodbav:"provided-services"`

	TotalNetPrice   float64 `yaml:"total-net-price" dynamodbav:"total-net-price"`
	TotalVATAmount  float64 `yaml:"total-vat-amount" dynamodbav:"total-vat-amount"`
	TotalGrossPrice float64 `yaml:"total-gross-price" dynamodbav:"total-gross-price"`

	Notes  string `yaml:"notes" dynamodbav:"notes"`
	Footer string `yaml:"footer" dynamodbav:"footer"`
}
