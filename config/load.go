package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	*Server
	*DB
}

type Server struct {
	Address string
	Port    string
}

type DB struct {
	User     string
	Password string
	Schema   string
	Host     string
	Port     string
}

func LoadEnv() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %v", err)
	}

	newConfig := &Config{
		Server: &Server{
			Address: os.Getenv("APP_ADDRESS"),
			Port:    os.Getenv("APP_PORT"),
		},
		DB: &DB{
			User:     os.Getenv("MYSQL_USER"),
			Password: os.Getenv("MYSQL_PASSWORD"),
			Schema:   os.Getenv("MYSQL_SCHEMA"),
			Host:     os.Getenv("MYSQL_HOST"),
			Port:     os.Getenv("MYSQL_PORT"),
		},
	}

	return newConfig, nil
}
