// Copyright (c) 2023, Renzo Beltran
// Use of this source code is governed by a MIT License
// license that can be found in the LICENSE file.

package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

// A config struct that serves houses the service configs from the config file
type Config struct {
	Environment string       `yaml:"environment"`
	LogLevel    string       `yaml:"log_level"`
	Server      ServerConfig `yaml:"server"`
}

type ServerConfig struct {
	ServerHost string `yaml:"server_host"`
	ServerPort string `yaml:"server_port"`
}

func ReadConfig() (config Config, err error) {
	configFile, err := os.ReadFile("./config/config.yml")
	if err != nil {
		return
	}
	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		return
	}
	return
}
