// Copyright (c) 2023, Renzo Beltran
// Use of this source code is governed by a MIT License
// license that can be found in the LICENSE file.

package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/zobeltran/api-resume/internal/config"
	"github.com/zobeltran/api-resume/internal/http/rest/routes"
	"github.com/zobeltran/api-resume/internal/middlewares"
	"github.com/zobeltran/api-resume/internal/utils"

	"github.com/sirupsen/logrus"
)

func main() {

	// Read Config
	conf, err := config.ReadConfig()
	if err != nil {
		logrus.Fatal(err)
	}

	logrus.Infof("Initializing API Resume Service (%s)", conf.Environment)
	// Define Fiber Config.
	fiberConfig := config.FiberConfig()

	// Define a new Fiber app with config.
	app := fiber.New(fiberConfig)

	// Middleware.
	middlewares.FiberMiddleware(app)

	// Routes.
	routes.PublicRoutes(app)

	// Start Server (with or without graceful shutdown)
	if os.Getenv("STAGE") == "Dev" {
		utils.StartServer(app, conf.Server)
	} else {
		utils.StartServerWithGracefulShutdown(app, conf.Server)
	}
}
