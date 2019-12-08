package utils

import (
	"context"

	log "github.com/sirupsen/logrus"
	"github.com/tsubro/document-renderer/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetMongoConnection() *mongo.Client {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	log.Info("Connected to MongoDB!")
	return client
}

func SetTaskState(job *models.RenderJob, index int, taskState string) {
	job.Inputs[index].TaskStatus = taskState
}

func SetJobState(job *models.RenderJob, jobState ...interface{}) {

	if jobState != nil {
		for _,val := range jobState {
			job.State = val.(string)
			return
		}
	} else {
		for _, val := range job.Inputs {
			if val.TaskStatus == "DONE" {
				job.State = "DONE"
				return
			}
		}
		job.State = "FAILED"
	}
}

