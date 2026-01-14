package config

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	Port     string
	MongoURI string
	Database string
}

func Load() (*Config, error) {
	godotenv.Load()

	viper.SetDefault("PORT", "8080")
	viper.AutomaticEnv()

	cfg := &Config{
		Port:     viper.GetString("PORT"),
		MongoURI: viper.GetString("MONGO_URI"),
		Database: viper.GetString("MONGO_DB"),
	}

	return cfg, nil
}
