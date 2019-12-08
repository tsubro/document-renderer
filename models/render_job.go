package models

type RenderJob struct {
	JobId string `json:"job_id"`

	State string 			`json:"state"`

	Inputs []Input 			`json:"inputs"`

	OutputLocation string   `json:"output_location"`
	OutputHeader   []Header `json:"output_header"`
	OutputMimeType string   `json:"output_mime_type"`

	Links []Selflinks 		`json:"links"`
}

type Input struct {
	InputLocation string   	`json:"input_location"`
	InputHeader   []Header 	`json:"input_header"`
	InputMimeType string   	`json:"input_mime_type"`
	InputSettings settings 	`json:"input_settings"`
	TaskStatus string 		`json:"task_status"`
}

type Header struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type settings struct {
	Orientation string `json:"orientation"`
	Size        string `json:"size"`
	Scale       string `json:"scale"`
	Colour      bool   `json:"color"`
	PageRange   string `json:"page_range"`
}

type Selflinks struct {
	Link   string `json:"link"`
	Method string `json:"method"`
}
