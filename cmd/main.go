package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Create a new router using the Gorilla Mux package
	router := mux.NewRouter()

	// Define routes for the /genpass and /healthz endpoints
	router.HandleFunc("/genpass", generatePasswords).Methods("POST")

	// Start the HTTP server and listen for incoming requests on port 8000
	fmt.Println("Server started ....")
	http.ListenAndServe(":8000", router)
}
