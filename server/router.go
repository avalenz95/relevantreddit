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
	route.HandleFunc("/u/{username}", handleUser).Methods("GET", "OPTIONS").Name("user")

	return route

}
