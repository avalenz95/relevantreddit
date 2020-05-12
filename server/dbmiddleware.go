package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/ablades/relevantreddit/tries/prefixtree"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*DB MIDDLEWARE*/

var users, tries *mongo.Collection

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

	users = client.Database("test").Collection("Users")

	tries = client.Database("test").Collection("Tries")

	fmt.Println("Connected to Collections: Users  and Tries")

}

//Add a trie to collection
func createTrie(name string) *mongo.SingleResult {
	fmt.Printf("Creating... trie: %s \n", name)
	//subname   string
	//bannerURL string
	//tree      prefixtree.PrefixTree

	//Create a prefix tree
	tree := prefixtree.New(name)

	trie := SubTrie{
		subname: name,
		tree:    tree,
	}

	//Insert into DB
	inserted, err := tries.InsertOne(context.Background(), trie)
	if err != nil {
		fmt.Printf("Failed to insert trie %+v  \n", trie)
		log.Fatal(err)
	}

	//Get and return recent insert
	query := tries.FindOne(context.Background(), inserted.InsertedID)
	return query

}

//look for a trie add it if it doesnt exist
func findTrie(name string) *mongo.SingleResult {
	filter := bson.M{"subname": name}
	query := tries.FindOne(context.Background(), filter)

	if query.Err() == mongo.ErrNoDocuments {
		fmt.Printf("Trie: %s not found.\n", name)
		return createTrie(name)
	}
	fmt.Printf("Trie: %s found. %v \n", name, query)

	return query
}

//Add new word to trie in db
func addToTrie(name string, keyword string, username string, triePtr *mongo.SingleResult) {

	var trie SubTrie
	//Decode db result into trie
	triePtr.Decode(&trie)

	//Insert word into trie
	trie.tree.InsertKeyword(keyword, name)

	//replace trie structure with new one
	tries.ReplaceOne(context.Background(), bson.M{"subname": name}, trie)

}

// get subreddit image from about section add it TODO: HAVE IT SERVE AS AN UPDATE FUNCTION
func addSubBanner(subname string) {
	//Set up request
	url := fmt.Sprintf("https://www.reddit.com/r/%s/about.json", subname)
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	request.Header.Set("User-Agent", fmt.Sprintf("relevant_for_reddit/0.0 (by /u/%s)", creds.Username))

	content := sendRequest(request)

	var as = aboutSubreddit{}
	json.Unmarshal(content, &as)
	fmt.Printf("Banner URL: %s \n", as.Data.BannerImg)

	//Find and decode trie
	var trie SubTrie
	triePtr := findTrie(subname)
	triePtr.Decode(&trie)
	//Add image to trie
	trie.bannerURL = as.Data.BannerImg

	//replace trie structure with new updated one
	tries.ReplaceOne(context.Background(), bson.M{"subname": subname}, trie)

}

//Insert user into DB
func insertUser(user UserProfile) {
	inserted, err := users.InsertOne(context.Background(), user)
	if err != nil {
		fmt.Printf("Failed to insert profile %+v  \n", user)
		log.Fatal(err)
	}

	fmt.Printf("\n Inserted profile %s --> %v \n", user.RedditName, inserted.InsertedID)
}

//findUser in DB
func findUser(userName string) primitive.M {
	filter := bson.M{"redditname": userName}
	result := users.FindOne(context.Background(), filter)
	if result.Err() != nil {
		fmt.Printf("User: %s not found. \n", userName)
		return nil
	}

	fmt.Printf("Found %s \n", userName)
	return filter
}

func getContent(userName string) UserProfile {

	var userProfile UserProfile
	//Find user in store result in user profile
	err := users.FindOne(context.Background(), bson.M{"redditname": userName}).Decode(&userProfile)
	if err != nil {
		log.Fatal(err)
	}

	//loop over all subreddits and content just a test TODO:REMOVE
	for key, element := range userProfile.Subreddits {
		fmt.Println("Key:", key, "=>", "Element:", element)
	}

	return userProfile
}

//Update keywords in database
func updateKeywords(userName string, subreddit string, newWord string) {
	filter := findUser(userName)
	key := fmt.Sprintf("subreddits.r/%s", subreddit)

	//Add all words to array
	update := bson.D{{"$addToSet", bson.D{{key, newWord}}}}

	fmt.Printf("%s ---> %s \n", newWord, subreddit)
	users.UpdateOne(context.Background(), filter, update)
}

//Remove keyword from the database
func removeKeyword(userName string, subreddit string, word string) {

	filter := findUser(userName)
	key := fmt.Sprintf("subreddits.%s", subreddit)

	update := bson.D{{"$pull", bson.D{{key, word}}}}
	users.UpdateOne(context.Background(), filter, update)
}

//Add Subreddit to database
func addSubreddit(userName string, subreddit string) {
	filter := findUser(userName)
	update := bson.D{{"$addToSet", bson.D{{subreddit, []string{}}}}}
	users.UpdateOne(context.Background(), filter, update)
}

func removeUser()      {}
func removeSubreddit() {}
func getAllUsers()     {}

func getUserSubreddits() {}
