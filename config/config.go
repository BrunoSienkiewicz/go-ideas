package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	DbUser     string
	DbPassword string
	DbName     string
	DbHost     string
	DbPort     string
}

type ConfigOption func(*Config)

func loadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func (c *Config) GetDbConnectionString() string {
	return "postgres://" + c.DbUser + ":" + c.DbPassword + "@" + c.DbHost + ":" + c.DbPort + "/" + c.DbName + "?sslmode=disable"
}

func NewConfig(opts ...ConfigOption) *Config {
	loadEnv()

	c := &Config{
		DbUser:     os.Getenv("DB_USER"),
		DbPassword: os.Getenv("DB_PASSWORD"),
		DbName:     os.Getenv("DB_NAME"),
		DbHost:     os.Getenv("DB_HOST"),
		DbPort:     os.Getenv("DB_PORT"),
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

func WithDbUser(user string, password string) ConfigOption {
	return func(c *Config) {
		c.DbUser = user
		c.DbPassword = password
	}
}
