// Copyright (c) 2023, Renzo Beltran
// Use of this source code is governed by a MIT License
// license that can be found in the LICENSE file.

package utils

import (
	"log"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
	"github.com/zobeltran/api-resume/internal/config"
)

// StartServerWithGracefulShutdown function for starting server with graceful shutdown.
func StartServerWithGracefulShutdown(a *fiber.App, conf config.ServerConfig) {
	// Create channel for idle connections
	idleConnsClosed := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt) // Catch OS Signals.
		<-sigint

		// Received an interrupt signal, shutdown.
		if err := a.Shutdown(); err != nil {
			// Error from closing listeners, or context timeout:
			log.Printf("Server is not shutting down! Reason %v", err)
		}

		close(idleConnsClosed)
	}()

	// Build Fiber connection URL.
	fiberConnUrl, _ := ConnectionBuilder("fiber", conf)

	// Run Server.
	if err := a.Listen(fiberConnUrl); err != nil {
		log.Printf("Server is not running! Reason: %v", err)
	}
}

// StartServer func for starting a simple server.
func StartServer(a *fiber.App, conf config.ServerConfig) {
	// Build Fiber connection URL
	fiberConnUrl, _ := ConnectionBuilder("fiber", conf)

	if err := a.Listen(fiberConnUrl); err != nil {
		log.Printf("Server is not running! Reason: %v", err)
	}
}
