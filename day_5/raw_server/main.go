package main

import (
	"encoding/json"
	"net/http"
)

// response structure (optional but clean)
type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// set header
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// response body
	res := Response{
		Success: true,
		Message: "Server is running on port 3001",
	}

	json.NewEncoder(w).Encode(res)
}

func main() {
	http.HandleFunc("/", homeHandler)

	// start server
	http.ListenAndServe(":3002", nil)
}
