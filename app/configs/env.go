package configs

import (
	"fmt"
	"log"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type EnvironmentVariable struct {
	APP_PORT                          string `env:"APP_PORT"`
	APP_USERNAME                      string `env:"APP_USERNAME"`
	APP_PASSWORD                      string `env:"APP_PASSWORD"`
	DB_HOST                           string `env:"DB_HOST"`
	DB_NAME                           string `env:"DB_NAME"`
	DB_USER                           string `env:"DB_USER"`
	DB_PASSWORD                       string `env:"DB_PASSWORD"`
	DB_SSL_MODE                       string `env:"DB_SSL_MODE" default:"disable"`
	DB_PORT                           int    `env:"DB_PORT"`
}

func LoadEnv() *EnvironmentVariable {
	envFile := ".env"
	if err := godotenv.Load(envFile); err != nil {
		log.Fatal(fmt.Sprintf("Error loading :%s file ", envFile))
	}

	fev := EnvironmentVariable{}

	if err := env.Parse(&fev); err != nil {
		log.Fatal(fmt.Sprintln("Error parse env ", err))
	}

	return &fev
}