package main

import (
	"context"
	"fmt"
	"log"

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

	fmt.Println("Connected to Collections: Users, Tries")

}

//Create/Update Banners
//

//Add a trie to collection
func createTrie(subname string) {
	fmt.Printf("Creating... trie: %s \n", subname)
	//subname   string
	//bannerURL string
	//tree      prefixtree.PrefixTree

	//Create a prefix tree
	tree := prefixtree.New(subname)
	bannerURL := fetchSubredditBanner(subname)

	trie := SubTrie{
		Subname:   subname,
		Tree:      tree,
		BannerURL: bannerURL,
	}

	//Insert into DB
	inserted, err := tries.InsertOne(context.Background(), trie)
	if err != nil {
		fmt.Printf("Failed to insert trie %+v  \n", trie)
		log.Fatal(err)
	}
	fmt.Printf("%+v", inserted)

}

//look for a trie add it if it doesnt exist REFORMAT LATER
func foundTrie(subname string) bool {
	filter := bson.M{"subname": subname}
	query := tries.FindOne(context.Background(), filter)

	if query.Err() == mongo.ErrNoDocuments {
		fmt.Printf("Trie: %s not found.\n", subname)
		createTrie(subname)
		return false
	}
	fmt.Printf("Trie: %s found. %v \n", subname, query)

	return true
}

//Add new word to trie in db
func addKeywordToTrie(subname string, keyword string, username string) {

	filter := bson.M{"subname": subname}
	query := tries.FindOne(context.Background(), filter)

	var trie SubTrie
	//Decode db result into trie
	query.Decode(&trie)
	fmt.Printf("\n TRIE: >>> %+v \n", trie)
	//Insert word into trie
	trie.Tree.InsertKeyword(keyword, username)

	fmt.Printf("\n INSERTED TRIE: >>> %+v \n", trie)
	//replace trie structure with new one
	tries.ReplaceOne(context.Background(), bson.M{"subname": subname}, trie)

}

func getTrieBanner(subname string) string {
	//get trie from db
	filter := bson.M{"subname": subname}
	query := tries.FindOne(context.Background(), filter)

	var trie SubTrie
	//Decode db result into trie
	query.Decode(&trie)

	return trie.BannerURL
}

// update a trie banner
func updateTrieBanner(subname string) {
	//get trie from db
	filter := bson.M{"subname": subname}
	query := tries.FindOne(context.Background(), filter)

	var trie SubTrie
	//Decode db result into trie
	query.Decode(&trie)
	fmt.Printf("\n Updating TRIE: >>> %v \n", trie.Subname)

	trie.BannerURL = fetchSubredditBanner(subname)

	//replace trie structure with new one
	tries.ReplaceOne(context.Background(), bson.M{"subname": subname}, trie)

}

//Routine that will update all banner images in db
func updateAllBanners() {}

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
func updateUserKeywords(userName string, subreddit string, newWord string) {
	filter := findUser(userName)
	key := fmt.Sprintf("subreddits.%s", subreddit)

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
