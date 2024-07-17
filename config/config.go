package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	DBDriver  = GetEnv("DB_DRIVER")
	DBName    = GetEnv("DB_NAME")
	DBHost    = GetEnv("DB_HOST")
	DBPort    = GetEnv("DB_PORT")
	DBUser    = GetEnv("DB_USER")
	DBPass    = GetEnv("DB_PASS")
	SSLMode   = GetEnv("SSL_MODE")
	JWTSecret = GetEnv("JWT_SECRET")
)

func GetEnv(key string, value ...string) string {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("err env", err)
		panic("Error Load file .env not found")
	}

	if os.Getenv(key) != "" {
		return os.Getenv(key)
	} else {
		if len(value) > 0 {
			return value[0]
		}
		return ""
	}
}
