package dao

import (
	"github.com/tsubro/document-renderer/models"
)

type RenderJobInterface interface {
	CreateJob(job *models.RenderJob)
	GetJob(jobId string) (*models.RenderJob, error)
	UpdateJob(job *models.RenderJob)
	// Delete()
}

type MongoImpl struct {
}

func (mo MongoImpl) CreateJob(job *models.RenderJob) {
	Insert(job)
}

func (mo MongoImpl) GetJob(jobId string) (*models.RenderJob, error) {
	return Get(jobId)
}

func (mo MongoImpl) UpdateJob(job *models.RenderJob) {
	Update(job)
}
