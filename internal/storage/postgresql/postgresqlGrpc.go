package postgresql

import (
	"SuperStub/internal/domain/models"
	"context"
	"fmt"
	"github.com/jhump/protoreflect/desc/protoparse"
	"log"
	"os"
	"strconv"
)

func (storage *Storage) SaveGrpcStub(ctx context.Context, stub models.GrpcStub) (int64, error) {
	const op = "storage.postgres.SaveStub"

	_, err := storage.db.NamedExec(
		"INSERT INTO grpc_stubs (name, project_id, created_at, path, response_body) VALUES (:name, :project_id, :created_at, :path, :response_body)",
		&stub,
	)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return 0, nil
}

func (storage *Storage) GetGrpcStubByFullPath(ctx context.Context, path string) (models.GrpcStub, error) {
	const op = "storage.postgres.GetGrpcStubByFullPath"

	var newStub models.GrpcStub
	err := storage.db.QueryRow(
		"SELECT id, name, created_at, path, response_body FROM grpc_stubs WHERE path = $1",
		path,
	).Scan(&newStub.ID, &newStub.Name, &newStub.CreatedAt, &newStub.Path, &newStub.ResponseBody)
	if err != nil {
		return models.GrpcStub{}, err
	}
	return newStub, nil
}

func (storage *Storage) GrpcStub(ctx context.Context, projectId string, stubId string) (models.GrpcStub, error) {
	const op = "storage.postgres.Stub"

	var newStub models.GrpcStub
	err := storage.db.QueryRow(
		"SELECT id, name, created_at, path, response_body FROM grpc_stubs WHERE id = $1 AND project_id = $2",
		stubId, projectId,
	).Scan(&newStub.ID, &newStub.Name, &newStub.CreatedAt, &newStub.Path, &newStub.ResponseBody)
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

func (storage *Storage) GetProtoByName(ctx context.Context, name string) (string, error) {
	const op = "storage.postgres.GetProtoByName"

	var res string
	err := storage.db.QueryRow(
		"SELECT proto_file FROM grpc_protos WHERE service_name = $1",
		name,
	).Scan(&res)
	if err != nil {
		return "", err
	}
	return res, nil
}

func (storage *Storage) SaveProto(ctx context.Context, projectId string, fileName string) (int64, error) {
	const op = "storage.postgres.SaveProto"

	protoFiles := []string{fileName}
	parser := protoparse.Parser{}
	fileDescriptors, err := parser.ParseFiles(protoFiles...)
	if err != nil {
		log.Fatalf("Failed to parse proto files: %v", err)
	}
	packageName := fileDescriptors[0].GetPackage()
	id, err := strconv.Atoi(projectId)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	data, err := os.ReadFile(fileName)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	proto := models.GrpcProto{
		Id:          0,
		ServiceName: packageName,
		ProjectId:   id,
		CreatedAt:   "now",
		ProtoFile:   string(data),
	}
	_, err = storage.db.NamedExec(
		"INSERT INTO grpc_protos (service_name, project_id, created_at, proto_file) VALUES (:service_name, :project_id, :created_at, :proto_file)",
		&proto,
	)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return 0, nil
}
