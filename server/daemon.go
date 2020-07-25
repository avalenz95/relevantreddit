package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

//pass in subreddit includes /r get posts from x time period
func fetchSubredditPosts(trie *SubTrie) {
	//build the initial url
	url := fmt.Sprintf("https://api.reddit.com/%s/new", trie.Subname) // best temporarily for consistent input data

	//send a request
	request, err := http.NewRequest("GET", url, nil)
	request.Header.Set("User-Agent", fmt.Sprintf("relevant_for_reddit/0.0 (by /u/%s)", creds.Username))

	if err != nil {
		log.Fatal(err)
	}
	data := sendRequest(request)

	fmt.Printf("VISITING: %s \n", url)
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
		//Get comments from each post
		fetchComments(post.Data.Permalink, trie)
	}
}

//parse comments for a given subreddit post
func fetchComments(relPath string, trie *SubTrie) {

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
	//TODO: THINK ABUOT HOW TO HANDLE DEPTH
	for _, c := range comments {
		for _, comment := range c.Data.Children {
			//Check words against Trie
			for _, word := comment.Data.Body {
				users := trie.Tree.Contains(word)
				//Send to channel
				if len(users) > 0:
			}
		}
	}

}


//Determine if post is within time range? may be redundant
func parsePosts(posts []redditPosts) {

}

//

func daemon() {
	//Make a notification map
	//Anytime a keyword returns add that post to users notification map
	// Get Tries Collection
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()
	var allTries []*SubTrie

	cursor, err := tries.Find(context.TODO(), bson.D{})
	defer cursor.Close(ctx)
	if err != nil {
		log.Fatal(err)
	} else {
		for cursor.Next(ctx) {
			var trie SubTrie
			cursor.Decode(&trie)
			allTries = append(allTries, &trie)
		}
	}

	fmt.Printf("Tries: %+v", allTries)
	//channel := make(chan struct {})
	for _, trie := range allTries {
		fmt.Printf("%s \n  ------ \n", trie.Subname)
		go fetchSubredditPosts(trie)
	}

	//Unmarshall
	//Iterate over all tries
	//Call fetchPosts for each trie should probably be done concurrently

	//Call fetchComments for each post
	//Once channel is empty (all subs have been procceded)
	//Start notifying users? maybe this should be concurrent instead? another channel?

}
