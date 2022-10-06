package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Load read the env files
func Load(files ...string) {
	isInCloud := len(os.Getenv("POD_IP")) > 0
	err := godotenv.Load(files...)
	if err != nil && !isInCloud {
		log.Println(err)
	}
}
