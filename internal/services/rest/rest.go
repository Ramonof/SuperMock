package rest

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
	"strings"
)

type StubSaver interface {
	SaveStub(
		ctx context.Context,
		stub models.RestStub,
	) (int64, error)
}

type StubProvider interface {
	Stub(
		ctx context.Context,
		projectId string,
		stubId string,
	) (models.RestStub, error)
	Stubs(
		ctx context.Context,
		projectId string,
	) ([]models.RestStub, error)
}

type StubUpdater interface {
	UpdateStub(
		ctx context.Context,
		stub models.RestStub,
	) (int64, error)
}

type StubDeleter interface {
	DeleteStub(
		ctx context.Context,
		projectId string,
		stubId string,
	) (int64, error)
}

type Rest struct {
	log          *slog.Logger
	stubSaver    StubSaver
	stubProvider StubProvider
	stubUpdater  StubUpdater
	stubDeleter  StubDeleter
}

func New(log *slog.Logger, stubSaver StubSaver, stubProvider StubProvider, updater StubUpdater, deleter StubDeleter) *Rest {
	return &Rest{log, stubSaver, stubProvider, updater, deleter}
}

func (service *Rest) GetAllRestStubs(w http.ResponseWriter, r *http.Request) {
	var stubs []models.RestStub
	params := mux.Vars(r)
	projectId := params["project_id"]
	stubs, err := service.stubProvider.Stubs(context.TODO(), projectId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(stubs)
}

func (service *Rest) GetRestStubById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	projectId := params["project_id"]
	stub, err := service.stubProvider.Stub(context.TODO(), projectId, id)
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

func (service *Rest) CreateRestStub(w http.ResponseWriter, r *http.Request) {
	var stub models.RestStub
	params := mux.Vars(r)
	projectId := params["project_id"]
	err := json.NewDecoder(r.Body).Decode(&stub)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	stub.ProjectId, err = strconv.Atoi(projectId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if !strings.HasPrefix(stub.Path, "/") {
		stub.Path = "/" + stub.Path
	}
	_, err = service.stubSaver.SaveStub(context.TODO(), stub)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	stub.ID = 0
	stub.CreatedAt = "now" // Placeholder
	json.NewEncoder(w).Encode(stub)
}

func (service *Rest) UpdateRestStub(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	projectId := params["project_id"]
	stubId := params["id"]
	var stub models.RestStub
	err := json.NewDecoder(r.Body).Decode(&stub)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	stub.ProjectId, err = strconv.Atoi(projectId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	stub.ID, err = strconv.Atoi(stubId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if !strings.HasPrefix(stub.Path, "/") {
		stub.Path = "/" + stub.Path
	}
	_, err = service.stubUpdater.UpdateStub(context.TODO(), stub)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	stub.ID, _ = strconv.Atoi(id)
	stub.CreatedAt = "now" // Placeholder
	json.NewEncoder(w).Encode(stub)
}

func (service *Rest) DeleteRestStub(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	projectId := params["project_id"]

	_, err := service.stubDeleter.DeleteStub(context.TODO(), projectId, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (service *Rest) ServeStub(w http.ResponseWriter, r *http.Request) {
	pathVars := strings.SplitN(r.RequestURI, "/", 4)
	if len(pathVars) != 4 || pathVars[1] != "projects" {
		http.Error(w, "404 page not found", http.StatusNotFound)
		return
	}

	var stubs []models.RestStub
	method := r.Method
	projectId := pathVars[2]
	path := "/" + pathVars[3]

	stubs, err := service.stubProvider.Stubs(context.TODO(), projectId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for _, stub := range stubs {
		if stub.Path == path && (stub.Method == "ANY" || stub.Method == method) {
			var jsonMap map[string]interface{}
			json.Unmarshal([]byte(stub.ResponseBody), &jsonMap)
			json.NewEncoder(w).Encode(jsonMap)
			return
		}
	}

	http.Error(w, "No such stub", http.StatusBadRequest)
}
