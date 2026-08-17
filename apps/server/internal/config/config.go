// Package config loads the server settings from a JSON file.
package config

import (
	"encoding/json"
	"fmt"
	"os"
)

// Config is the full configuration file. Each subsystem gets its own section.
type Config struct {
	HTTP HTTP `json:"http"`
}

// HTTP contains the settings needed to run the HTTP server.
type HTTP struct {
	// ListenAddress is "host:port" — the network interface and port the server
	// binds to. Go lets you pick ANY valid address:port combo, there's nothing
	// special about 8080 or any other port. We use 127.0.0.1:8787 because:
	//
	//   "127.0.0.1" = loopback only — rejects connections from other machines.
	//       This is a security choice: during development, only YOUR computer
	//       can talk to it. In production, Cloudflare Tunnel handles external
	//       access instead of exposing the port directly.
	//
	//   "8787" = deliberately not 8080/3000/5000 to avoid collisions with other
	//       dev tools you might be running. Totally arbitrary — change it to
	//       anything from 1024-65535 in development.json if something else
	//       wants this port.
	//
	// If you changed it to "0.0.0.0:8787", any device on your network could
	// connect. If you used ":8787" (empty host), Go defaults to 0.0.0.0 too.
	ListenAddress string `json:"listenAddress"`

	// ShutdownTimeoutSeconds is how long the server waits for in-flight requests
	// to finish when shutting down, before forcefully closing connections.
	ShutdownTimeoutSeconds int `json:"shutdownTimeoutSeconds"`
}

// Load reads and validates one JSON configuration file.
func Load(path string) (Config, error) {
	contents, err := os.ReadFile(path)
	if err != nil {
		return Config{}, fmt.Errorf("read configuration: %w", err)
	}
	var result Config
	if err := json.Unmarshal(contents, &result); err != nil {
		return Config{}, fmt.Errorf("parse configuration: %w", err)
	}
	if result.HTTP.ListenAddress == "" {
		return Config{}, fmt.Errorf("http.listenAddress is required")
	}
	if result.HTTP.ShutdownTimeoutSeconds < 1 || result.HTTP.ShutdownTimeoutSeconds > 300 {
		return Config{}, fmt.Errorf("http.shutdownTimeoutSeconds must be from 1 through 300")
	}
	return result, nil
}
