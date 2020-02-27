package main

import (
	"log"

	"github.com/joho/godotenv"
)

func main() {

	//Load Environment Variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading environment variables")
	}
}
