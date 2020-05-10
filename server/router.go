package main

import (
	"github.com/gorilla/mux"
)

var route = mux.NewRouter()

//Set routes dealing with api and db
func router() *mux.Router {

	/* Reddit Auth Routes */
	route.HandleFunc("/", handleMain).Methods("GET", "OPTIONS")
	route.HandleFunc("/r/login", handleRedditLogin).Methods("GET", "OPTIONS")
	route.HandleFunc("/r/callback", handleRedditCallback).Methods("GET", "OPTIONS")
	/* DB/Endpoint routes */
	route.HandleFunc("/user/{username}", getUserContent).Methods("GET", "OPTIONS").Name("user") //Users homepage

	route.HandleFunc("/img/r/{subreddit}", getSubredditImg).Methods("GET", "OPTIONS") //route for pulling subreddit image

	//User wants to add a word
	//User wants to delete a word

	return route

}
