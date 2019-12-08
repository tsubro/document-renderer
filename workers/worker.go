package workers

import (
	log "github.com/sirupsen/logrus"
	"github.com/tsubro/document-renderer/dao"
	"github.com/tsubro/document-renderer/models"
	"github.com/tsubro/document-renderer/utils"
	
	"sync"
)

var renderJobInterface dao.RenderJobInterface = dao.MongoImpl{}

var imageMimeTypes = []string{"image/gif", "image/jpeg"}
var officeMimeTypes = []string{"application/docx", "application/xlsx"}

func TaskDeligator(job *models.RenderJob) *models.RenderJob {

	job.JobId = utils.GetUUID()
	job.State = "processing"
	

	go renderJobInterface.CreateJob(job)
	go metaTaskDelegator(job)
	
	return job
}

func metaTaskDelegator(job *models.RenderJob) {
	var wg sync.WaitGroup
	m := make(map[int][]byte)
	task := func(index int, job *models.RenderJob, b []byte, m map[int][]byte) {

		input := job.Inputs[index]
		var task Task
		if utils.Contains(imageMimeTypes, input.InputMimeType) {
			task = ImageTaskImpl{}
		} else if "application/pdf" == input.InputMimeType {
			task = PdfTaskImpl{}
		} else if utils.Contains(officeMimeTypes, input.InputMimeType) {
			task = OfficeTaskImpl{}
		} else {
			log.Info("Not my cake bro!!!")
		}
		task.Transform(index, job, b, m)
	}

	for i:=0; i<len(job.Inputs); i++ {
		wg.Add(1)
		go utils.Download(i, job, task, &wg, m)
	}
	wg.Wait()
	utils.SetJobState(job)
	Collect(job, m)

	go renderJobInterface.UpdateJob(job)	
}

func GetJobStatus(id string) (*models.RenderJob, error) {
	var renderJobInterface dao.RenderJobInterface
	renderJobInterface = dao.MongoImpl{}
	return renderJobInterface.GetJob(id)
}
