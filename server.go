package main

import (
	"fmt"
	"net/http"
)

// StudentServer returns pong or pong
type StudentServer struct {
	http.Handler
}

// NewStudentServer returns a new instance of StudentServer
func NewStudentServer() *StudentServer {
	server := new(StudentServer)
	router := http.NewServeMux()
	router.Handle("/ping", http.HandlerFunc(handlePing))
	router.Handle("/pong", http.HandlerFunc(handlePong))
	server.Handler = router
	return server
}

func handlePing(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ping")
}

func handlePong(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "pong")
}
