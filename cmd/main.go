package main

import (
	"fmt"
	"log"
	"net/http"
	"sync/atomic"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	// Create a new router using the Gorilla Mux package
	router := mux.NewRouter()

	// Create an atomic variable to track the readiness status
	isReady := &atomic.Value{}
	isReady.Store(false)

	// Start a goroutine that will set the readiness status to true after a delay
	go func() {
		time.Sleep(5 * time.Second)
		isReady.Store(true)
		log.Printf("Ready....")
	}()

	// Define a handler function for the /readyz endpoint
	readyzHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	// Add the /readyz endpoint to the router and specify that it only accepts GET requests
	router.HandleFunc("/readyz", readyzHandler).Methods("GET")

	// Add middleware function to the router that checks the readiness status before each request
	router.Use(readyz(isReady))

	// Define routes for the /genpass and /healthz endpoints
	router.HandleFunc("/genpass", generatePasswords).Methods("POST")
	router.HandleFunc("/healthz", healthz)

	// Start the HTTP server and listen for incoming requests on port 8000
	fmt.Println("Server started ....")
	http.ListenAndServe(":8000", router)
}
