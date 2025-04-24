// main.go
package main

import (
	"log"      // For logging server startup and fatal errors
	"net/http" // For starting the HTTP server

	"github.com/akarkareddy/ethereum-multisig-wallet/api" // Import the custom API package where handlers and routes are defined

	"github.com/gorilla/mux" // Gorilla Mux is a powerful router for HTTP request routing(Rest Api)
)

func main() {
	router := mux.NewRouter() // Create a new router using Gorilla Mux to handle different HTTP routes
	api.SetupRoutes(router)   // Call a helper function from the API package to set or register all end points

	log.Println("Server started on http://localhost:8080") // Log to console that the server has started and is listening on port 8080
	log.Fatal(http.ListenAndServe(":8080", router))        // If the server crashes or can't start, log.Fatal will print the error and exit the program
}
