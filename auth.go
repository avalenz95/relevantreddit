//Handle user authentication
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

//Store user credentials for easier access
type credentials struct {
	SecretKey string
	Client    string
	Username  string
	Password  string
}

//Load enviornment variables
func loadEnvironment() credentials {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading environment variables")
	}

	var cred = credentials{}

	cred.SecretKey = os.Getenv("SCRIPT_REDDIT_SECRET")
	cred.Client = os.Getenv("APP_CLIENT")
	cred.Username = os.Getenv("USERNAME")
	cred.Password = os.Getenv("PASSWORD")

	return cred
}

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

func handleRedditLogin(w http.ResponseWriter, r *http.Request) {

	cred := loadEnvironment()

	req, err := http.NewRequest("GET", "https://www.reddit.com/api/v1/authorize.compact", nil)
	if err != nil {
		log.Print(err)
	}

	req.Header.Set("User-Agent", fmt.Sprintf("relevant_for_reddit/0.0 (by /u/%s)", cred.Username))
	//Build request query string
	q := req.URL.Query()
	q.Add("client_id", cred.Client)
	q.Add("response_type", "code")
	q.Add("state", "foobar")                                       //verify user is user CSRF
	q.Add("redirect_uri", "http://localhost:3000/RedditCallback.") //temp redirect url
	q.Add("duration", "temporary")                                 //temp for now may switch to perm later
	q.Add("scope", "mysubreddits identity history")

	req.URL.RawQuery = q.Encode()

	url := req.URL.String()

	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func handleRedditCallback(w http.ResponseWriter, r *http.Request) {
}

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
