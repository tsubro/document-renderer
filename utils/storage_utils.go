package utils

import (
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
	"bytes"

	log "github.com/sirupsen/logrus"
	"github.com/tsubro/document-renderer/models"
)

func Download(index int, job *models.RenderJob, a func(index int, job *models.RenderJob, b []byte, m map[int][]byte),
	wg *sync.WaitGroup, m map[int][]byte) {

	defer wg.Done()
	input := job.Inputs[index]

	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}

	client := &http.Client{Transport: tr}
	req, _ := http.NewRequest("GET", input.InputLocation, nil)

	for _, val := range input.InputHeader {
		req.Header.Add(val.Name, val.Value)
	}

	resp, _ := client.Do(req)
	defer resp.Body.Close()

	log.Info(resp.StatusCode)
	if resp.StatusCode == 200 {
		bodyByte, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			log.Info("Error in downloading document")
			SetTaskState(job, index, "FAILED")
			return
		} 
		
		a(index, job, bodyByte, m)
	}  else {
		SetTaskState(job, index, "FAILED")
	}
}

func Upload(job *models.RenderJob, b []byte) {
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}

	bodyByte := bytes.NewReader(b)

	client := &http.Client{Transport: tr}
	req, _ := http.NewRequest("POST", job.OutputLocation, bodyByte)
	for _, val := range job.OutputHeader {
		req.Header.Add(val.Name, val.Value)
	}

	resp, _ := client.Do(req)
	defer resp.Body.Close()

	log.Info(resp.StatusCode)
	if resp.StatusCode != 201 && resp.StatusCode != 200 {
		SetJobState(job, "FAILED")
	}
}

func SaveOption(jobId string, index int) string {
	filePath := GetOutputDir(jobId)
	os.MkdirAll(filePath, os.ModePerm)
	fileName := filePath + "/" + jobId + "#" + strconv.Itoa(index) + ".pdf"
	return fileName
}

func GetOutputDir(jobId string) string {
	return "/Users/schakraborty/Downloads/" + jobId
}
