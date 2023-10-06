// Copyright (c) 2023, Renzo Beltran
// Use of this source code is governed by a MIT License
// license that can be found in the LICENSE file.

package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zobeltran/api-resume/internal/controllers"
)

// PublicRoutes func for decribing group of public routes
func PublicRoutes(a *fiber.App) {
	// Create route group
	route := a.Group("/api/v1")

	// Routes for GET method
	route.Get("/health", controllers.HealthCheck)
}
