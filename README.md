# Overview

Project meant to generate invoices as PDF. It takes a client and a issuer as configuration during the init phase.
During execution some information needs to be provided.
The project idea came about after offering consultancy in the German market (IT Beratung) as a freelancer.

# Design
Data structure documentation can be found [here](https://drive.google.com/file/d/1TqzJZUncPmCPW45T3SlCebUZuGn5hV-N/view?usp=sharing).

# Usage
## Init
Initial configuration will be done using the `init` command. Interactively fill in the client and the issuer information. This data will be saved under your home directory by default: `~/.invoicerender`. Override this using the `INVOICERENDER_HOME` environment variable.
```
$ invoicerender init --client-count=2

Please provide the following information for the invoicerender initial configuration:

Header text: John Doe - Consulting
Footer text: No VAT shall be requested for these services under the EU law ...
Header logo: /tmp/images/logo.png

Please profide the issuer's information:
Name: John Doe
Street name and number: Müsterman str. 123
Post Code: 40123
City: München
Country: România
VAT No.: DE01234567
Income and sales tax number: 123456123456
Bank name: ING
IBAN: DE0123929838192389128907894567412
Swift/BIC Code: COBADEFXXXX

Client1 Information
...

Client2 Information
...

```

## List
Use to list items that have been provided or generated.

### Clients
List the registered clients
```bash
$ invoicerender list clients
NAME       ID
First      1
Second     2 
```

## Issue

Generate an invoice for a named client with just one billable service. Use this when just one item description is required:
```bash
$ invoicerender issue --client=<client name> \
                      --item-description="Services provided last month" \
                      --item-unit-price=100 \
                      --item-currency=EUR \
                      --item-quantity=21 \
                      --item-vat-percentage=19
```

Generate an invoice for a named client with multiple billabe items:
```bash
$ invoicerender issue --client=<client name> --interactive --item-count=2

item1 Description: Services provided last month
item1 Unit Price: 100
item1 Currency: EUR
item1 Quantity: 21
item1 VAT Percentage: 19

item2 Description: Miscellaneous billable items
item2 Unit Price: 50
item2 Currency: EUR
item2 Quantity: 3
item2 VAT Percentage: 19
```