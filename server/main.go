package main

import (
	"fmt"
	"log"
	"net/http"
)

//displays subreddit
func subcribedReddits(rc subreddits, u *UserProfile) {
	//Loop through all of a requests subreddits
	for _, item := range rc.Data.Children {
		fmt.Printf("Added: %s to user profile. \n", item.Data.DisplayNamePrefixed)

		u.Subreddits[item.Data.DisplayNamePrefixed] = nil
	}

	fmt.Println(rc.Data.After)
}

func main() {
	loadEnvironment()

	//Start router
	r := router()

	fmt.Println("Starting server on the port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
