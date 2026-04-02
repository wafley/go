package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"school/internal/service"
)

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)

	json.NewEncoder(w).Encode(map[string]string{
		"error": "Page not found",
	})
}

func main() {
	service.InitLogger()

	http.HandleFunc("/students", service.GetStudents)
	http.HandleFunc("/", notFound)

	println("The server is running at http://localhost:8080")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        fmt.Println("Failed to start server:", err)
    }
}