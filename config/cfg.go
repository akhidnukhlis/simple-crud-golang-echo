package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Configuration struct {
	DB_DRIVER   string
	DB_USERNAME string
	DB_PASSWORD string
	DB_HOST     string
	DB_PORT     string
	DB_NAME     string
}

func GetConfig() Configuration {
	conf := Configuration{}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Gagal memuat file .env: %v", err)
	}

	conf.DB_DRIVER = os.Getenv("DB_DRIVER")
	conf.DB_USERNAME = os.Getenv("DB_USERNAME")
	conf.DB_PASSWORD = os.Getenv("DB_PASSWORD")
	conf.DB_PORT = os.Getenv("DB_PORT")
	conf.DB_HOST = os.Getenv("DB_HOST")
	conf.DB_NAME = os.Getenv("DB_NAME")

	return conf
}
