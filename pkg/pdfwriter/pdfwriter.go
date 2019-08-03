package pdfwriter

import (
	"fmt"
	"strconv"

	"github.com/jung-kurt/gofpdf"
	"golang.org/x/text/language"
	"golang.org/x/text/message"

	"github.com/CloudAutomationSolutions/invoicerender/models"
)

const (
	baseLeftAlignment = 8
)

var (
	tableColumnSizes = []float64{10, 55, 10, 23, 25, 15, 23, 25}
)

// writeColumn - writes a column based on a string array. Each entry is a cell.
func writeColumn(x, y, cellWidth float64, pdf *gofpdf.Fpdf, textList []string) {
	pdf.SetFont("literata", "", 12)
	pdf.SetFillColor(220, 237, 255)
	pdf.SetY(y)

	tr := pdf.UnicodeTranslatorFromDescriptor("cp1258")

	for _, str := range textList {
		pdf.SetX(x)
		pdf.CellFormat(cellWidth, 6, tr(str), "1", 0, "L", true, 0, "")
		pdf.Ln(-1)
	}
}

// writeHeader - writes a column based on a string array. Each entry is a cell.
func writeHeader(pdf *gofpdf.Fpdf, text, imagePath string) {
	pdf.SetFillColor(220, 237, 255)

	pdf.SetXY(baseLeftAlignment, 15)

	tr := pdf.UnicodeTranslatorFromDescriptor("cp1258")
	pdf.CellFormat(100, 7, tr(text), "0", 0, "C", true, 0, "")

	image(imagePath, pdf)
}

// writeLabel - writes a column based on a string array. Each entry is a cell.
func writeLabel(x, y float64, pdf *gofpdf.Fpdf, text string) {
	pdf.SetFont("literata", "", 12)
	pdf.SetFillColor(255, 255, 255)

	pdf.SetXY(x, y)

	pdf.CellFormat(20, 6, text, "0", 0, "L", true, 0, "")
}

func image(imagePath string, pdf *gofpdf.Fpdf) *gofpdf.Fpdf {
	var opt gofpdf.ImageOptions

	pdf.ImageOptions(imagePath, 170, 10, 25, 0, false, opt, 0, "")

	return pdf
}

// writeTableHeader - writes the header so that each column is properly identified
// TODO: maybe add the size of the cell as well as a parameter along with the item in the hdr slice?
func writeTableHeader(pdf *gofpdf.Fpdf, hdr []string) {
	pdf.SetFont("literata", "", 9)
	pdf.SetFillColor(220, 237, 255)

	for i, str := range hdr {
		pdf.CellFormat(tableColumnSizes[i], 7, str, "1", 0, "C", true, 0, "")
	}

	pdf.Ln(-1)
}

// writeTable - writes the data into the table, line by line
// TODO: This table needs to take in an array of data struct representing an item. That way we can have ints and calculate the total
func writeTable(pdf *gofpdf.Fpdf, tbl [][]string) {
	pdf.SetFont("literata", "", 9)
	pdf.SetFillColor(255, 255, 255)

	tr := pdf.UnicodeTranslatorFromDescriptor("cp1258")
	for i, line := range tbl {
		pdf.CellFormat(tableColumnSizes[0], 7, fmt.Sprintf("%d.", i+1), "1", 0, "C", false, 0, "")
		for j, str := range line {
			// We use j+1 because we already have the first cell on the first column done above.
			pdf.CellFormat(tableColumnSizes[j+1], 7, tr(str), "1", 0, "C", true, 0, "")
		}
		pdf.Ln(-1)
	}
}

// writeTotal - writes the data into the table, line by line
func writeTotal(pdf *gofpdf.Fpdf, amount string) {
	pdf.SetX(2 + baseLeftAlignment + tableColumnSizes[0] + tableColumnSizes[1] + tableColumnSizes[2] + tableColumnSizes[3])

	tr := pdf.UnicodeTranslatorFromDescriptor("cp1258")

	pdf.SetFont("literataBold", "", 9)
	pdf.SetFillColor(220, 237, 255)
	pdf.CellFormat(tableColumnSizes[4], 7, "Grand Total", "1", 0, "C", true, 0, "")
	pdf.CellFormat(tableColumnSizes[5], 7, "", "1", 0, "C", true, 0, "")
	pdf.CellFormat(tableColumnSizes[6], 7, "", "1", 0, "C", true, 0, "")

	pdf.CellFormat(tableColumnSizes[7], 7, tr(amount), "1", 0, "C", true, 0, "")
}

// writeFooter - writes the data into the table, line by line
func writeFooter(pdf *gofpdf.Fpdf, notes, footer string) {
	pdf.SetDrawColor(220, 237, 255)
	pdf.SetLineWidth(1)
	pdf.SetFont("literata", "", 8)
	pdf.SetLineCapStyle("round")

	tr := pdf.UnicodeTranslatorFromDescriptor("cp1258")

	pdf.Line(baseLeftAlignment, 248, 200, 248)
	pdf.SetXY(baseLeftAlignment, 249)

	pdf.CellFormat(tableColumnSizes[2], 5, "Notes:", "", 0, "L", false, 0, "")
	notesLines := pdf.SplitLines([]byte(notes), 185)
	for _, noteLine := range notesLines {
		pdf.Ln(-1)
		pdf.CellFormat(185, 5, tr(string(noteLine)), "1", 0, "L", true, 0, "")
	}
	pdf.SetY(271)
	pdf.CellFormat(185, 5, tr(footer), "", 0, "L", false, 0, "")
}

