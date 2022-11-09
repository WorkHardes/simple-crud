package users

import (
	"net/http"

	"github.com/gorilla/mux"
)

func InitRouter(r *mux.Router) {
	usersRouter := r.PathPrefix("/users").Subrouter()

	usersRouter.HandleFunc("/list", IndexUsers).Methods("GET")
}

func IndexUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(201)
}
