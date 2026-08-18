package main

import (
	"context"
	"errors"
	"flag"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/Pashweetie/KnightOwl/apps/server/internal/config"
	"github.com/Pashweetie/KnightOwl/apps/server/internal/httpserver"
)

func main() {
	configPath := flag.String("config", "config/development.json", "path to the JSON configuration file")
	flag.Parse()
	settings, err := config.Load(*configPath)
	if err != nil {
		slog.Error("invalid configuration", "error", err.Error())
		os.Exit(1)
	}

	// We create the server object first, THEN set up signal handling below.
	// This order matters: the server doesn't open a port until Run() is called,
	// so nothing is listening yet. We're just preparing the struct with its
	// settings. The signal context (ctx) gets passed INTO Run(), where it
	// actually starts listening and watching for shutdown at the same time.
	app := httpserver.New(settings.HTTP, slog.Default())

	// signal.NotifyContext does two things:
	//
	// 1. Creates a "context" (ctx) — think of this as a cancellation token that
	//    gets flipped to "done" when the OS sends SIGINT (Ctrl+C) or SIGTERM
	//    (Docker/K8s telling the container to stop). Any code holding this ctx
	//    can watch it and react when it flips.
	//
	// 2. Returns a cleanup function (stop). Internally, Go allocated an OS-level
	//    signal channel to intercept those signals. If we never call stop(), that
	//    channel leaks — the OS keeps routing signals to it even after we don't
	//    care anymore. "defer stop()" means "call stop() automatically when
	//    main() exits, no matter how it exits (normal return, early error, etc)."
	//
	// KEY DISTINCTION: stop() does NOT cause the shutdown. The ctx getting
	// cancelled is what causes the shutdown (see waitForShutdownSignal in
	// server.go). stop() just tells Go "I'm done intercepting signals, free
	// that internal plumbing." It's housekeeping, not a trigger.
	//
	// "defer" in Go = "run this when the function returns." Same idea as
	// try/finally in Python or a destructor — guarantees cleanup happens.
	// https://go.dev/tour/flowcontrol/12
	// https://pkg.go.dev/os/signal#NotifyContext
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	if err := app.Run(ctx); err != nil && !errors.Is(err, context.Canceled) {
		slog.Error("server stopped", "error", err.Error())
		os.Exit(1)
	}
}
