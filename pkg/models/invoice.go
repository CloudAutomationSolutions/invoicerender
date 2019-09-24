package models

// Invoice - The details of the party receiving the invoice.
type Invoice struct {
	HeaderText string `yaml:"header-text"`
	LogoPath   string `yaml:"logo-path"`

	ID        string `yaml:"id"`
	IssueDate string `yaml:"issue-date"`
	DueDate   string `yaml:"due-date"`

	Issuer *Issuer `yaml:"issuer"`
	Client *Client `yaml:"client"`

	Subject          []string           `yaml:"subject"`
	ProvidedServices []ProvidedService `yaml:"provided-services"`

	TotalNetPrice   float64 `yaml:"total-net-price"`
	TotalVATAmount  float64 `yaml:"total-vat-amount"`
	TotalGrossPrice float64 `yaml:"total-gross-price"`

	Notes  string `yaml:"notes"`
	Footer string `yaml:"footer"`
}
