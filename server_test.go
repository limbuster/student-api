package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPing(t *testing.T) {
	t.Run("GET /ping should return `ping`", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/ping", nil)
		response := httptest.NewRecorder()

		s := NewStudentServer()
		s.ServeHTTP(response, request)

		got := response.Body.String()
		want := "ping"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("GET /pong should return `pong`", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/pong", nil)
		response := httptest.NewRecorder()

		s := NewStudentServer()
		s.ServeHTTP(response, request)

		got := response.Body.String()
		want := "pong"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
