package api

import (
	"database/sql"
	"jwt-auth/internal/service/user"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

// NewAPIServer creates a new API server
func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

// Run starts the API server
func (s *APIServer) Run() error {
	router := mux.NewRouter()

	// Group API routes under /api/v1
	subRouter := router.PathPrefix("/api/v1").Subrouter()

	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subRouter)

	log.Println("Starting server on ", s.addr)

	return http.ListenAndServe(s.addr, router)
}
