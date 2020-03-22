//Handle user authentication
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

//Store user credentials for easier access
type credentials struct {
	Secret   string
	Client   string
	Username string
	Password string
	Redirect string
	MongoURI string
}

//TokenRequest stores authentication request
type token struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Expires     int    `json:"expires_in"`
	Scope       string `json:"scope"`
}

var creds credentials

//Load enviornment variables
func loadEnvironment() {
	err := godotenv.Load("../.env")

	if err != nil {
		log.Fatal("Error loading environment variables")
	}

	creds.Secret = os.Getenv("APP_SECRET")
	creds.Client = os.Getenv("APP_CLIENT")
	creds.Username = os.Getenv("USERNAME")
	creds.Password = os.Getenv("PASSWORD")
	creds.Redirect = "http://localhost:8080/r/callback"
}

//Request a token from reddit server
func requestToken(code string) token {

	body := strings.NewReader(fmt.Sprintf("grant_type=authorization_code&code=%s&redirect_uri=%s", code, creds.Redirect))
	//Create new http post request
	request, err := http.NewRequest("POST", "https://www.reddit.com/api/v1/access_token", body)
	if err != nil {
		log.Fatal(err)
	}

	//Set authorization header request
	request.SetBasicAuth(creds.Client, creds.Secret)
	request.Header.Set("User-Agent", fmt.Sprintf("relevant_for_reddit/0.0 (by /u/%s)", creds.Username))

	response := sendRequest(request)

	//fmt.Println(string(content))
	//Create empty token request variable
	var tokenRequest = token{}

	//parse json into token struct
	json.Unmarshal(response, &tokenRequest)

	return tokenRequest
}

func getUserInfo(t token, endpoint string) userInfo {
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		log.Fatal(err)
	}

	//Set required headers
	req.Header.Set("User-Agent", fmt.Sprintf("relevant_for_reddit/0.0 (by /u/%s)", creds.Username))
	req.Header.Set("Authorization", fmt.Sprintf("%s %s", t.TokenType, t.AccessToken))

	//send request
	content := sendRequest(req)

	var userContent = userInfo{}

	json.Unmarshal(content, &userContent)

	return userContent
}

func useToken(t token, url string) subreddits {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	//api/v1/me

	//Set required headers
	req.Header.Set("User-Agent", fmt.Sprintf("relevant_for_reddit/0.0 (by /u/%s)", creds.Username))
	req.Header.Set("Authorization", fmt.Sprintf("%s %s", t.TokenType, t.AccessToken))

	//send request
	content := sendRequest(req)

	//fmt.Println(string(content))
	var rc = subreddits{}

	json.Unmarshal(content, &rc)

	return rc
}

//Sends an http request returns response in bytes
func sendRequest(request *http.Request) []byte {
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	//set up middleware? to handle request ect?

	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println(string(content))
	return content
}
