package service

import (
	"github.com/gorilla/mux"
)

// Router is exported and used in main.go
func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/book/{id}", CheckIP).Methods("GET", "OPTIONS")

	return router
}