package workers

import (
	"bytes"
	"io"
	"io/ioutil"

	"github.com/jung-kurt/gofpdf"
	"github.com/jung-kurt/gofpdf/contrib/gofpdi"
	unipdf "github.com/unidoc/unidoc/pdf/model"

	log "github.com/sirupsen/logrus"
	"github.com/tsubro/document-renderer/models"
	"github.com/tsubro/document-renderer/utils"
)

func TransformPdf(index int, job *models.RenderJob, b []byte, m map[int][]byte, dependentTask ...interface{}) {
	log.Info("Inside pdf task!!")

	var buff io.ReadSeeker = bytes.NewReader(b)
	pdf := gofpdf.New("P", "mm", "A4", "")

	//Getting total number of pages
	pdfReader, err := unipdf.NewPdfReader(buff)
	if err != nil {
		log.Error("Failed to read PDF file: %v\n", err)
		utils.SetTaskState(job, index, "FAILED")
		return
	}
	numPages, err := pdfReader.GetNumPages()

	if err != nil {
		log.Error("Failed to get number of pages: %v\n", err)
		utils.SetTaskState(job, index, "FAILED")
		return
	}
	
	//Building PDF
	for i := 1; i <= numPages; i++ {
		tp := gofpdi.ImportPageFromStream(pdf, &buff, i, "/MediaBox")
		pdf.AddPage()
		gofpdi.UseImportedTemplate(pdf, tp, 20, 50, 150, 0)
	}

	var buf bytes.Buffer
	err = pdf.Output(&buf)
	m[index] = buf.Bytes()
	ioutil.WriteFile(utils.SaveOption(job.JobId, index), m[index], 0644)
	utils.SetTaskState(job, index, "DONE")
}
