package workers

import (
	"bytes"
	"io/ioutil"

	"github.com/jung-kurt/gofpdf"
	log "github.com/sirupsen/logrus"
	"github.com/tsubro/document-renderer/models"
	"github.com/tsubro/document-renderer/utils"
)

func TransformImage(index int, job *models.RenderJob, b []byte, m map[int][]byte, dependentTask ...interface{}) {
	log.Info("Inside Image task!!")

	buff := bytes.NewReader(b)
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	tp := pdf.ImageTypeFromMime("image/jpeg")

	infoPtr := pdf.RegisterImageReader("bird", tp, buff)
	if pdf.Ok() {
		imgWd, imgHt := infoPtr.Extent()
		pdf.Image("bird", (210-imgWd)/2.0, pdf.GetY()+5,
			imgWd, imgHt, false, tp, 0, "")
	}

	var buf bytes.Buffer
	err := pdf.Output(&buf)
	m[index] = buf.Bytes()

	err = ioutil.WriteFile(utils.SaveOption(job.JobId, index), m[index], 0644)
	if err != nil {
		log.Error(err)
		utils.SetTaskState(job, index, "FAILED")
		return
	}
	utils.SetTaskState(job, index, "DONE")
}
