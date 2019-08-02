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

	Subject          string             `yaml:"subject"` // TODO: Figure out if this should be []string because we also have a second row that is a reference to a contract. Or add a new field for that?
	ProvidedServices *[]ProvidedService `yaml:"provided-services"`

	TotalNetPrice   float32 `yaml:"total-net-price"`
	TotalVATAmount  float32 `yaml:"total-vat-amount"`
	TotalGrossPrice float32 `yaml:"total-gross-price"`

	Notes string `yaml:"notes"`
	Footer string `yaml:"footer"`
}
