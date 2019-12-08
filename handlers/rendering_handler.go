package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/tsubro/document-renderer/models"
	"github.com/tsubro/document-renderer/utils"
	"github.com/tsubro/document-renderer/workers"
)

func CreateJob(w http.ResponseWriter, r *http.Request) {
	log.Info("Called HTTP Method", r.Method)

	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()

	job := models.RenderJob{}
	err := d.Decode(&job)
	if err != nil {
		log.Error("Failed to parse request body ", err)
		utils.BuildErrorResponse([]string{err.Error()}, w)
		return
	}

	validationErrors := utils.ValidateInputPayload(&job)

	w.Header().Set("Content-Type", "application/json")
	if validationErrors != nil && len(validationErrors) > 0 {
		log.Error(validationErrors)
		utils.BuildErrorResponse(validationErrors, w)
		return
	}

	workers.TaskDeligator(&job)
	utils.BuildCreateJobResponse(&job, w)
	return
}

func GetJob(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id := params["id"]
	job, err := workers.GetJobStatus(id)

	if err != nil {
		log.Error(err)

		errorResponse := struct {
			error `json:"errors"`
		}{
			err,
		}
		js, _ := json.Marshal(errorResponse)
		w.WriteHeader(500)
		w.Write(js)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	js, _ := json.Marshal(job)
	w.WriteHeader(200)
	w.Write(js)
}
