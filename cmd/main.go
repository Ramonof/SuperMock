package main

import (
	"SuperStub/internal/services/rest"
	"SuperStub/internal/storage/postgresql"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres_user"
	password = "postgres_password"
	dbname   = "postgres"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	storage, err := postgresql.New(psqlInfo)
	if err != nil {
		panic(err)
	}

	restService := rest.New(nil, storage, storage, storage, storage)

	router := mux.NewRouter()

	router.HandleFunc("/{project_id}/stub", restService.GetAllRestStubs).Methods("GET")
	router.HandleFunc("/{project_id}/stub/{id}", restService.GetRestStubById).Methods("GET")
	router.HandleFunc("/{project_id}/stub", restService.CreateRestStub).Methods("POST")
	router.HandleFunc("/{project_id}/stub/{id}", restService.UpdateRestStub).Methods("PUT")
	router.HandleFunc("/{project_id}/stub/{id}", restService.DeleteRestStub).Methods("DELETE")

	router.HandleFunc("/{project_id}/{path}", restService.ServeStub).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))
}
