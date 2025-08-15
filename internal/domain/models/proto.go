package models

type GrpcProto struct {
	Id          int    `json:"id" db:"id"`
	ServiceName string `json:"service_name" db:"service_name"`
	ProjectId   int    `json:"project_id" db:"project_id"`
	CreatedAt   string `json:"created_at" db:"created_at"`
	ProtoFile   string `json:"proto_file" db:"proto_file"`
}
