package api

import (
	"database/sql"
	"jwt-auth/internal/auth"
	"jwt-auth/internal/config"
	"jwt-auth/internal/service/user"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	db  *sql.DB
	cfg *config.Config
}

// NewAPIServer creates a new API server
func NewAPIServer(db *sql.DB, cfg *config.Config) *APIServer {
	return &APIServer{
		db:  db,
		cfg: cfg,
	}
}

// Run starts the API server
func (s *APIServer) Run() error {
	router := mux.NewRouter()

	// Group API routes under /api/v1
	subRouter := router.PathPrefix("/api/v1").Subrouter()

	userStore := user.NewStore(s.db)
	jwtService := auth.NewJWT(s.cfg)
	userHandler := user.NewHandler(userStore, jwtService)
	userHandler.RegisterRoutes(subRouter)

	log.Println("Starting server on :", s.cfg.Port)

	return http.ListenAndServe(":"+s.cfg.Port, router)
}
