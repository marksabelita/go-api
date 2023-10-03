package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnv(c string) string {
	c = os.Getenv(c)
	return c
}
 
func EnvMongoURI() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("MONGO_URI")
}