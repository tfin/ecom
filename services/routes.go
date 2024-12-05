package user

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/users", h.handelRegister).Methods("GET")
	// router.HandleFunc("/register", h.CreateUser).Methods("POST")
	// router.HandleFunc("/users/{id}", h.GetUser).Methods("GET")
	// router.HandleFunc("/users/{id}", h.UpdateUser).Methods("PUT")
	// router.HandleFunc("/users/{id}", h.DeleteUser).Methods("DELETE")
	// router.HandleFunc("/users", h.GetUsers).Methods("GET")

}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	// TODO
}

func (h *Handler) handelRegister(w http.ResponseWriter, r *http.Request) {
	// TODO
}
