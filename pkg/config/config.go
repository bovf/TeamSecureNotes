package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

// Config stores all configuration for the application.
type Config struct {
	MongoDBUser     string
	MongoDBPassword string
	MongoDBHost     string
	MongoDBPort     string 
	MongoDBName     string
	MongoDBURI      string
	Port            string
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file, using environment variables")
	}


	config := Config{
		MongoDBUser:     getEnv("MONGODB_USER", ""),
		MongoDBPassword: getEnv("MONGODB_PASSWORD", ""),
		MongoDBHost:     getEnv("MONGODB_HOST", "localhost"),
		MongoDBPort:     getEnv("MONGODB_PORT", "27017"),
		MongoDBName:     getEnv("MONGODB_DBNAME", ""),
		Port:            getEnv("PORT", "8080"),
	}

	return config
}

// getEnv reads an environment variable or returns a default value.
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

