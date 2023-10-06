// Copyright (c) 2023, Renzo Beltran
// Use of this source code is governed by a MIT License
// license that can be found in the LICENSE file.

package utils

import (
	"fmt"
	"log"

	"github.com/zobeltran/api-resume/internal/config"
)

// ConnectionBuilder func for building connections.
func ConnectionBuilder(n string, config config.ServerConfig) (string, error) {
	// Define URL to connection
	var url string

	switch n {
	case "fiber":
		// URL for Fiber Connection.
		url = fmt.Sprintf(
			"%s:%s",
			config.ServerHost,
			config.ServerPort,
		)
		log.Printf(url)
	default:
		// Return error message.
		return "", fmt.Errorf("connection name '%v' is not supported", n)
	}

	return url, nil
}
