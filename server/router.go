package main

import (
	"github.com/gorilla/mux"
)

var routes = mux.NewRouter()

//Set routes dealing with api and db
func router() *mux.Router {

	/* Reddit Auth Routes */
	routes.HandleFunc("/r/login", handleRedditLogin).Methods("GET", "OPTIONS")
	routes.HandleFunc("/r/callback", handleRedditCallback).Methods("GET", "OPTIONS")
	/* DB/Endpoint routes */
	routes.HandleFunc("/user/{username}", getUserContent).Methods("GET", "OPTIONS").Name("user") //Users homepage

	// Get sub banner image
	routes.HandleFunc("/banner/r/{subreddit}", getSubBanner).Methods("GET", "OPTIONS") //route for pulling subreddit image

	//User wants to add a word
	routes.HandleFunc("/addkeyword", addKeyword).Methods("POST", "OPTIONS")
	//User wants to delete a word

	return routes

}
