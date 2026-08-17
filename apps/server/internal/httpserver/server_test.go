package httpserver

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Pashweetie/KnightOwl/apps/server/internal/config"
)

func TestHealthLifecycle(t *testing.T) {
	server := New(config.HTTP{ListenAddress: "127.0.0.1:8787", ShutdownTimeoutSeconds: 1}, nil)
	response := httptest.NewRecorder()
	server.Handler().ServeHTTP(response, httptest.NewRequest(http.MethodGet, "/health/live", nil))
	if response.Code != http.StatusNoContent {
		t.Fatalf("live: got %d", response.Code)
	}
	server.started = true
	response = httptest.NewRecorder()
	server.Handler().ServeHTTP(response, httptest.NewRequest(http.MethodPost, "/admin/drain", nil))
	response = httptest.NewRecorder()
	server.Handler().ServeHTTP(response, httptest.NewRequest(http.MethodGet, "/health/ready", nil))
	if response.Code != http.StatusServiceUnavailable {
		t.Fatalf("ready: got %d", response.Code)
	}
}