// WriteInvoicePDF - Based on the defined Invoice model, this function takes care of the fields it needs to write in the PDF. It prints and formats them based on the configuration.
func WriteInvoicePDF(invoice *models.Invoice, outputPDFPath string) error {

	// We are instantiating the main PDF object
	pdf := gofpdf.New("P", "mm", "A4", "../font/literata")
	pdf.AddPage()

	// We are importing a font that supports latin charatcters. This PDF library we are using does not support utf-8. TTF files need to be compiled with the proper pages to get some characters.
	pdf.AddFont("literata", "", "Literata-Regular.json")
	pdf.AddFont("literataBold", "", "Literata-Bold.json")
	pdf.SetFont("literata", "", 16)

	// We write the header title and the logo.
	// TODO: Make logo optional
	writeHeader(pdf, invoice.HeaderText, invoice.LogoPath)

	// Write the labels for the invoice identification fields.
	writeLabel(baseLeftAlignment, 40, pdf, "Invoice ID:")
	writeLabel(baseLeftAlignment, 46, pdf, "Issue date:")
	writeLabel(baseLeftAlignment, 52, pdf, "Due date:")
	// Write the fields next to the appropiate labels above
	writeColumn(40, 40, 25, pdf, []string{invoice.ID, invoice.IssueDate, invoice.DueDate})

	// write the labels. Align the Issuer and Client columns so that they match this.
	writeLabel(baseLeftAlignment, 70, pdf, "From:")
	writeLabel(baseLeftAlignment, 94, pdf, "VAT no.:")
	writeLabel(baseLeftAlignment, 100, pdf, "ESt/USt no.:")
	writeLabel(baseLeftAlignment, 106, pdf, "Bank:")
	writeLabel(baseLeftAlignment, 112, pdf, "Iban:")
	writeLabel(baseLeftAlignment, 118, pdf, "BIC:")
	// Add one more label. Override the "From:" label on the client side.
	writeLabel(110, 70, pdf, "For:")

	// Issuer details
	writeColumn(40, 70, 55, pdf, []string{
		invoice.Issuer.Name,
		invoice.Issuer.Address.StreetAndNumber,
		fmt.Sprintf("%s, %s", invoice.Issuer.Address.PostCode, invoice.Issuer.Address.City),
		invoice.Issuer.Address.Country,
		invoice.Issuer.VATNumber,
		invoice.Issuer.IncomeSalesTaxNumber,
		invoice.Issuer.Bank.Name,
		invoice.Issuer.Bank.IBAN,
		invoice.Issuer.Bank.Swift,
	})

	// Client details
	writeColumn(135, 70, 55, pdf, []string{
		invoice.Client.Name,
		invoice.Client.Address.StreetAndNumber,
		fmt.Sprintf("%s, %s", invoice.Client.Address.PostCode, invoice.Client.Address.City),
		invoice.Client.Address.Country,
		invoice.Client.VATNumber,
	})

	// Invoice subject label
	writeLabel(baseLeftAlignment, 135, pdf, "Subject:")
	// Invoice subject
	writeColumn(40, 135, 150, pdf, invoice.Subject)

	// Set position for the invoice items
	pdf.SetY(160)
	// Write the header of the items table
	writeTableHeader(pdf, []string{"No.", "Item", "Qty", "Unit Net Price", "Total net", "VAT %", "VAT amount", "Total Gross"})

	// This printer object is used to split larget numbers. The format is now European: 11.200,34
	// TODO: Make configurable for the UK/US/Indian market. Maybe switch to the go-money library?
	p := message.NewPrinter(language.Make("de"))

	// Itterate over all the Provided services declared on the invoice an format them so that they look appropiate on the PDF
	var billableItems [][]string
	for _, providedService := range *invoice.ProvidedServices {
		billableItems = append(
			billableItems,
			[]string{
				providedService.Name,
				strconv.Itoa(providedService.Quantity),
				p.Sprintf("%.2f €", providedService.UnitPrice),
				p.Sprintf("%.2f €", providedService.TotalNetPrice),
				p.Sprintf("%d %", providedService.VATPercentage),
				p.Sprintf("%.2f €", providedService.VATAmount),
				p.Sprintf("%.2f €", providedService.TotalGross),
			})
	}

	// We want to have at least 4 rows to make the invoice look pretty.
	if len(billableItems) < 4 {
		for i := 0; i < 4-len(billableItems); i++ {
			billableItems = append(billableItems, []string{"", "", "", "", "", "", ""})
		}
	}

	// Write the items that the invoice is generated for
	// TODO: Limit the items to 10
	writeTable(pdf, billableItems)
	// Write the total amount
	// TODO: Add total VAT and total Net
	writeTotal(pdf, p.Sprintf("%.2f €", invoice.TotalGrossPrice))

	// Write the footer with all the notes
	writeFooter(pdf, invoice.Notes, invoice.Footer)

	// Write out the PDF at the specified path
	err := pdf.OutputFileAndClose(outputPDFPath)
	if err != nil {
		return err
	}

	return nil
}
