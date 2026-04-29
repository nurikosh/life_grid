package config

import (
	"fmt"
	"os"
	"strings"
)

type Config struct {
	DatabaseURL string
	Port        string
}

func Load() (*Config, error) {
	databaseURL, ok := os.LookupEnv("DATABASE_URL")
	if !ok || databaseURL == "" {
		return nil, fmt.Errorf("DATABASE_URL is required")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	} else if !strings.Contains(port, ":") {
		port = ":" + port
	}

	return &Config{
		DatabaseURL: databaseURL,
		Port:        port,
	}, nil
}
