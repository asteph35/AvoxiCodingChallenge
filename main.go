package main

import (
	"AvoxiCodingChallenge/repository"
	"AvoxiCodingChallenge/service"
	"fmt"
	"log"
	"net/http"
)

func main() {
	s := service.New(repository.New())
	r := service.Router(s)

	fmt.Println("Starting server on the port 8080...")

	log.Fatal(http.ListenAndServe(":8080", r))
}
