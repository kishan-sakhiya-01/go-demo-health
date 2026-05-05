package main

import (
	"fmt"
	"net/http"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK")
}

func main() {
	http.HandleFunc("/health", healthHandler)
	fmt.Println("Server starting on port 5999")
	http.ListenAndServe(":5999", nil)
}
