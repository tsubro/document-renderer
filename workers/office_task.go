package workers

import (
	"io/ioutil"
	"os/exec"
	"os"
	"runtime"

	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/tsubro/document-renderer/models"
	"github.com/tsubro/document-renderer/utils"
)

func TransformOffice(index int, job *models.RenderJob, b []byte, m map[int][]byte,dependentTask ...interface{}) {
	log.Info("Inside Office Task")

	var arg0 string
	var arg1 string

	if runtime.GOOS == "darwin" {
		log.Info("Mac OS Detected")
		arg0 = "/Applications/LibreOffice.app/Contents/MacOS/soffice"
		arg1 = "--headless" //This command is optional, it will help to disable the splash screen of LibreOffice.

	} else {
		log.Info("Linux OS Detected")
		arg0 = "lowriter"
		arg1 = "--invisible" //This command is optional, it will help to disable the splash screen of LibreOffice.

	}

	arg2 := "--convert-to"
	arg3 := "pdf:writer_pdf_Export"
	arg4 := "--outdir"

	outFileName := utils.SaveOption(job.JobId, index)
	inFileName, err := writeInputFile(outFileName, b, job.Inputs[index].InputMimeType)
	if err != nil {
		log.Error(err)
		utils.SetTaskState(job, index, "FAILED")
		return
	}
	defer os.Remove(inFileName)

	_, err = exec.Command(arg0, arg1, arg2, arg3, arg4, utils.GetOutputDir(job.JobId) ,inFileName).Output()
	log.Info("office byte array ", len(b))
	if err != nil {
		log.Error(err)
		utils.SetTaskState(job, index, "FAILED")
		return
	}
	b, err = ioutil.ReadFile(outFileName)
	m[index] = b
	utils.SetTaskState(job, index, "DONE")
}

func writeInputFile(fileName string, b []byte, mimeType string) (string, error) {
	log.Info("Writing Input Office File ")

	inFileName := strings.Split(fileName, ".")[0] + "." + strings.Split(mimeType, "/")[1]
	err := ioutil.WriteFile(inFileName, b, 0644)
	
	return inFileName, err
}

