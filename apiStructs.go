package main

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

//Content from each subreddit
type subredditContent []struct {
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
