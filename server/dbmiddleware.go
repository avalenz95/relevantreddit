package main

import (

	"context"
	"log"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*DB MIDDLEWARE*/

var collection *mongo.Collection

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

	collection = client.Database("test").Collection("Users")

	fmt.Println("Collection: Users instance created!")

}

//Insert user into DB
func insertUser(user UserProfile) {
	inserted, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		fmt.Printf("Failed to insert profile %+v  \n", user)
		log.Fatal(err)
	}

	fmt.Printf("\n Inserted profile %s --> %v \n", user.RedditName, inserted.InsertedID)
}

//findUser in DB
func findUser(userName string) primitive.M {
	filter := bson.M{"redditname": userName}
	result := collection.FindOne(context.Background(), filter)
	if result.Err() != nil {
		fmt.Printf("User: %s not found. \n", userName)
		return nil
	}

	fmt.Printf("Found %s \n", userName)
	return filter
}


//Update keywords in database
func updateKeywords(userName string, subreddit string, newWords []string) {

	filter := findUser(userName)
	key := fmt.Sprintf("subreddits.%s", subreddit)

	//Add all words to array
	for _, word := range newWords {
		update := bson.D{{"$addToSet", bson.D{{key, word}}}}

		fmt.Printf("%s ---> %s \n", word, subreddit)
		collection.UpdateOne(context.Background(), filter, update)
	}

}

//Remove keyword from the database
func removeKeyword(userName string, subreddit string, word string) {

	filter := findUser(userName)
	key := fmt.Sprintf("subreddits.%s", subreddit)

	update := bson.D{{"$pull", bson.D{{key, word}}}}
	collection.UpdateOne(context.Background(), filter, update)
}

//Add Subreddit to database
func addSubreddit(userName string, subreddit string) {
	filter := findUser(userName)
	update := bson.D{{"$addToSet", bson.D{{subreddit, []string{}}}}}
	collection.UpdateOne(context.Background(), filter, update)
}


func removeUser()      {}
func removeSubreddit() {}
func getAllUsers() {}

func getUserSubreddits() {}