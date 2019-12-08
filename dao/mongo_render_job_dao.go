package dao

import (
	"context"

	log "github.com/sirupsen/logrus"
	"github.com/tsubro/document-renderer/models"
	"github.com/tsubro/document-renderer/utils"
	"go.mongodb.org/mongo-driver/bson"
)

func Insert(job *models.RenderJob) {

	collection := utils.GetMongoConnection().Database("document-renderer").Collection("render-job")

	insertResult, err := collection.InsertOne(context.TODO(), job)
	if err != nil {
		log.Fatal(err)
	}

	log.Info("Inserted a single document: ", insertResult.InsertedID)
}

func Get(jobId string) (*models.RenderJob , error) {
	collection := utils.GetMongoConnection().Database("document-renderer").Collection("render-job")

	var job models.RenderJob

	err := collection.FindOne(context.TODO(), bson.D{{"jobid", jobId}}).Decode(&job)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	log.Info("Fetched a single document succefully")
	return &job, nil
}

func Update(job *models.RenderJob) {
	collection := utils.GetMongoConnection().Database("document-renderer").Collection("render-job")

	update := bson.M{
		"$set": bson.M{
		  "state": job.State,
		  "inputs": job.Inputs,
		},
	  }
	_, err := collection.UpdateOne(context.TODO(), bson.D{{"jobid", job.JobId}}, update)
	if err != nil {
		log.Error(err)
	}
}
