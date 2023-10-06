// Copyright (c) 2023, Renzo Beltran
// Use of this source code is governed by a MIT License
// license that can be found in the LICENSE file.

package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func FiberMiddleware(a *fiber.App) {
	a.Use(
		// Add RequestId
		requestid.New(),
	)
}
