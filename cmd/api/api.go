package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux" // Importing Gorilla Mux for routing
	"github/nitin3339/ecom/cmd/api/service/user" // Importing user service package
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter() // Creating a new router using Gorilla Mux
	subrouter := router.PathPrefix("/api/v1").Subrouter() // Creating a subrouter for API versioning
	userHandler := user.NewHandler() // Creating a new instance of the user handler
	userHandler.RegisterRoute(subrouter) // Registering user routes with the subrouter
	log.Println("Listening on", s.addr) // Logging server address
	return http.ListenAndServe(s.addr, router) // Starting HTTP server and passing router as handler
}
