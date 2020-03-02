package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

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

//TokenRequest stores authentication request
type tokenRequest struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Expires     int    `json:"expires_in"`
	Scope       string `json:"scope"`
}

func readResponse() {

}

//Sends an http request returns response in bytes
func sendRequest(request *http.Request) []byte {

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	return content
}

func requestToken() tokenRequest {
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

	//header entries
	req.Header.Set("User-Agent", fmt.Sprintf("relevant_for_reddit/0.0 (by /u/%s)", username))

	content := sendRequest(req)
	//Create empty token request variable
	var tokenRequest = tokenRequest{}

	//parse json into token struct
	json.Unmarshal(content, &tokenRequest)

	return tokenRequest
}

func useToken(tr tokenRequest) {

	req, err := http.NewRequest("GET", "https://oauth.reddit.com/api/v1/me", nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("%s %s", tr.TokenType, tr.AccessToken))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

}

func main() {

	fmt.Print(requestToken())

}
