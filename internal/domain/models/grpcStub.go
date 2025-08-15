package models

type GrpcStub struct {
	ID           int    `json:"id" db:"id"`
	Name         string `json:"name" db:"name"`
	ProjectId    string `json:"projectId" db:"project_id"`
	CreatedAt    string `json:"created_at" db:"created_at"`
	ProtoFile    string `json:"proto_file" db:"proto_file"`
	ProtoMethod  string `json:"proto_method" db:"proto_method"`
	ResponseBody string `json:"response_body" db:"response_body"`
}
