package config

import (
	"backend/internal/utils"
)

var (
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string
)

func LoadConfig() {	
	DBUser = utils.MustGetEnv("DB_USER")
	DBPassword = utils.MustGetEnv("DB_PASSWORD")
	DBHost = utils.MustGetEnv("DB_HOST")
	DBPort = utils.MustGetEnv("DB_PORT")
	DBName = utils.MustGetEnv("DB_NAME")
}