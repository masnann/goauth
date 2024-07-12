package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	DBDriver       = GetEnv("DB_DRIVER")
	DBName         = GetEnv("DB_NAME")
	DBHost         = GetEnv("DB_HOST")
	DBPort         = GetEnv("DB_PORT")
	DBUser         = GetEnv("DB_USER")
	DBPass         = GetEnv("DB_PASS")
	SSLMode        = GetEnv("SSL_MODE")
	UploadAPIURL   = GetEnv("UPLOAD_API_URL")
	UploadFolder   = GetEnv("UPLOAD_FOLDER")
	UploadUsername = GetEnv("UPLOAD_USERNAME")
	UploadPassword = GetEnv("UPLOAD_PASSWORD")
	UpdateCashUrl  = GetEnv("UPDATE_CASH_URL")
	UpdateCashAuth = GetEnv("UPDATE_CASH_AUTH")
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
