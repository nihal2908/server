package main

import (
	"fmt"
	"net/http"
)

// handler for the root endpoint
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"message": "Hello, World!" first change. second change. third change. forth change.}`)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprint(w, "Method not allowed")
	}
}

func main() {
	// Register the handler for the root path
	http.HandleFunc("/", helloHandler)

	// Start the server on port 8080
	fmt.Println("Server is running on http://localhost:3000")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
