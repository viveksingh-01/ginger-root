package config

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	Port     string
	MongoURI string
}

func Load() (*Config, error) {
	godotenv.Load()

	viper.SetDefault("PORT", "8080")
	viper.AutomaticEnv()

	cfg := &Config{
		Port:     viper.GetString("PORT"),
		MongoURI: viper.GetString("MONGODB_URI"),
	}

	return cfg, nil
}
