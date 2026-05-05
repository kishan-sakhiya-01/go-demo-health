package main

import (
	"fmt"
	"log"
	"net/http"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK")
}

func main() {
	http.HandleFunc("/health", healthHandler)
	fmt.Println("Server starting on port 5999")
	if err := http.ListenAndServe(":5999", nil); err != nil {
		log.Fatal(err)
	}
}
