// Copyright (c) 2023, Renzo Beltran
// Use of this source code is governed by a MIT License
// license that can be found in the LICENSE file.

package routes

import (
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestPublicRoutes(t *testing.T) {
	tests := []struct {
		Description   string
		Route         string
		ExpectedError bool
		ExpectedCode  int
	}{
		{
			Description:   "Health Check",
			Route:         "/api/v1/health",
			ExpectedError: false,
			ExpectedCode:  200,
		},
	}

	// Log suppression when testing with `go test ./... -v`
	log.SetOutput(io.Discard)

	// Define Fiber app.
	app := fiber.New()

	// Define routes.
	PublicRoutes(app)

	// Iterate through test single cases
	for _, test := range tests {
		// Create a new http request with the route from the test case.
		req := httptest.NewRequest("GET", test.Route, http.NoBody)
		req.Header.Set("Content-Type", "application/json")

		// Perform the request plain with the app.
		res, err := app.Test(req, -1) //the -1 disables request latency

		// Verify that no error occurred that is not expected
		assert.Equalf(t, test.ExpectedError, err != nil, test.Description)

		// As expected errors lead to broken request,
		// the next test case needs to be processed.

		if test.ExpectedError {
			continue
		}

		// Verify if the status code is a expected.
		assert.Equalf(t, test.ExpectedCode, res.StatusCode, test.Description)
	}
}
