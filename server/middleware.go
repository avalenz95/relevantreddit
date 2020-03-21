package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*DB MIDDLEWARE*/

var collection *mongo.Collection

//Create mongodb cllection

func init() {
	//Client options
	clientOptions := options.Client().ApplyURI("mongodb+srv://ablades:atlaspass@cluster0-pwh5o.mongodb.net/test?retryWrites=true&w=majority")

	//Connect to DB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	collection = client.Database("test").Collection("Users")

	fmt.Println("Collection: Users instance created!")

}

//Insert user into DB
func insertUser(user UserProfile) {
	inserted, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Inserted profile %+v --> %v \n", user, inserted.InsertedID)
}

//findUser in DB
func findUser(userName string) bool {
	filter := bson.M{"RedditName": userName}
	err := collection.FindOne(context.Background(), filter)
	if err != nil {
		fmt.Printf("User: %s not found. \n", userName)
		return false
	}
	return true
}

func getAllUsers() {}
func getUserSubreddits() {

}
func updateKeywords()  {}
func removeUser()      {}
func removeKeyword()   {}
func removeSubreddit() {}
func addSubreddit()    {}

func handleUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	w.Write([]byte(params["username"]))
	fmt.Printf("Passed in Username is: %s ", params["username"])
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
	fmt.Println("here")
	//Get first parameter with query name
	code := r.FormValue("code")
	state := r.FormValue("state") //TODO : Verify states are same
	fmt.Println(state)
	fmt.Println(code)
	//
	token := requestToken(code)

	var appUser UserProfile
	appUser.Subreddits = make(map[string][]string)

	url := "https://oauth.reddit.com/subreddits/mine/subscriber.json?limit=100"
	rc := useToken(token, url)

	subcribedReddits(rc, &appUser)

	//send multiple requests till all are pulled
	for rc.Data.After != "" {
		url = fmt.Sprintf("https://oauth.reddit.com/subreddits/mine/subscriber.json?limit=100&after=%s", rc.Data.After)
		rc = useToken(token, url)

		subcribedReddits(rc, &appUser)
	}

	endpoint := "https://oauth.reddit.com/api/v1/me"

	appUser.RedditName = getUserInfo(token, endpoint).Name

	fmt.Print(appUser)

	//redirect user to their homepage
	routeURL, _ := route.Get("user").URL("username", appUser.RedditName)
	http.Redirect(w, r, routeURL.String(), http.StatusPermanentRedirect)

}
