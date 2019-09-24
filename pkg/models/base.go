package models

// Address - The address structure is used for both parties when rendering an invoice.
type Address struct {
	StreetAndNumber string `yaml:"street"`
	PostCode        string `yaml:"post-code"`
	City            string `yaml:"city"`
	Country         string `yaml:"country"`
}

// Bank - These bank details are meant for the party that generated the invoice. Here is where the payment should go after the invoice is processed by the receiving party.
type Bank struct {
	Name  string `yaml:"name"`
	IBAN  string `yaml:"iban"`
	Swift string `yaml:"swift"`
}

// ProvidedService - The items that will be invoiced for will have this structure. Multiple can be present on an invoice.
type ProvidedService struct {
	Name          string  `yaml:"name"`
	Quantity      int     `yaml:"quantity"`
	VATPercentage float64 `yaml:"vat-percentage"`

	UnitPrice     float64 `yaml:"unit-price"`
	TotalNetPrice float64 `yaml:"total-net-price"`
	VATAmount     float64 `yaml:"vat-amount"`
	TotalGross    float64 `yaml:"total-gross"`
}
