package grpc

import (
	"SuperStub/internal/domain/models"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"log/slog"
	"net/http"
	"strconv"
)

type StubSaver interface {
	SaveGrpcStub(
		ctx context.Context,
		stub models.GrpcStub,
	) (int64, error)
}

type StubProvider interface {
	GrpcStub(
		ctx context.Context,
		projectId string,
		stubId string,
	) (models.GrpcStub, error)
	GrpcStubs(
		ctx context.Context,
		projectId string,
	) ([]models.GrpcStub, error)
}

type StubUpdater interface {
	UpdateGrpcStub(
		ctx context.Context,
		stub models.GrpcStub,
	) (int64, error)
}

type StubDeleter interface {
	DeleteGrpcStub(
		ctx context.Context,
		projectId string,
		stubId string,
	) (int64, error)
}

type Grpc struct {
	log          *slog.Logger
	stubSaver    StubSaver
	stubProvider StubProvider
	stubUpdater  StubUpdater
	stubDeleter  StubDeleter
}

func New(log *slog.Logger, stubSaver StubSaver, stubProvider StubProvider, updater StubUpdater, deleter StubDeleter) *Grpc {
	return &Grpc{log, stubSaver, stubProvider, updater, deleter}
}

func (service *Grpc) GetAllGrpcStubs(w http.ResponseWriter, r *http.Request) {
	var stubs []models.GrpcStub
	params := mux.Vars(r)
	projectId := params["project_id"]
	stubs, err := service.stubProvider.GrpcStubs(context.TODO(), projectId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(stubs)
}

func (service *Grpc) GetGrpcStubById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	projectId := params["project_id"]
	stub, err := service.stubProvider.GrpcStub(context.TODO(), projectId, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.NotFound(w, r)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	json.NewEncoder(w).Encode(stub)
}

func (service *Grpc) CreateGrpcStub(w http.ResponseWriter, r *http.Request) {
	var stub models.GrpcStub
	params := mux.Vars(r)
	projectId := params["project_id"]
	err := json.NewDecoder(r.Body).Decode(&stub)
	stub.ProjectId = projectId

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = service.stubSaver.SaveGrpcStub(context.TODO(), stub)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	stub.ID = 0
	stub.CreatedAt = "now" // Placeholder
	json.NewEncoder(w).Encode(stub)
}

func (service *Grpc) UpdateGrpcStub(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	projectId := params["project_id"]
	var stub models.GrpcStub
	err := json.NewDecoder(r.Body).Decode(&stub)
	stub.ProjectId = projectId
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = service.stubUpdater.UpdateGrpcStub(context.TODO(), stub)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	stub.ID, _ = strconv.Atoi(id)
	stub.CreatedAt = "now" // Placeholder
	json.NewEncoder(w).Encode(stub)
}

func (service *Grpc) DeleteGrpcStub(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	projectId := params["project_id"]

	_, err := service.stubDeleter.DeleteGrpcStub(context.TODO(), projectId, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

//func (service *Grpc) ServeStub(w http.ResponseWriter, r *http.Request) {
//	var stubs []models.GrpcStub
//	params := mux.Vars(r)
//	projectId := params["project_id"]
//	path := params["path"]
//	stubs, err := service.stubProvider.Stubs(context.TODO(), projectId)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//
//	for _, stub := range stubs {
//		if stub.Path == path {
//			var jsonMap map[string]interface{}
//			json.Unmarshal([]byte(stub.ResponseBody), &jsonMap)
//			json.NewEncoder(w).Encode(jsonMap)
//			return
//		}
//	}
//
//	http.Error(w, "No such stub", http.StatusBadRequest)
//}
