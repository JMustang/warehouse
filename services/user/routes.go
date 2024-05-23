package user

import (
	"encoding/json"
	"net/http"

	"github.com/JMustang/warehouse/types"
	"github.com/gorilla/mux"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {

}
func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	// get JSON payload
	if r.Body == nil {

	}
	var payload types.RegisterUserPayload
	err := json.NewDecoder(r.Body).Decode(payload)
	// check if the user already exists
	// create a new user
	// return JWT
}
