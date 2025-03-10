package models

type RestStub struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	ProjectId    string `json:"projectId"`
	CreatedAt    string `json:"created_at"`
	Path         string `json:"path"`
	ResponseBody string `json:"response_body"`
}
