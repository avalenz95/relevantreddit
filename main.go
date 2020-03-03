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

type subredditContent struct {
	Subreddit struct {
		BannerImg         string      `json:"banner_img"`
		CommunityIcon     string      `json:"community_icon"`
		IconColor         string      `json:"icon_color"`
		DisplayName       string      `json:"display_name"`
		HeaderImg         interface{} `json:"header_img"`
		Title             string      `json:"title"`
		IconImg           string      `json:"icon_img"`
		Description       string      `json:"description"`
		Subscribers       int         `json:"subscribers"`
		Name              string      `json:"name"`
		URL               string      `json:"url"`
		UserIsModerator   bool        `json:"user_is_moderator"`
		PublicDescription string      `json:"public_description"`
		SubredditType     string      `json:"subreddit_type"`
		UserIsSubscriber  bool        `json:"user_is_subscriber"`
	} `json:"subreddit"`

	OauthClientID string `json:"oauth_client_id"`
	Name          string `json:"name"`
	CommentKarma  int    `json:"comment_karma"`
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

//TokenRequest stores authentication request
type token struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Expires     int    `json:"expires_in"`
	Scope       string `json:"scope"`
}

//Store user credentials for easier access
type credentials struct {
	SecretKey string
	Client    string
	Username  string
	Password  string
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

//
func useToken(t token, creds credentials) []byte {
	//get api endpoint
	req, err := http.NewRequest("GET", "https://oauth.reddit.com/api/v1/me", nil)
	if err != nil {
		log.Fatal(err)
	}
	//Set required headers
	req.Header.Set("User-Agent", fmt.Sprintf("relevant_for_reddit/0.0 (by /u/%s)", creds.Username))
	req.Header.Set("Authorization", fmt.Sprintf("%s %s", t.TokenType, t.AccessToken))

	//send request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	//convert response
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return content
}

func main() {

	credentials := loadEnvironment()

	token := requestToken(credentials)

	fmt.Println(token)

	fmt.Println(string(useToken(token, credentials)))

}
