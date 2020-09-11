package main

import (
	"fmt"
	"log"
	"net/http"
)

//displays subreddit
func subscribedReddits(rc subreddits, u *UserProfile) {
	//Loop through all of a requests subreddits
	for _, item := range rc.Data.Children {
		fmt.Printf("Added: %s to user profile. \n", item.Data.DisplayNamePrefixed)

		u.Subreddits[item.Data.DisplayNamePrefixed] = make([]string, 0)
	}

	fmt.Println(rc.Data.After)
}

func main() {
	loadEnvironment()

	//Start router
	r := router()
	//daemon()

	fmt.Println("Starting server on the port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
