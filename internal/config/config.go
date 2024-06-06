package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Database Database
}

type Database struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int64  `yaml:"port"`
	DbName   string `yaml:"dbName"`
	SslMode  string `yaml:"sslMode"`
}

func MustLoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("Config path is not set")
	}

	// check if file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file %s does not exist", configPath)
	}

	var config Config

	if err := cleanenv.ReadConfig(configPath, &config); err != nil {
		log.Fatalf("Cannot read config: %s", err)
	}

	return &config, nil
}
