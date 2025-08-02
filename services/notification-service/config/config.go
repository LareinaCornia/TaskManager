package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	RedisAddr string
	SMTPHost  string
	SMTPPort  int
	SMTPUser  string
	SMTPPass  string
	FromEmail string
	ToEmail   string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Config{
		RedisAddr: os.Getenv("REDIS_ADDR"),
		SMTPHost:  os.Getenv("SMTP_HOST"),
		SMTPPort:  587,
		SMTPUser:  os.Getenv("SMTP_USER"),
		SMTPPass:  os.Getenv("SMTP_PASS"),
		FromEmail: os.Getenv("FROM_EMAIL"),
		ToEmail:   os.Getenv("TO_EMAIL"),
	}
}
