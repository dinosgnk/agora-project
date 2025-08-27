package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

func LoadConfig[T any]() *T {
	environment := os.Getenv("ENVIRONMENT")
	if environment == "" {
		currentDir, _ := os.Getwd()
		configPath := filepath.Join(currentDir, ".env.local")
		if _, err := os.Stat(configPath); err == nil {
			err := godotenv.Load(configPath)
			if err != nil {
				log.Printf("cannot load configuration file, error: %v", err)
			}
		}
	}

	var cfg T
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("cannot parse env variables, error: %v", err)
	}

	return &cfg
}
