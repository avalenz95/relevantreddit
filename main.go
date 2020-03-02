package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

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

func request() string {
	//Load Environment Variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading environment variables")
	}

	secretKey := os.Getenv("SCRIPT_REDDIT_SECRET")
	client := os.Getenv("SCRIPT_CLIENT")
	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")

	body := strings.NewReader(fmt.Sprintf("grant_type=password&username=%s&password=%s", username, password))

	//Create new http post request
	req, err := http.NewRequest("POST", "https://www.reddit.com/api/v1/access_token", body)
	if err != nil {
		log.Fatal(err)
	}

	//Set authorization request
	req.SetBasicAuth(client, secretKey)

	//Curl command header entries
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
	}
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {

	}

	return string(content)
}

func main() {
	fmt.Println(request())

}
