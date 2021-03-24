package router

import (
	handler "github.com/Fuerback/go-crud/api/handlers"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/useraccount/{id}", handler.GetUser).Methods("GET", "OPTIONS")

	return router
}
