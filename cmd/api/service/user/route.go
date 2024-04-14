package user

import (
	"net/http"

	"github.com/gorilla/mux" // Importing Gorilla Mux for routing
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoute(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST") // Handling login endpoint
	router.HandleFunc("/register", h.handleRegister).Methods("POST") // Handling register endpoint
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	// Handling login request
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	// Handling register request
}