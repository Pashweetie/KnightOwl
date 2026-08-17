// Package httpserver owns HTTP routes and the server shutdown lifecycle.
package httpserver

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"sync"
	"time"

	"github.com/Pashweetie/KnightOwl/apps/server/internal/config"
)

// Server is the HTTP part of KnightOwl. Later tickets add routes here.
type Server struct {
	settings config.HTTP
	logger   *slog.Logger

	// mu is an RWMutex — Go's reader-writer lock. Multiple goroutines handle
	// requests concurrently, and they all read "started" and "stopping" on every
	// health check. Without a lock, a read could race with a write during
	// shutdown. RWMutex allows unlimited concurrent readers (RLock) but exclusive
	// access for writers (Lock). Go's race detector will crash the program if
	// you skip this, which is how Go enforces thread safety at development time.
	// https://pkg.go.dev/sync#RWMutex
	mu       sync.RWMutex
	started  bool
	stopping bool
	http     *http.Server
}

// New creates the server. The & returns a pointer (memory address) to the
// struct so every caller shares one instance. Without it, Go copies the struct
// by value and changes to one copy wouldn't be visible to others.
// https://go.dev/tour/moretypes/1
func New(settings config.HTTP, logger *slog.Logger) *Server {
	if logger == nil {
		logger = slog.Default()
	}
	return &Server{settings: settings, logger: logger}
}

// Handler registers the routes this server responds to.
func (server *Server) Handler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health/live", server.live)
	mux.HandleFunc("GET /health/startup", server.startup)
	mux.HandleFunc("GET /health/ready", server.ready)
	mux.HandleFunc("POST /admin/drain", server.drain)
	return mux
}

// Run starts the HTTP server and blocks until shutdown completes.
func (server *Server) Run(ctx context.Context) error {
	httpServer := &http.Server{
		Addr:              server.settings.ListenAddress,
		Handler:           server.Handler(),
		ReadHeaderTimeout: 5 * time.Second,
	}

	server.mu.Lock()
	server.started = true
	server.http = httpServer
	server.mu.Unlock()

	// "go" launches a goroutine — a lightweight concurrent thread managed by
	// Go's runtime (~2KB stack, not an OS thread). This one watches for the
	// shutdown signal while ListenAndServe blocks below. It's the standard Go
	// pattern for "do something when the context is cancelled."
	// https://go.dev/tour/concurrency/1
	go server.waitForShutdownSignal(ctx)

	server.logger.Info("server listening", "address", server.settings.ListenAddress)
	err := httpServer.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		return nil
	}
	return err
}

// waitForShutdownSignal blocks on a channel receive (<-ctx.Done()). Channels
// are Go's typed pipes between goroutines. <-ch means "wait until something is
// sent into ch." ctx.Done() returns a channel that closes when the context is
// cancelled (signal received), unblocking this line.
// https://go.dev/tour/concurrency/2
func (server *Server) waitForShutdownSignal(ctx context.Context) {
	<-ctx.Done()
	server.stop()
}

func (server *Server) live(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}

func (server *Server) startup(w http.ResponseWriter, _ *http.Request) {
	server.mu.RLock()
	started := server.started
	server.mu.RUnlock()
	if !started {
		http.Error(w, "starting", http.StatusServiceUnavailable)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (server *Server) ready(w http.ResponseWriter, _ *http.Request) {
	server.mu.RLock()
	ready := server.started && !server.stopping
	server.mu.RUnlock()
	if !ready {
		http.Error(w, "not ready", http.StatusServiceUnavailable)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// drain handles POST /admin/drain. "Drain" is infrastructure jargon for
// "graceful removal from the load balancer pool" — stop accepting new requests,
// let in-flight ones finish, then exit.
func (server *Server) drain(w http.ResponseWriter, _ *http.Request) {
	server.markStopping()
	w.WriteHeader(http.StatusNoContent)
	go server.stop()
}

func (server *Server) markStopping() {
	server.mu.Lock()
	server.stopping = true
	server.mu.Unlock()
}

func (server *Server) stop() {
	server.markStopping()
	server.mu.RLock()
	httpServer := server.http
	server.mu.RUnlock()
	if httpServer == nil {
		return
	}
	timeout := time.Duration(server.settings.ShutdownTimeoutSeconds) * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	// "defer" schedules a function call to run when the enclosing function
	// returns — regardless of whether it returns normally or early via error.
	// It's Go's equivalent of try/finally. Used here to release the timer.
	// https://go.dev/tour/flowcontrol/12
	defer cancel()
	_ = httpServer.Shutdown(ctx)
}
