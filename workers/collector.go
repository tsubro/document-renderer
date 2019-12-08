package workers

import (
	"io"
	// "io/ioutil"

	log "github.com/sirupsen/logrus"
	"github.com/tsubro/document-renderer/models"
	"github.com/tsubro/document-renderer/utils"

	"bytes"

	"github.com/jung-kurt/gofpdf"
	"github.com/jung-kurt/gofpdf/contrib/gofpdi"
	unipdf "github.com/unidoc/unidoc/pdf/model"
)

func Collect(job *models.RenderJob, m map[int][]byte) {
	log.Info("Inside Collector")
	merge(job.JobId, m)
}

func merge(jobId string, m map[int][]byte) {
	log.Info("Merging all pdfs")

	pdf := gofpdf.New("P", "mm", "A4", "")
	defer pdf.OutputFileAndClose(utils.GetOutputDir(jobId)+"/output.pdf")

	for i:=0; i<len(m); i++ {
		b := m[i]
		
		if len(b) == 0 {
			continue
		}
		var buff io.ReadSeeker = bytes.NewReader(b)

		//Getting total number of pages
		pdfReader, err := unipdf.NewPdfReader(buff)
		if err != nil {
			log.Error("Failed to read PDF file: %v\n", err)
		}
		numPages, err := pdfReader.GetNumPages()
		log.Info("index ", "bytes ", "pageno ", i, len(b), numPages)
		if err != nil {
			log.Error("Failed to get number of pages: %v\n", err)
		}

		//Building PDF
		for i := 1; i <= numPages; i++ {
			tp := gofpdi.ImportPageFromStream(pdf, &buff, i, "/MediaBox")
			pdf.AddPage()
			gofpdi.UseImportedTemplate(pdf, tp, 20, 50, 150, 0)
		}
	}
}

