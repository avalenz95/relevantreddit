package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	//Load Environment Variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading environment variables")
	}

	secretKey := os.Getenv("REDDIT_SECRET")

}
