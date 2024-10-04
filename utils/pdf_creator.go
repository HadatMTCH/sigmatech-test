package utils

import (
	"base-api/constants"
	"bytes"
	"strings"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

func GeneratePDF(html string, pdfFileName string) (string, error) {
	pdfFilePath := constants.TempDownloadedFileDir + pdfFileName

	// create a new page based on the HTML
	page := wkhtmltopdf.NewPageReader(strings.NewReader(html))
	page.NoBackground.Set(true)
	page.DisableExternalLinks.Set(false)

	// create a new instance of the PDF generator
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		return pdfFilePath, err
	}

	// add page to the PDF generator
	pdfg.AddPage(page)

	// set dpi of the content
	pdfg.Dpi.Set(350)
	pdfg.PageSize.Set(wkhtmltopdf.PageSizeA4)

	// set margins to zero at all direction
	pdfg.MarginBottom.Set(0)
	pdfg.MarginTop.Set(10)
	pdfg.MarginLeft.Set(0)
	pdfg.MarginRight.Set(0)

	// create the exact pdf
	err = pdfg.Create()
	if err != nil {
		return pdfFilePath, err
	}

	// write it into a file
	err = pdfg.WriteFile(constants.TempDownloadedFileDir + pdfFileName)
	if err != nil {
		return pdfFilePath, err
	}

	return pdfFilePath, nil
}

func GeneratePDFMultiplePage(html []string, pdfFileName string) (string, error) {
	pdfFilePath := constants.TempDownloadedFileDir + pdfFileName

	//boffer to combine html string
	buf := new(bytes.Buffer)

	for _, x := range html {
		buf.WriteString(x)
		buf.WriteString(`<html><body><P style="page-break-before: always"><body/></html>`)
	}

	// create a new page based on the combined
	page := wkhtmltopdf.NewPageReader(buf)
	page.NoBackground.Set(true)
	page.DisableExternalLinks.Set(false)

	// create a new instance of the PDF generator
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		return pdfFilePath, err
	}

	// add page to the PDF generator
	pdfg.AddPage(page)
	// set dpi of the content
	pdfg.Dpi.Set(350)
	pdfg.PageSize.Set(wkhtmltopdf.PageSizeA4)

	// set margins to zero at all direction
	pdfg.MarginBottom.Set(30)
	pdfg.MarginTop.Set(30)
	pdfg.MarginLeft.Set(30)
	pdfg.MarginRight.Set(30)

	// create the exact pdf
	err = pdfg.Create()
	if err != nil {
		return pdfFilePath, err
	}

	// write it into a file
	err = pdfg.WriteFile(constants.TempDownloadedFileDir + pdfFileName)
	if err != nil {
		return pdfFilePath, err
	}

	return pdfFilePath, nil
}
