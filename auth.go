//Handle user authentication
package main

import (
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
	requestToken(code)

}

/*
func authRequest(cred credentials) {
	req, err := http.NewRequest("GET", "https://www.reddit.com/api/v1/authorize.compact", nil)
	if err != nil {
		log.Print(err)
	}

	req.Header.Set("User-Agent", fmt.Sprintf("relevant_for_reddit/0.0 (by /u/%s)", cred.Username))
	//Build request query string
	q := req.URL.Query()
	q.Add("client_id", cred.Client)
	q.Add("response_type", "code")
	q.Add("state", "foobar")                                               //verify user is user CSRF
	q.Add("redirect_uri", "https://www.github.com/ablades/relevantreddit") //temp redirect url
	q.Add("duration", "temporary")                                         //temp for now may switch to perm later
	q.Add("scope", "mysubreddits identity history")

	req.URL.RawQuery = q.Encode()

	fmt.Println(req.URL.String())
}
*/
//Request a token from reddit server
func requestToken(code string) {
	//Load Environment Variables
	body := strings.NewReader(fmt.Sprintf("grant_type=authorization_code&code=%s&redirect_uri=%s", code, creds.Redirect))
	//Create new http post request
	request, err := http.NewRequest("POST", "https://www.reddit.com/api/v1/access_token", body)
	if err != nil {
		log.Fatal(err)
	}

	//Set authorization header request
	request.SetBasicAuth(creds.Client, creds.Secret)
	request.Header.Set("User-Agent", fmt.Sprintf("relevant_for_reddit/0.0 (by /u/%s)", creds.Username))

	sendRequest(request)

	//fmt.Println(string(content))
	//Create empty token request variable
	//var tokenRequest = token{}

	//parse json into token struct
	//json.Unmarshal(content, &tokenRequest)

	//return tokenRequest
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
