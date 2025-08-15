package postgresql

import (
	"SuperStub/internal/domain/models"
	"context"
	"fmt"
)

func (storage *Storage) SaveGrpcStub(ctx context.Context, stub models.GrpcStub) (int64, error) {
	const op = "storage.postgres.SaveStub"

	_, err := storage.db.NamedExec(
		"INSERT INTO grpc_stubs (name, project_id, created_at, service_id, method, response_body) VALUES (:name, :project_id, :created_at, :service_id,:method, :response_body)",
		&stub,
	)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return 0, nil
}

func (storage *Storage) GrpcStub(ctx context.Context, projectId string, stubId string) (models.GrpcStub, error) {
	const op = "storage.postgres.Stub"

	var newStub models.GrpcStub
	err := storage.db.QueryRow(
		"SELECT id, name, created_at, service_id, method, response_body FROM grpc_stubs WHERE id = $1 AND project_id = $2",
		stubId, projectId,
	).Scan(&newStub.ID, &newStub.Name, &newStub.CreatedAt, &newStub.ServiceId, &newStub.Method, &newStub.ResponseBody)
	if err != nil {
		return models.GrpcStub{}, err
	}
	return newStub, nil
}

func (storage *Storage) GrpcStubs(ctx context.Context, projectId string) ([]models.GrpcStub, error) {
	const op = "storage.postgres.GrpcStubs"
	var stubs []models.GrpcStub

	err := storage.db.Select(&stubs, "SELECT * FROM grpc_stubs WHERE project_id = $1", projectId)

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return stubs, nil
}

func (storage *Storage) UpdateGrpcStub(ctx context.Context, stub models.GrpcStub) (int64, error) {
	//TODO
	return 0, nil
}

func (storage *Storage) DeleteGrpcStub(ctx context.Context, projectId string, stubId string) (int64, error) {
	//TODO
	return 0, nil
}
