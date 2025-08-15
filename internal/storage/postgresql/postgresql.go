package postgresql

import (
	"SuperStub/internal/domain/models"
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type Storage struct {
	db *sqlx.DB
}

func New(psqlInfo string) (*Storage, error) {
	const op = "storage.postgresql.New"

	db, err := sqlx.Connect("postgres", psqlInfo)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to database!")

	return &Storage{db: db}, nil
}

func (storage *Storage) Close() error {
	err := storage.db.Close()
	if err != nil {
		return err
	}
	return nil
}
func (storage *Storage) SaveStub(ctx context.Context, stub models.RestStub) (int64, error) {
	const op = "storage.postgres.SaveStub"

	_, err := storage.db.NamedExec(
		"INSERT INTO reststubs (name, project_id, created_at, path, response_body) VALUES (:name, :project_id, :created_at, :path, :response_body)",
		&stub,
	)
	if err != nil {
		//var sqlErr pq.Error
		//
		//if errors.As(err, &sqlErr) && sqlErr.ExtendedCode == pq.ErrConstraintUnique {
		//	return 0, fmt.Errorf("%s: %w", op, storage.ErrUserExists)
		//}

		return 0, fmt.Errorf("%s: %w", op, err)
	}

	//id, err := res.LastInsertId()
	//if err != nil {
	//	return 0, fmt.Errorf("%s: %w", op, err)
	//}

	return 0, nil
}

func (storage *Storage) Stub(ctx context.Context, projectId string, stubId string) (models.RestStub, error) {
	const op = "storage.postgres.Stub"

	var newStub models.RestStub
	err := storage.db.QueryRow(
		"SELECT id, name, created_at, path, response_body FROM reststubs WHERE id = $1 AND project_id = $2",
		stubId, projectId,
	).Scan(&newStub.ID, &newStub.Name, &newStub.CreatedAt, &newStub.Path, &newStub.ResponseBody)
	if err != nil {
		return models.RestStub{}, err
	}
	return newStub, nil
}

func (storage *Storage) Stubs(ctx context.Context, projectId string) ([]models.RestStub, error) {
	const op = "storage.postgres.createUser"
	var stubs []models.RestStub

	err := storage.db.Select(&stubs, "SELECT * FROM reststubs WHERE project_id = $1", projectId)

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return stubs, nil
}

func (storage *Storage) UpdateStub(ctx context.Context, stub models.RestStub) (int64, error) {
	//TODO
	return 0, nil
}

func (storage *Storage) DeleteStub(ctx context.Context, projectId string, stubId string) (int64, error) {
	//TODO
	return 0, nil
}
