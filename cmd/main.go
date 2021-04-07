package main

import (
	"fmt"

	u "go-html-to-pdf-demo/src/pdf-generator"
)

type InvoiceRow struct {
	Number string
	Name string
	Description string
	UnitPrice string
	Quantity int
	Total string
}

func main() {
	r := u.NewPdf("")

	//html template path
	templatePath := "templates/index.html"

	//path for download pdf
	outputPath := "./storage/invoice.pdf"

	//html template data
	templateData := struct {
		InvoiceName string
		Rows []InvoiceRow
	}{
		InvoiceName: "INVOICE #889926",
		Rows: getInvoiceRows(),
	}

	if err := r.ParseTemplate(templatePath, templateData); err != nil {
		fmt.Println(err)
		return
	}

	ok, _ := r.GeneratePDF(outputPath)
	fmt.Println(ok, "pdf generated successfully")
}

func getInvoiceRows() []InvoiceRow {
	return []InvoiceRow{
		{
			Number: "01",
			Name: "Website Design",
			Description: "Creating a recognizable design solution based on the company's existing visual identity",
			UnitPrice: "$40.00",
			Quantity: 30,
			Total: "$1,200.00",
		},
		{
			Number: "02",
			Name: "Website Development",
			Description: "Developing a Content Management System-based Website",
			UnitPrice: "$40.00",
			Quantity: 80,
			Total: "$3,200.00",
		},
		{
			Number: "03",
			Name: "Search Engines Optimization",
			Description: "Optimize the site for search engines (SEO)",
			UnitPrice: "$40.00",
			Quantity: 20,
			Total: "$800.00",
		},
	}
}
