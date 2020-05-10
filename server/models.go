package main

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Tries for all subreddits
type SubredditTries struct {
	subreddits map[string][]tries.prefix
}

//UserProfile that goes into DB
type UserProfile struct {
	ID         primitive.ObjectID  `json:"_id,omitempty" bson:"_id,omitempty"`
	RedditName string              `json:"username,omitempty"`
	Subreddits map[string][]string `json:"subreddits,omitempty"` //key is r/[subreddit] value list of keywords
}

//User info from api endpoint
type userInfo struct {
	PrefNightmode bool   `json:"pref_nightmode"`
	OauthClientID string `json:"oauth_client_id"`
	Name          string `json:"name"`
}

//subreddits holds content from api
type subreddits struct {
	Data struct {
		Children []struct {
			Data struct {
				DisplayName         string `json:"display_name"`
				Subscribers         int    `json:"subscribers"`
				Name                string `json:"name"`
				ID                  string `json:"id"`
				DisplayNamePrefixed string `json:"display_name_prefixed"`
				Description         string `json:"description"`
				URL                 string `json:"url"`
			} `json:"data"`
		} `json:"children"`
		After string `json:"after"`
	} `json:"data"`
}

//structure of a subreddit post
type rPosts []struct {
	Kind string `json:"kind"`
	Data struct {
		Modhash  string `json:"modhash"`
		Dist     int    `json:"dist"`
		Children []struct {
			Kind string `json:"kind"`
			Data struct {
				ApprovedAtUtc       interface{} `json:"approved_at_utc"`
				Subreddit           string      `json:"subreddit"`
				Title               string      `json:"title"`
				Name                string      `json:"name"`
				Ups                 int         `json:"ups"`
				TotalAwardsReceived int         `json:"total_awards_received"`
				Edited              bool        `json:"edited"`
				ContentCategories   interface{} `json:"content_categories"`
				Created             float64     `json:"created"`
				ViewCount           interface{} `json:"view_count"`
				Archived            bool        `json:"archived"`
				Score               int         `json:"score"`
				Over18              bool        `json:"over_18"`
				Spoiler             bool        `json:"spoiler"`
				Locked              bool        `json:"locked"`
				SubredditID         string      `json:"subreddit_id"`
				Author              string      `json:"author"`
				NumComments         int         `json:"num_comments"`
				Permalink           string      `json:"permalink"`
				URL                 string      `json:"url"`
				CreatedUtc          float64     `json:"created_utc"`
			} `json:"data"`
		} `json:"children"`
		After string `json:"after"`
	} `json:"data"`
}

//the about page for a subreddit
type aboutSubreddit struct {
	Data struct {
		DisplayName           string  `json:"display_name"`
		HeaderImg             string  `json:"header_img"`
		Title                 string  `json:"title"`
		PrimaryColor          string  `json:"primary_color"`
		ActiveUserCount       int     `json:"active_user_count"`
		IconImg               string  `json:"icon_img"`
		DisplayNamePrefixed   string  `json:"display_name_prefixed"`
		Subscribers           int     `json:"subscribers"`
		Name                  string  `json:"name"`
		PublicDescription     string  `json:"public_description"`
		CommunityIcon         string  `json:"community_icon"`
		BannerBackgroundImage string  `json:"banner_background_image"`
		DescriptionHTML       string  `json:"description_html"`
		HeaderTitle           string  `json:"header_title"`
		HeaderSize            []int   `json:"header_size"`
		KeyColor              string  `json:"key_color"`
		Created               float64 `json:"created"`
		PublicDescriptionHTML string  `json:"public_description_html"`
		BannerImg             string  `json:"banner_img"`
		BannerBackgroundColor string  `json:"banner_background_color"`
		ID                    string  `json:"id"`
		Description           string  `json:"description"`
		URL                   string  `json:"url"`
		BannerSize            []int   `json:"banner_size"`
		MobileBannerImage     string  `json:"mobile_banner_image"`
	} `json:"data"`
}

//The comments for a subreddit
type rComments []struct {
	Kind string `json:"kind"`
	Data struct {
		Modhash  string      `json:"modhash"`
		Dist     interface{} `json:"dist"`
		Children []struct {
			Kind string `json:"kind"`
			Data struct {
				TotalAwardsReceived int         `json:"total_awards_received"`
				ApprovedAtUtc       interface{} `json:"approved_at_utc"`
				Ups                 int         `json:"ups"`
				ModReasonBy         interface{} `json:"mod_reason_by"`
				BannedBy            interface{} `json:"banned_by"`
				AuthorFlairType     string      `json:"author_flair_type"`
				LinkID              string      `json:"link_id"`
				Replies             string      `json:"replies"`
				ID                  string      `json:"id"`
				Gilded              int         `json:"gilded"`
				Archived            bool        `json:"archived"`
				Author              string      `json:"author"`
				ParentID            string      `json:"parent_id"`
				Score               int         `json:"score"`
				SubredditID         string      `json:"subreddit_id"`
				Body                string      `json:"body"`
				Downs               int         `json:"downs"`
				BodyHTML            string      `json:"body_html"`
				Gildings            struct {
				} `json:"gildings"`
				Stickied              bool   `json:"stickied"`
				Subreddit             string `json:"subreddit"`
				Permalink             string `json:"permalink"`
				Locked                bool   `json:"locked"`
				Name                  string `json:"name"`
				Created               int    `json:"created"`
				CreatedUtc            int    `json:"created_utc"`
				SubredditNamePrefixed string `json:"subreddit_name_prefixed"`
				Controversiality      int    `json:"controversiality"`
				Depth                 int    `json:"depth"`
				Distinguished         string `json:"distinguished"`
			} `json:"data,omitempty"`
		} `json:"children"`
		After  interface{} `json:"after"`
		Before interface{} `json:"before"`
	} `json:"data"`
}

/* https://api.reddit.com/r/politics/comments/fmzu7n/trump_has_completely_mishandled_the_coronavirus?limit=2/*/
