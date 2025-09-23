package config

import (
	"os"
	"path/filepath"

	"github.com/caarlos0/env/v11"
	"github.com/dinosgnk/agora-project/internal/pkg/logger"
	"github.com/joho/godotenv"
)

func LoadConfig[T any](log logger.Logger) *T {
	environment := os.Getenv("ENVIRONMENT")
	if environment == "" {
		currentDir, _ := os.Getwd()
		configPath := filepath.Join(currentDir, ".env.local")
		if _, err := os.Stat(configPath); err == nil {
			err := godotenv.Load(configPath)
			if err != nil {
				log.Warn("Cannot load configuration file", "path", configPath, "error", err.Error())
			} else {
				log.Info("Loaded configuration file", "path", configPath)
			}
		}
	}

	var cfg T
	if err := env.Parse(&cfg); err != nil {
		log.Error("Cannot parse environment variables", "error", err.Error())
		os.Exit(1)
	}

	log.Info("Configuration loaded successfully")
	return &cfg
}
