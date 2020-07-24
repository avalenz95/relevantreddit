package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

//pass in subreddit includes /r get posts from x time period
func fetchSubredditPosts(subreddit string) redditPosts {
	//build the initial url
	url := fmt.Sprintf("https://api.reddit.com/%s/new", subreddit) // best temporarily for consistent input data

	//send a request
	request, err := http.NewRequest("GET", url, nil)
	request.Header.Set("User-Agent", fmt.Sprintf("relevant_for_reddit/0.0 (by /u/%s)", creds.Username))

	if err != nil {
		log.Fatal(err)
	}
	data := sendRequest(request)

	fmt.Printf("visiting: %s \n", url)
	//check each post to make sure it falls within the time constraints
	//add to list if it does. break if it does not
	//make another requset if we still haven't hit the time limit or after still exist
	var posts redditPosts

	//parse json subreddit struct
	json.Unmarshal(data, &posts)
	// use permalink for each post to pull comments
	for _, post := range posts.Data.Children {
		fmt.Println(post.Data.Title)
		fmt.Printf("Fetching Comments for: %s \n", post.Data.Permalink)
	}

	return posts
}

//parse comments for a given subreddit post
func fetchComments(relPath string) {

	//Url to comments of a post
	url := fmt.Sprintf("https://api.reddit.com%s", relPath)

	//send a request
	request, err := http.NewRequest("GET", url, nil)
	request.Header.Set("User-Agent", fmt.Sprintf("relevant_for_reddit/0.0 (by /u/%s)", creds.Username))

	if err != nil {
		log.Fatal(err)
		fmt.Printf("ERROR: %s", err)
	}

	data := sendRequest(request)

	var comments redditComments

	json.Unmarshal(data, &comments)

	for _, c := range comments {
		for _, comment := range c.Data.Children {
			fmt.Printf(" -- %s \n", comment.Data.Body)
		}
	}

}

//determine if keywords exist in a given comment
func evaluateComments() {}

//Determine if post is within time range? may be redundant
func parsePosts(posts []redditPosts) {

}

func daemon() {
	// Get Tries Collection
	//Unmarshall
	//Iterate over all tries
	//Call fetchPosts for each trie
	//Call fetchComments for each post

}
