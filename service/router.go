package service

import (
	"github.com/gorilla/mux"
)

// Router is exported and used in main.go
func Router(s Service) *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/ip-check", s.CheckIP).Methods("GET")
	return router
}
