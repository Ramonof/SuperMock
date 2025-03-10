package main

import (
	"SuperStub/internal/services/rest"
	"SuperStub/internal/storage/postgresql"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type RestStub struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	ProjectId string `json:"projectId"`
	CreatedAt string `json:"created_at"`
}

var db *sql.DB

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres_user"
	password = "postgres_password"
	dbname   = "postgres"
)

func main() {
	// Initialize database connection
	//psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	//	"password=%s dbname=%s sslmode=disable",
	//	host, port, user, password, dbname)
	//var err error
	//db, err = sql.Open("postgres", psqlInfo)
	//if err != nil {
	//	panic(err)
	//}
	//defer db.Close()
	//
	//err = db.Ping()
	//if err != nil {
	//	panic(err)
	//}

	//fmt.Println("Successfully connected!")

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	storage, err := postgresql.New(psqlInfo)
	if err != nil {
		panic(err)
	}
	db = storage.Db

	restService := rest.New(nil, storage, storage, storage, storage)

	// Initialize router
	router := mux.NewRouter()

	// Define API routes
	//router.HandleFunc("/{project_id}/stub", getRestStub).Methods("GET")            // Fetch all users
	//router.HandleFunc("/{project_id}/stub/{id}", getRestStubById).Methods("GET")   // Fetch a user by ID
	//router.HandleFunc("/{project_id}/stub", createRestStub).Methods("POST")        // Create a new user
	//router.HandleFunc("/{project_id}/stub/{id}", updateRestStub).Methods("PUT")    // Update a user by ID
	//router.HandleFunc("/{project_id}/stub/{id}", deleteRestStub).Methods("DELETE") // Delete a user by ID

	router.HandleFunc("/{project_id}/stub", restService.GetAllRestStubs).Methods("GET")        // Fetch all users
	router.HandleFunc("/{project_id}/stub/{id}", restService.GetRestStubById).Methods("GET")   // Fetch a user by ID
	router.HandleFunc("/{project_id}/stub", restService.CreateRestStub).Methods("POST")        // Create a new user
	router.HandleFunc("/{project_id}/stub/{id}", restService.UpdateRestStub).Methods("PUT")    // Update a user by ID
	router.HandleFunc("/{project_id}/stub/{id}", restService.DeleteRestStub).Methods("DELETE") // Delete a user by ID

	router.HandleFunc("/{project_id}/{path}", restService.ServeStub).Methods("GET") // Delete a user by ID

	// Start server on port 8000
	log.Fatal(http.ListenAndServe(":8000", router))
}

func getRestStub(w http.ResponseWriter, r *http.Request) {
	var stubs []RestStub
	params := mux.Vars(r)
	projectId := params["project_id"]
	rows, err := db.Query("SELECT * FROM reststub WHERE project_id = $1", projectId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var stub RestStub
		if err := rows.Scan(&stub.ID, &stub.Name, &stub.ProjectId, &stub.CreatedAt); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		stubs = append(stubs, stub)
	}
	json.NewEncoder(w).Encode(stubs)
}

func getRestStubById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	projectId := params["project_id"]
	var stub RestStub
	err := db.QueryRow("SELECT id, name, created_at FROM reststub WHERE id = $1 AND project_id = $2", id, projectId).Scan(&stub.ID, &stub.Name, &stub.CreatedAt)
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

func createRestStub(w http.ResponseWriter, r *http.Request) {
	var stub RestStub
	params := mux.Vars(r)
	projectId := params["project_id"]
	err := json.NewDecoder(r.Body).Decode(&stub)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = db.Exec("INSERT INTO reststub (name, project_id, created_at) VALUES ($1, $2, $3)", stub.Name, projectId, "now")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//id, err := result.LastInsertId()
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}
	stub.ID = 0
	stub.CreatedAt = "now" // Placeholder
	json.NewEncoder(w).Encode(stub)
}

func updateRestStub(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	projectId := params["project_id"]
	var user RestStub
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = db.Exec("UPDATE reststub SET name = $1 WHERE id = $2 AND project_id = $3", user.Name, id, projectId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user.ID, _ = strconv.Atoi(id)
	user.CreatedAt = "now" // Placeholder
	json.NewEncoder(w).Encode(user)
}

func deleteRestStub(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	projectId := params["project_id"]

	_, err := db.Exec("DELETE FROM reststub WHERE id = $1 AND project_id = $2", id, projectId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
