# Overview

Project meant to generate invoices as PDF. It takes a client and a issuer as configuration during the init phase.
During execution some information needs to be provided.
The project idea came about after offering consultancy in the German market (IT Beratung) as a freelancer.

# Design
Data structure documentation can be found [here](https://drive.google.com/file/d/1TqzJZUncPmCPW45T3SlCebUZuGn5hV-N/view?usp=sharing).

# Usage
## Init
Initial configuration will be done using the `init` command. Interactively fill in the client and the issuer information. This data will be saved under your home directory by default: `~/.invoicerender`. Override this using the `INVOICERENDER_HOME` environment variable.
```bash
invoicerender init
```

## Issue

Generate an invoice for a named client with just one billable service:
```bash
invoicerender issue --client=<client name> \
                    --item-description="Services provided last month" \
                    --item-unit-price=100 \
                    --item-currency=EUR \
                    --item-quantity=21 \
                    --item-vat-percentag=19
```

Generate an invoice for a named client with multiple billabe items:
```bash
invoicerender issue --client=<client name> --interactive --item-count=2

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