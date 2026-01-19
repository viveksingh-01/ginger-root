package config

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	Port          string
	MongoURI      string
	Database      string
	AllowedOrigin string
}

func Load() (*Config, error) {
	godotenv.Load()

	viper.SetDefault("PORT", "8080")
	viper.AutomaticEnv()

	cfg := &Config{
		Port:          viper.GetString("PORT"),
		MongoURI:      viper.GetString("MONGO_URI"),
		Database:      viper.GetString("MONGO_DB"),
		AllowedOrigin: viper.GetString("ALLOWED_ORIGIN"),
	}

	// Basic validation
	if cfg.MongoURI == "" {
		return nil, fmt.Errorf("MONGO_URI is required")
	}

	if cfg.Database == "" {
		return nil, fmt.Errorf("MONGO_DB is required")
	}

	return cfg, nil
}
