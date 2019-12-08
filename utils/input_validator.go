package utils

import (

	log "github.com/sirupsen/logrus"
	"net/url"
	"github.com/tsubro/document-renderer/models"
	"sync"
)

var wg sync.WaitGroup
var inputMimeTypes = [...]string{"application/pdf", "image/jpeg", "image/gif", "application/docx", "application/xlsx"}

func ValidateInputPayload(job *models.RenderJob)  []string {
	log.Info("Inside Validate Input Payload")
	validationErrorsChannel := make(chan string, 100)

	validationErrors := []string{}

	wg.Add(1)
	go validateInputDetails(job, validationErrorsChannel)
	wg.Add(1)
	go validateOutputDetails(job, validationErrorsChannel)
	wg.Add(1)
	go validateInputSettings(job, validationErrorsChannel)
	
	wg.Wait()
	close(validationErrorsChannel)

	for val := range validationErrorsChannel {
		validationErrors = append(validationErrors, val)
	}
	
	return validationErrors
}

//Input detail validation
func validateInputDetails(job *models.RenderJob, validationErrorsChannel chan<- string) {
	defer wg.Done()
	for _, input := range job.Inputs {
		
		if len(input.InputLocation) == 0 {
			validationErrorsChannel <- "Input Location cannot be empty"
		} else {
			_, err := url.ParseRequestURI(input.InputLocation)
			if err != nil {
				validationErrorsChannel <- "Input Location is malformed"
			}
		}
	
		if len(input.InputMimeType) == 0 {
			validationErrorsChannel <- "Input MimeType cannot be empty"
		} else {
			found := false
			for _, val := range inputMimeTypes {
				if input.InputMimeType == val {
					found = true
					break
				}
			}
	
			if found == false {
				validationErrorsChannel <- "Unsupported Input MimeType"
			}
		}
	}
}

//Output details validation
func validateOutputDetails(job *models.RenderJob, validationErrorsChannel chan<- string) {
	defer wg.Done()
	if len(job.OutputLocation) == 0 {
		validationErrorsChannel <- "Output Location cannot be empty"
	} else {
		_, err := url.ParseRequestURI(job.OutputLocation)
		if err != nil {
			validationErrorsChannel <- "Output Location is malformed"
		}
	}

	if len(job.OutputMimeType) == 0 {
		validationErrorsChannel <- "Output MimeType cannot be empty"
	} else if job.OutputMimeType != "application/pdf" {
		validationErrorsChannel <- "Unsupported Output MimeType"
	}
}

func validateInputSettings(job *models.RenderJob, validationErrors chan <- string) {
	defer wg.Done()
}
