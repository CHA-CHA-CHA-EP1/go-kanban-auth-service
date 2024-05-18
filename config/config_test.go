package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigValidation(t *testing.T) {
	config := &Config{
		Env: "dev",
		App: App{Name: "TestApp"},
		Server: Server{
			Port:    8080,
			Timeout: 10,
		},
		Redis: Redis{
			Host:     "localhost",
			Port:     6379,
			DB:       0,
			Password: "securepassword",
		},
	}

	err := config.Validate()
	assert.Nil(t, err, "Config should be valid")
}

func TestConfigValidationFails(t *testing.T) {
	config := &Config{
		Env: "dev",
		App: App{Name: ""},
		Server: Server{
			Port:    0,
			Timeout: 0,
		},
		Redis: Redis{
			Host:     "invalid-host",
			Port:     70000, // Invalid port
			DB:       -1,    // Invalid DB index
			Password: "",
		},
	}

	err := config.Validate()
	assert.NotNil(t, err, "Config should be invalid due to incorrect values")
}
