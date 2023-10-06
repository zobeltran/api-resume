// Copyright (c) 2023, Renzo Beltran
// Use of this source code is governed by a MIT License
// license that can be found in the LICENSE file.

package config

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

// Fiber Configuration
func FiberConfig() fiber.Config {
	// Define Server Settings
	readTimeoutSecondsCount, _ := strconv.Atoi("5")

	// Return Fiber Configurations
	return fiber.Config{
		ReadTimeout: time.Second * time.Duration(readTimeoutSecondsCount),
	}

}
