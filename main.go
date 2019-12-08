package main

import (
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	r "github.com/tsubro/document-renderer/handlers"
)

const port = ":8080"
const create_job_url = "/job-api"

func main() {
	log.Info("Starting HTTP Server on the PORT ", port)

	httpMux := mux.NewRouter()
	httpMux.HandleFunc("/job-api", r.CreateJob).Methods("POST")
	httpMux.HandleFunc("/job-api/{id}", r.GetJob).Methods("GET")
	http.ListenAndServe(port, httpMux)
}
