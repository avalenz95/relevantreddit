package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

//pass in subreddit includes /r get posts from x time period
func fetchSubredditPosts(trie *SubTrie, queue chan notifcation, wg *sync.WaitGroup) {
	//defer wg.Done() // wait for goroutine to finish before decrementing
	defer fmt.Printf("\033[32m Done with %s \033[0m", trie.Subname)
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
	// use permalink for each post to pull comments TODO:Seems like first comment of every post is an empty one
	for _, post := range posts.Data.Children {
		fmt.Println(post.Data.Title)
		fmt.Printf("Fetching Comments for: %s \n", post.Data.Permalink)
		//Get comments from each post
		fetchComments(post.Data.Permalink, trie, queue)
		fmt.Printf("----DONE FETCHING FOR %s \n", post.Data.Permalink)
	}
}

//parse comments for a given subreddit post
func fetchComments(relPath string, trie *SubTrie, queue chan notifcation) {

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
	//Process Comment
	for _, c := range comments {
		for _, comment := range c.Data.Children {
			//Check words against Trie
			processComment(comment.Data.Body, trie, queue)
		}
	}
}

//Strip comment of punctuation and other characters
func processComment(comment string, trie *SubTrie, queue chan notifcation) {
	fmt.Printf("  ---------  Processing comment: %s \n", comment)
	//Remove links and urls
	re, _ := regexp.Compile(`((\[(.+)\]\(.+\))|(https?:\/\/)?([\da-z\.-]+)\.([a-z\.]{2,6})([\/\w \.-]*))`)

	r := strings.NewReplacer(",", "", ".", "", ";", "")
	parsedComment := strings.Fields(r.Replace(comment))

	for _, word := range parsedComment {
		users := trie.Tree.Contains(word)
		if len(users) > 0 {
			for _, user := range users {
				fmt.Printf("\033[32m Added Notification to channel for User: %s  with word: %s \033[0m \n ", user, word)
				//Add to channel

			}
		}
	}

}

//Determine if post is within time range? may be redundant
func parsePosts(posts []redditPosts) {

}

//
type notifcation struct {
	name string
	msg  string
}

func daemon() {
	//Make a notification map
	//notificationMap := make(map[string][]string)
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

	//Maintains count of go routines
	var wg sync.WaitGroup

	notificationQueue := make(chan notifcation)
	fmt.Println("Tries")
	//Gets posts for each trie concurrently
	wgCounter := 0
	for _, trie := range allTries {
		fmt.Printf(" %s \n", trie.Subname)
		wgCounter++
		//wg.Add(1)
		fetchSubredditPosts(trie, notificationQueue, &wg)
		fmt.Println(wgCounter)
	}
	fmt.Println("BEFORE WAIT")
	//wg.Wait() // Wait till goroutines finish
	fmt.Println("AFTER WAIT")
	close(notificationQueue) // close channel - no more values will be added

	//Create Map based off values in channel
	notificationMap := make(map[string][]string)

	fmt.Printf("\n --Map of Notifications-- \n  %+v \n  END DAEMON :))))) \n \n", notificationMap)

}
