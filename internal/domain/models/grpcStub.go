package models

type GrpcStub struct {
	ID           int    `json:"id" db:"id"`
	Name         string `json:"name" db:"name"`
	ProjectId    int    `json:"project_id" db:"project_id"`
	CreatedAt    string `json:"created_at" db:"created_at"`
	ServiceId    int    `json:"service_id" db:"service_id"`
	Method       string `json:"method" db:"method"`
	ResponseBody string `json:"response_body" db:"response_body"`
}
