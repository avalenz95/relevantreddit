package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
)

func getUserContent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	//payload :=
	//json.NewEncoder(w).Encode(payload)
	params := mux.Vars(r)
	redditMap := getContent(params["username"])
	json.NewEncoder(w).Encode(redditMap)
	//fmt.Printf("Passed in Username is: %s \n", params["username"])
}

//temp till connect react frontend
func handleMain(w http.ResponseWriter, r *http.Request) {
	const htmlIndex = `
	<html>
		<body>
			<a href="/r/login">Authenticate with Reddit</a>
		</body>
	</html>
	`
	fmt.Fprintf(w, htmlIndex)
}

//Redirect user to authorization page
func handleRedditLogin(w http.ResponseWriter, r *http.Request) {

	url, err := url.Parse("reddit.com/api/v1/authorize")
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
	fmt.Printf("Redirecting to: %s \n", url)

	http.Redirect(w, r, url.String(), http.StatusTemporaryRedirect)
}

//Handle unser response from reddit
func handleRedditCallback(w http.ResponseWriter, r *http.Request) {
	//Get first parameter with query name
	code := r.FormValue("code")
	state := r.FormValue("state") //TODO : Verify states are same
	fmt.Println(state)
	fmt.Println(code)

	//Request token from reddit server
	token := requestToken(code)

	//Pull user info from endpoint
	userInfo := getUserInfo(token, "https://oauth.reddit.com/api/v1/me")

	//Check if user is in DB
	if findUser(userInfo.Name) == nil {
		fmt.Println("User Not Found! Creating User")
		//Create User Profile
		var appUser UserProfile
		//Update fields
		appUser.RedditName = userInfo.Name
		appUser.Subreddits = make(map[string][]string)

		url := "https://oauth.reddit.com/subreddits/mine/subscriber.json?limit=100"
		rc := useToken(token, url)
		//Get all of a users subreddits
		subscribedReddits(rc, &appUser)
		for rc.Data.After != "" {
			url = fmt.Sprintf("https://oauth.reddit.com/subreddits/mine/subscriber.json?limit=100&after=%s", rc.Data.After)
			rc = useToken(token, url)

			subscribedReddits(rc, &appUser)
		}

		insertUser(appUser)

	} else {
		fmt.Println("User found... Redirecting")
		//updateKeywords(userInfo.Name, "r/apexlegends", []string{"test1", "test2"})
	}

	//redirect user to their homepage
	routeURL, _ := route.Get("user").URL("username", userInfo.Name)
	http.Redirect(w, r, routeURL.String(), http.StatusPermanentRedirect)

}
