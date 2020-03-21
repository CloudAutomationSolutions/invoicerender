package models

// Address - The address structure is used for both parties when rendering an invoice.
type Address struct {
	StreetAndNumber string `json:"street" yaml:"street" dynamodbav:"street"`
	PostCode        string `json:"post-code" yaml:"post-code" dynamodbav:"post-code"`
	City            string `json:"city" yaml:"city" dynamodbav:"city"`
	Country         string `json:"country" yaml:"country" dynamodbav:"country"`
}

// Bank - These bank details are meant for the party that generated the invoice. Here is where the payment should go after the invoice is processed by the receiving party.
type Bank struct {
	Name  string `json:"name" yaml:"name" dynamodbav:"name"`
	IBAN  string `json:"iban" yaml:"iban" dynamodbav:"iban"`
	Swift string `json:"swift" yaml:"swift" dynamodbav:"swift"`
}

// ProvidedService - The items that will be invoiced for will have this structure. Multiple can be present on an invoice.
type ProvidedService struct {
	Name          string  `json:"name" yaml:"name" dynamodbav:"name"`
	Quantity      int     `json:"quantity" yaml:"quantity" dynamodbav:"quantity"`
	VATPercentage float64 `json:"vat-percentage" yaml:"vat-percentage" dynamodbav:"vat-percentage"`

	UnitPrice     float64 `json:"unit-price" yaml:"unit-price" dynamodbav:"unit-price"`
	TotalNetPrice float64 `json:"total-net-price" yaml:"total-net-price" dynamodbav:"total-net-price"`
	VATAmount     float64 `json:"vat-amount" yaml:"vat-amount" dynamodbav:"vat-amount"`
	TotalGross    float64 `json:"total-gross" yaml:"total-gross" dynamodbav:"total-gross"`
}
