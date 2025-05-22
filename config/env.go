package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUser                 string
	DBPassword             string
	DBName                 string
	JWTExpirationInSeconds int64
	JWTSecret              string
}

var Envs = initConfig()

func initConfig() Config {
	godotenv.Load()
	return Config{
		DBUser:                 getEnv("DB_USER", "postgres"),
		DBPassword:             getEnv("DB_PASSWORD", "admin"),
		DBName:                 getEnv("DB_NAME", "BackendAPI"),
		JWTExpirationInSeconds: getEnvAsInt("JWT_EXP", 3600*24*7),
		JWTSecret:              getEnv("JWT_SECRET", "supersecret123"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getEnvAsInt(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fallback
		}

		return i
	}

	return fallback
}
