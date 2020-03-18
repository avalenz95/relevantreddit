//Handle user authentication
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
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
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading environment variables")
	}

	creds.Secret = os.Getenv("APP_SECRET")
	creds.Client = os.Getenv("APP_CLIENT")
	creds.Username = os.Getenv("USERNAME")
	creds.Password = os.Getenv("PASSWORD")
	creds.Redirect = "http://localhost:3000/RedditCallback"
}

//Handles request/response intermidary actions
func redditMiddleware() {
	http.HandleFunc("/", handleMain)
	http.HandleFunc("/RedditLogin", handleRedditLogin)
	http.HandleFunc("/RedditCallback", handleRedditCallback)
	fmt.Println(http.ListenAndServe(":3000", nil))

}

func handleMain(w http.ResponseWriter, r *http.Request) {
	const htmlIndex = `
	<html>
		<body>
			<a href="/RedditLogin">Log in with Google</a>
		</body>
	</html>
	`
	fmt.Fprintf(w, htmlIndex)
}

//Redirect user to authorization page
func handleRedditLogin(w http.ResponseWriter, r *http.Request) {

	url, err := url.Parse("reddit.com/api/v1/authorize.compact")
	if err != nil {
		log.Fatal(err)
	}
	url.Scheme = "https"
	q := url.Query()
	q.Add("client_id", creds.Client)
	q.Add("response_type", "code")
	q.Add("state", "foobar")              //verify user is user CSRF
	q.Add("redirect_uri", creds.Redirect) //temp redirect url
	q.Add("duration", "temporary")        //temp for now may switch to perm later
	q.Add("scope", "mysubreddits identity history")

	url.RawQuery = q.Encode()
	fmt.Println(url)

	http.Redirect(w, r, url.String(), http.StatusTemporaryRedirect)
}

//Handle unser response from reddit
func handleRedditCallback(w http.ResponseWriter, r *http.Request) {
	fmt.Println("here")
	//Get first parameter with query name
	code := r.FormValue("code")
	state := r.FormValue("state") //TODO : Verify states are same
	fmt.Println(state)
	fmt.Println(code)
	//
	token := requestToken(code)

	useToken(token)

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

//variadic argument - closest thing to an optional argument accepts a variable number of arguments usd as a list
func useToken(t token, after ...string) {
	//variables init
	var req *http.Request
	var err error

	//Pull a users subreddits
	if len(after) == 1 {
		req, err = http.NewRequest("GET", fmt.Sprintf("https://oauth.reddit.com/subreddits/mine/subscriber.json?limit=100&after=%s", after[0]), nil)
	} else {
		req, err = http.NewRequest("GET", "https://oauth.reddit.com/subreddits/mine/subscriber.json?limit=100", nil)
	}
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

	fmt.Println(string(content))
	/*
		var rc = subreddits{}

		json.Unmarshal(content, &rc)
		subcribedReddits(rc)

		return rc */
}

//Sends an http request returns response in bytes
func sendRequest(request *http.Request) []byte {
	fmt.Println("Before request")
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

	fmt.Println(string(content))
	return content
}
