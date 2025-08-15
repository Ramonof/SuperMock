package project

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

type Saver interface {
	SaveProject(
		ctx context.Context,
		stub models.Project,
	) (int64, error)
}

type Provider interface {
	GetByIdProject(
		ctx context.Context,
		stubId int,
	) (models.Project, error)
	GetAllProjects(
		ctx context.Context,
	) ([]models.Project, error)
}

type Updater interface {
	UpdateProject(
		ctx context.Context,
		stub models.Project,
	) (int64, error)
}

type Deleter interface {
	DeleteProject(
		ctx context.Context,
		projectId string,
		stubId string,
	) (int64, error)
}

type Project struct {
	log      *slog.Logger
	saver    Saver
	provider Provider
	updater  Updater
	deleter  Deleter
}

func New(log *slog.Logger, saver Saver, provider Provider, updater Updater, deleter Deleter) *Project {
	return &Project{log, saver, provider, updater, deleter}
}

func (service *Project) GetAll(w http.ResponseWriter, r *http.Request) {
	var stubs []models.Project
	stubs, err := service.provider.GetAllProjects(context.TODO())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(stubs)
}

func (service *Project) GetById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["project_id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	stub, err := service.provider.GetByIdProject(context.TODO(), id)
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

func (service *Project) Create(w http.ResponseWriter, r *http.Request) {
	var stub models.Project
	err := json.NewDecoder(r.Body).Decode(&stub)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = service.saver.SaveProject(context.TODO(), stub)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	stub.ID = 0
	stub.CreatedAt = "now" // Placeholder
	json.NewEncoder(w).Encode(stub)
}

func (service *Project) Update(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	var stub models.Project
	err := json.NewDecoder(r.Body).Decode(&stub)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = service.updater.UpdateProject(context.TODO(), stub)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	stub.ID, _ = strconv.Atoi(id)
	stub.CreatedAt = "now" // Placeholder
	json.NewEncoder(w).Encode(stub)
}

func (service *Project) DeleteProject(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	projectId := params["project_id"]

	_, err := service.deleter.DeleteProject(context.TODO(), projectId, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
