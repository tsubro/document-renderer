package utils

import (
	"encoding/json"
	"net/http"

	"github.com/tsubro/document-renderer/models"
)

func BuildCreateJobResponse(job *models.RenderJob, w http.ResponseWriter) {

	getJobLinks := models.Selflinks{"http://localhost:8080/job-api/" + job.JobId, "GET"}
	cancelJobLinks := models.Selflinks{"http://localhost:8080/job-api/" + job.JobId, "DELETE"}

	job.Links = append(job.Links, getJobLinks)
	job.Links = append(job.Links, cancelJobLinks)

	js, _ := json.Marshal(job)
	w.WriteHeader(201)
	w.Write(js)
}

func BuildErrorResponse(validationErrors  []string, w http.ResponseWriter) {
	
	errorResponse := struct {
		Errors []string `json:"errors"`
	}{
		validationErrors,
	}
	js, _ := json.Marshal(errorResponse)
	w.WriteHeader(400)
	w.Write(js)
}

