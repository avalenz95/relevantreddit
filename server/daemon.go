package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

//pass in subreddit includes /r get posts from x time period
func parseSubreddit(subreddit string, postAge float64) []rPosts {
	//build the initial url
	url := fmt.Sprintf("https://api.reddit.com/%s/best", subreddit) // best temporarily for consistent input data

	//send a request
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
		return nil
	} else {
		data := sendRequest(request)

		//check each post to make sure it falls within the time constraints
		//add to list if it does. break if it does not
		//make another requset if we still haven't hit the time limit or after still exists
		var postList []rPosts

		var posts = rPosts{}

		//parse json subreddit struct
		json.Unmarshal(data, &posts)

		return postList
	}
}

//parse comments for a given subreddit post
func parseComments(relPath string) {

	//Url to comments of a post
	url := fmt.Sprintf("https://api.reddit.com/%s", relPath)

	//send a request
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	} else {
		data := sendRequest(request)

		var comments rComments

		json.Unmarshal(data, &comments)
	}

}

//determine if keywords exist in a given comment
func evaluateComments() {}

//Determine if post is within time range? may be redundant
func evaluatePosts() {}
