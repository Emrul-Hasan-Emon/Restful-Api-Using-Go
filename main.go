package main

import (
	"log"
	"net/http"
	"restfulapi/api"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// Define the api endpoint for hompeage
	router.HandleFunc("/", api.GetItems).Methods("GET")

	// Define the api endpoint for getting all items list
	router.HandleFunc("/items", api.GetItems).Methods("GET")

	// Define the api endpoint for getting specific item
	router.HandleFunc("/items/{id}", api.GetItem).Methods("GET")

	// Define the api endpoint for creating a item
	router.HandleFunc("/items", api.CreateItem).Methods("POST")

	// Define the api endpoint for upadating a item details
	router.HandleFunc("/items/{id}", api.UpdateItem).Methods("PUT")

	// Define the api endpoint for deleting a item
	router.HandleFunc("/items/{id}", api.Deleteitem).Methods("DELETE")

	// Start the http Server
	log.Fatal(http.ListenAndServe(":8080", router))
}
