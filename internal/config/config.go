package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Port      string
	AppEnv    string
	DBUrl     string
	JwtSecret string
	JwtExpire int
}

var C Config

func Load() {
	_ = godotenv.Load() // .env が無くてもエラーにしない

	C = Config{
		Port:      getEnv("PORT", "8080"),
		AppEnv:    getEnv("APP_ENV", "development"),
		DBUrl:     getEnv("DB_URL", ""),
		JwtSecret: getEnv("JWT_SECRET", ""),
		JwtExpire: getEnvInt("JWT_EXPIRE", 3600),
	}
}

func getEnv(key, def string) string {
	if v, ok := os.LookupEnv(key); ok {
		return v
	}
	return def
}

func getEnvInt(key string, def int) int {
	if v, ok := os.LookupEnv(key); ok {
		i, err := strconv.Atoi(v)
		if err == nil {
			return i
		}
	}
	return def
}
