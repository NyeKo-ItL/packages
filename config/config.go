package config

import (
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/NyeKo-ItL/packages/utils"
	"github.com/joho/godotenv"
)

// Config is a struct for the database and webservice config
type Config struct {
	PublicHost           string
	Port                 string
	DBUser               string
	DBPassword           string
	DBHost               string
	DBName               string
	JWTSecret            string
	JWTExpirationInHours int
	ApplicationName      string
	LastVersion          string
}

var baseDir, mod = utils.GetBaseDirectory()

// CFG initialize the configuration file and return the config
func CFG() Config {

	envPath := filepath.Join(baseDir, "../../.env")

	if err := godotenv.Load(envPath); err != nil {
		log.Printf("Loading defaults envs: %s", err)
	}

	// Load configuration values from environment variables
	return Config{
		PublicHost:           getEnv("PUBLIC_HOST", "http://localhost"),
		Port:                 getEnv("PORT", "3333"),
		DBUser:               getEnv("DB_USER", "user"),
		DBPassword:           getEnv("DB_PASSWORD", "root"),
		DBHost:               getEnv("DB_HOST", "postgres"),
		DBName:               getEnv("DB_NAME", ""),
		JWTSecret:            getEnv("JWT_SECRET", "reifg5349Gfjg3FIO4G945"),
		JWTExpirationInHours: getEnvAsInt("JWT_LIFETIME", 24),
		ApplicationName:      getEnv("APPLICATION", "myApp"),
		LastVersion:          getEnv("LAST_VERSION", "1.0"),
	}
}

// Gets the env by key or fallbacks
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func getEnvAsInt(key string, fallback int) int {
	if value, ok := os.LookupEnv(key); ok {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return fallback
}
