package config

import (
	"os"

	"github.com/gofiber/fiber/v2/log"

	"github.com/joho/godotenv"
)

// AppEnvironment is a type for application environment
type AppEnvironment string

// AppEnvironment constants
const (
	AppEnvDevelopment AppEnvironment = "development"
)

// appEnv indicates the environment of the application
var appEnv AppEnvironment

// InitConfig initializes the application configuration
func InitConfig() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Warn("Error loading .env file, using default environment variables")
	}

	// Set application environment
	appEnvStr := os.Getenv("APP_ENV")
	if appEnvStr == "" {
		appEnv = AppEnvDevelopment
	} else {
		appEnv = AppEnvironment(appEnvStr)
	}

}

// AppEnv returns the application environment
func AppEnv() AppEnvironment {
	return appEnv
}
