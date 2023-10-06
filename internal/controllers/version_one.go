// Copyright (c) 2023, Renzo Beltran
// Use of this source code is governed by a MIT License
// license that can be found in the LICENSE file.

package controllers

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func HealthCheck(c *fiber.Ctx) error {
	response := fiber.Map{
		"status": "OK",
	}

	c.Status(fiber.StatusOK)
	log.Println(c.Locals("requestid"))
	return c.JSON(response)
}
