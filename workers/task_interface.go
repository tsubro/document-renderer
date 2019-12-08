package workers

import "github.com/tsubro/document-renderer/models"

type Task interface {
	Transform(index int, job *models.RenderJob, b []byte, m map[int][]byte,dependentTask ...interface{})
}

type PdfTaskImpl struct {
}

func (t PdfTaskImpl) Transform(index int, job *models.RenderJob, b []byte, m map[int][]byte, dependentTask ...interface{}) {
	TransformPdf(index, job, b, m)
}

type ImageTaskImpl struct {
}

func (t ImageTaskImpl) Transform(index int, job *models.RenderJob, b []byte, m map[int][]byte, dependentTask ...interface{}) {
	TransformImage(index, job, b, m)
}

type OfficeTaskImpl struct {
}

func (t OfficeTaskImpl) Transform(index int, job *models.RenderJob, b []byte, m map[int][]byte, dependentTask ...interface{}) {
	TransformOffice(index, job, b, m)
}
