package pdfgenerator

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

type Pdf struct {
	body string
}

func NewPdf(body string) *Pdf {
	return &Pdf{
		body: body,
	}
}

// ParseTemplate parsing template function
func (r *Pdf) ParseTemplate(templateFileName string, data interface{}) error {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return err
	}

	r.body = buf.String()
	return nil
}

// GeneratePDF generate pdf function
func (r *Pdf) GeneratePDF(pdfPath string) (bool, error) {
	t := time.Now().Unix()

	err1 := ioutil.WriteFile(getTempHtmlPath(t), []byte(r.body), 0644)
	if err1 != nil {
		panic(err1)
	}

	htmlFile, err := os.Open(getTempHtmlPath(t))
	if htmlFile != nil {
		defer htmlFile.Close()
	}
	if err != nil {
		log.Fatal(err)
	}

	pdfGenerator, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		removeTempHtml(getTempHtmlPath(t))
		log.Fatal(err)
	}

	pageReader := wkhtmltopdf.NewPageReader(htmlFile)
	pageReader.PageOptions.EnableLocalFileAccess.Set(true)

	pdfGenerator.AddPage(pageReader)
	pdfGenerator.PageSize.Set(wkhtmltopdf.PageSizeA4)
	pdfGenerator.Dpi.Set(300)

	err = pdfGenerator.Create()
	if err != nil {
		log.Fatal(err)
	}

	err = pdfGenerator.WriteFile(pdfPath)
	if err != nil {
		log.Fatal(err)
	}

	removeTempHtml(getTempHtmlPath(t))

	return true, nil
}

func getTempHtmlPath(t int64) string {
	return "storage/" + strconv.FormatInt(int64(t), 10) + ".html"
}

func removeTempHtml(path string) {
	if err := os.Remove(path); err != nil {
		log.Fatal(err)
	}
}
