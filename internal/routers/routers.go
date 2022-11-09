package routers

import (
	"example.com/simple-crud/internal/routers/users"
	"github.com/gorilla/mux"
)

func New() *mux.Router {
	router := mux.NewRouter()
	apiRouter := router.PathPrefix("/api/v1").Subrouter()
	apiRouter.HandleFunc("/list", users.IndexUsers).Methods("GET")
	users.InitRouter(apiRouter)

	return router
}
