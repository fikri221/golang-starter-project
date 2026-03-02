package user

import (
	"jwt-auth/internal/types"
	"jwt-auth/internal/utils"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.UserStore
}

// NewHandler creates a new user handler
func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	// Get user data from request body
	var payload types.RegisterUserRequest
	err := utils.ParseJSON(r, &payload)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	// Validate payload
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, "Invalid request body, "+errors.Error())
		return
	}

	// Check if user already exists
	_, err = h.store.GetUserByEmail(payload.Email)
	if err == nil {
		utils.WriteError(w, http.StatusBadRequest, "User already exists")
		return
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(payload.Password)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "Error hashing password")
		return
	}

	// Create user in database
	err = h.store.CreateUser(types.User{
		Email:     payload.Email,
		Password:  hashedPassword,
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
	})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "Error creating user")
		return
	}

	utils.WriteJSON(w, http.StatusCreated, map[string]string{"message": "User created successfully"})
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	// Get user data from request body
	var payload types.LoginUserRequest
	err := utils.ParseJSON(r, &payload)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	// Validate user data
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, "Invalid request body, "+errors.Error())
		return
	}

	// Check if user exists
	u, err := h.store.GetUserByEmail(payload.Email)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Invalid email or password")
		return
	}

	// Check password
	if !utils.ComparePassword(u.Password, []byte(payload.Password)) {
		utils.WriteError(w, http.StatusBadRequest, "Invalid email or password")
		return
	}

	// Here we would typically generate a JWT token
	// For now, let's just return success
	utils.WriteJSON(w, http.StatusOK, map[string]string{"token": ""})
}
