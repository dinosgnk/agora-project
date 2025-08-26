package config

type AppConfig struct {
	Environment string `env:"ENVIRONMENT"`
	Port        string `env:"PORT"`
}
