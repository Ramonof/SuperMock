package postgresql

import (
	"SuperStub/internal/domain/models"
	"context"
	"fmt"
)

func (storage *Storage) SaveProject(ctx context.Context, stub models.Project) (int64, error) {
	const op = "storage.postgres.SaveStub"

	_, err := storage.db.NamedExec(
		"INSERT INTO projects (name, created_at) VALUES (:name, :created_at)",
		&stub,
	)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return 0, nil
}

func (storage *Storage) GetByIdProject(ctx context.Context, stubId int) (models.Project, error) {
	const op = "storage.postgres.Stub"

	var newStub models.Project
	err := storage.db.QueryRow(
		"SELECT id, name, created_at FROM projects WHERE id = $1",
		stubId,
	).Scan(&newStub.ID, &newStub.Name, &newStub.CreatedAt)
	if err != nil {
		return models.Project{}, err
	}
	return newStub, nil
}

func (storage *Storage) GetAllProjects(ctx context.Context) ([]models.Project, error) {
	const op = "storage.postgres.GrpcStubs"
	var stubs []models.Project

	err := storage.db.Select(&stubs, "SELECT * FROM projects")

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return stubs, nil
}

func (storage *Storage) UpdateProject(ctx context.Context, stub models.Project) (int64, error) {
	//TODO
	return 0, nil
}

func (storage *Storage) DeleteProject(ctx context.Context, projectId string, stubId string) (int64, error) {
	//TODO
	return 0, nil
}
