package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type redditAuth struct {
	ClientID      string
	ClientSecret  string
	response_type "code"
}

// func future_auth() {
// 	ctx := context.Background()

// 	oauth2.SetAuthURLParam("key", "value")

// 	conf := &oauth2.Config{
// 		ClientID:     os.Getenv("CLIENT"),
// 		ClientSecret: os.Getenv("REDDIT_SECRET"),
// 		Scopes:       nil,
// 		Endpoint:     oauth2.Endpoint{},
// 	}

// //Authentication configuration variable
// var OauthConfig *oauth2.Config

// //Load Environment Variables
// err := godotenv.Load()
// if err != nil {
// 	log.Fatal("Error loading environment variables")
// }

// secretKey := os.Getenv("REDDIT_SECRET")

// }

func main() {

	//Load Environment Variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading environment variables")
	}

	secretKey := os.Getenv("SCRIPT_REDDIT_SECRET")
	client := os.Getenv("SCRIPT_CLIENT")
	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")

}
