package models

type RestStub struct {
	ID           int    `json:"id" db:"id"`
	Name         string `json:"name" db:"name"`
	ProjectId    string `json:"projectId" db:"project_id"`
	CreatedAt    string `json:"created_at" db:"created_at"`
	Path         string `json:"path" db:"path"`
	ResponseBody string `json:"response_body" db:"response_body"`
}
