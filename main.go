package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	server := NewStudentServer()
	if err := http.ListenAndServe(":"+port, server); err != nil {
		log.Fatalf("could not listen on port %d %v", port, err)
	}
}
