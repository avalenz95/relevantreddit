//Handle user authentication
package auth

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
	SecretKey string
	Client    string
	Username  string
	Password  string
}

func authRequest() {
	req, err := http.NewRequest("GET", "https://www.reddit.com/api/v1/authorize", nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	//Build request query string
	q := req.URL.Query()
	q.Add("client_id", os.Getenv("APP_CLIENT"))
	q.Add("response_type", "code")
	q.Add("state", "foo & bar")
	q.Add("redirect_uri", "https://www.github.com/ablades/relevantreddit") //temp redirect url
	q.Add("duration", "temporary")                                         //temp for now may switch to perm later
	q.Add("scope", "mysubreddits identity history")

	req.URL.RawQuery = q.Encode()

	fmt.Println(req.URL.String())
}

/*
//Request a token from reddit server
func requestToken(creds credentials) token {
	//Load Environment Variables

	body := strings.NewReader(fmt.Sprintf("grant_type=password&username=%s&password=%s", creds.Username, creds.Password))

	//Create new http post request
	req, err := http.NewRequest("POST", "https://www.reddit.com/api/v1/access_token", body)
	if err != nil {
		log.Fatal(err)
	}

	//Set authorization request
	req.SetBasicAuth(creds.Client, creds.SecretKey)

	//header entries
	req.Header.Set("User-Agent", fmt.Sprintf("relevant_for_reddit/0.0 (by /u/%s)", creds.Username))

	content := sendRequest(req)
	//Create empty token request variable
	var tokenRequest = token{}

	//parse json into token struct
	json.Unmarshal(content, &tokenRequest)

	return tokenRequest
}

//Load enviornment variables
func loadEnvironment() credentials {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading environment variables")
	}

	var cred = credentials{}

	cred.SecretKey = os.Getenv("SCRIPT_REDDIT_SECRET")
	cred.Client = os.Getenv("SCRIPT_CLIENT")
	cred.Username = os.Getenv("USERNAME")
	cred.Password = os.Getenv("PASSWORD")

	return cred

}

func auth() {

	credentials := loadEnvironment()

	token := requestToken(credentials)

	//Use token once
	rc := useToken(token, credentials)

	//send multiple requests till all are pulled
	for rc.Data.After != "" {
		rc = useToken(token, credentials, rc.Data.After)
	}
}
*/
