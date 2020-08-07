package main

import (
	"github.com/ablades/prefix"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//SubTrie subreddit trie
type SubTrie struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Subname   string             `json:"subname,omitempty" bson:"subname,omitempty"`
	BannerURL string             `json:"bannerurl,omitempty" bson:"bannerurl,omitempty"`
	Tree      prefix.Tree        `json:"tree,omitempty" bson:"tree,omitempty"`
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
type redditPosts struct {
	Kind string `json:"kind"`
	Data struct {
		Modhash  string `json:"modhash"`
		Dist     int    `json:"dist"`
		Children []struct {
			Kind string `json:"kind"`
			Data struct {
				ApprovedAtUtc       interface{} `json:"approved_at_utc"`
				Subreddit           string      `json:"subreddit"`
				Selftext            string      `json:"selftext"`
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

//The comments for a subreddit TODO: DEAL WITH REPLIES
type redditComments []struct {
	Kind string `json:"kind"`
	Data struct {
		Modhash  string      `json:"modhash"`
		Dist     interface{} `json:"dist"`
		Children []struct {
			Kind string `json:"kind"`
			Data struct {
				TotalAwardsReceived   int           `json:"total_awards_received"`
				Ups                   int           `json:"ups"`
				Awarders              []interface{} `json:"awarders"`
				ModReasonBy           interface{}   `json:"mod_reason_by"`
				BannedBy              interface{}   `json:"banned_by"`
				AuthorFlairType       string        `json:"author_flair_type"`
				RemovalReason         interface{}   `json:"removal_reason"`
				LinkID                string        `json:"link_id"`
				AuthorFlairTemplateID string        `json:"author_flair_template_id"`
				Likes                 interface{}   `json:"likes"`
				Replies               struct {
					Kind string `json:"kind"`
					Data struct {
						Modhash  string      `json:"modhash"`
						Dist     interface{} `json:"dist"`
						Children []struct {
							Kind string `json:"kind"`
							Data struct {
								TotalAwardsReceived   int           `json:"total_awards_received"`
								ApprovedAtUtc         interface{}   `json:"approved_at_utc"`
								Ups                   int           `json:"ups"`
								Awarders              []interface{} `json:"awarders"`
								ModReasonBy           interface{}   `json:"mod_reason_by"`
								BannedBy              interface{}   `json:"banned_by"`
								AuthorFlairType       string        `json:"author_flair_type"`
								RemovalReason         interface{}   `json:"removal_reason"`
								LinkID                string        `json:"link_id"`
								AuthorFlairTemplateID interface{}   `json:"author_flair_template_id"`
								Likes                 interface{}   `json:"likes"`
								Replies               string        `json:"replies"`
								UserReports           []interface{} `json:"user_reports"`
								Saved                 bool          `json:"saved"`
								ID                    string        `json:"id"`
								BannedAtUtc           interface{}   `json:"banned_at_utc"`
								ModReasonTitle        interface{}   `json:"mod_reason_title"`
								Gilded                int           `json:"gilded"`
								Archived              bool          `json:"archived"`
								NoFollow              bool          `json:"no_follow"`
								Author                string        `json:"author"`
								CanModPost            bool          `json:"can_mod_post"`
								SendReplies           bool          `json:"send_replies"`
								ParentID              string        `json:"parent_id"`
								Score                 int           `json:"score"`
								AuthorFullname        string        `json:"author_fullname"`
								ReportReasons         interface{}   `json:"report_reasons"`
								ApprovedBy            interface{}   `json:"approved_by"`
								AllAwardings          []interface{} `json:"all_awardings"`
								SubredditID           string        `json:"subreddit_id"`
								Collapsed             bool          `json:"collapsed"`
								Body                  string        `json:"body"`
								Edited                bool          `json:"edited"`
								AuthorFlairCSSClass   interface{}   `json:"author_flair_css_class"`
								IsSubmitter           bool          `json:"is_submitter"`
								Downs                 int           `json:"downs"`
								AuthorFlairRichtext   []interface{} `json:"author_flair_richtext"`
								AuthorPatreonFlair    bool          `json:"author_patreon_flair"`
								BodyHTML              string        `json:"body_html"`
								Gildings              struct {
								} `json:"gildings"`
								CollapsedReason              interface{}   `json:"collapsed_reason"`
								AssociatedAward              interface{}   `json:"associated_award"`
								Stickied                     bool          `json:"stickied"`
								AuthorPremium                bool          `json:"author_premium"`
								SubredditType                string        `json:"subreddit_type"`
								CanGild                      bool          `json:"can_gild"`
								TopAwardedType               interface{}   `json:"top_awarded_type"`
								AuthorFlairTextColor         interface{}   `json:"author_flair_text_color"`
								ScoreHidden                  bool          `json:"score_hidden"`
								Permalink                    string        `json:"permalink"`
								NumReports                   interface{}   `json:"num_reports"`
								Locked                       bool          `json:"locked"`
								Name                         string        `json:"name"`
								Created                      int           `json:"created"`
								Subreddit                    string        `json:"subreddit"`
								AuthorFlairText              interface{}   `json:"author_flair_text"`
								TreatmentTags                []interface{} `json:"treatment_tags"`
								CreatedUtc                   int           `json:"created_utc"`
								SubredditNamePrefixed        string        `json:"subreddit_name_prefixed"`
								Controversiality             int           `json:"controversiality"`
								Depth                        int           `json:"depth"`
								AuthorFlairBackgroundColor   interface{}   `json:"author_flair_background_color"`
								CollapsedBecauseCrowdControl interface{}   `json:"collapsed_because_crowd_control"`
								ModReports                   []interface{} `json:"mod_reports"`
								ModNote                      interface{}   `json:"mod_note"`
								Distinguished                interface{}   `json:"distinguished"`
							} `json:"data"`
						} `json:"children"`
						After  string      `json:"after"`
						Before interface{} `json:"before"`
					} `json:"data"`
				} `json:"replies"`
				UserReports         []interface{} `json:"user_reports"`
				Saved               bool          `json:"saved"`
				ID                  string        `json:"id"`
				BannedAtUtc         interface{}   `json:"banned_at_utc"`
				ModReasonTitle      interface{}   `json:"mod_reason_title"`
				Gilded              int           `json:"gilded"`
				Archived            bool          `json:"archived"`
				NoFollow            bool          `json:"no_follow"`
				Author              string        `json:"author"`
				CanModPost          bool          `json:"can_mod_post"`
				SendReplies         bool          `json:"send_replies"`
				ParentID            string        `json:"parent_id"`
				Score               int           `json:"score"`
				AuthorFullname      string        `json:"author_fullname"`
				ReportReasons       interface{}   `json:"report_reasons"`
				SubredditID         string        `json:"subreddit_id"`
				Body                string        `json:"body"`
				Edited              bool          `json:"edited"`
				Downs               int           `json:"downs"`
				AuthorFlairCSSClass string        `json:"author_flair_css_class"`
				IsSubmitter         bool          `json:"is_submitter"`
				Collapsed           bool          `json:"collapsed"`
				AuthorFlairRichtext []interface{} `json:"author_flair_richtext"`
				AuthorPatreonFlair  bool          `json:"author_patreon_flair"`
				BodyHTML            string        `json:"body_html"`
				Gildings            struct {
				} `json:"gildings"`
				Stickied              bool        `json:"stickied"`
				AuthorPremium         bool        `json:"author_premium"`
				SubredditType         string      `json:"subreddit_type"`
				Permalink             string      `json:"permalink"`
				NumReports            interface{} `json:"num_reports"`
				Locked                bool        `json:"locked"`
				Name                  string      `json:"name"`
				Created               int         `json:"created"`
				Subreddit             string      `json:"subreddit"`
				CreatedUtc            int         `json:"created_utc"`
				SubredditNamePrefixed string      `json:"subreddit_name_prefixed"`
				Controversiality      int         `json:"controversiality"`
				Depth                 int         `json:"depth"`
			} `json:"data"`
		} `json:"children"`
		After  string      `json:"after"`
		Before interface{} `json:"before"`
	} `json:"data"`
}

/* https://api.reddit.com/r/politics/comments/fmzu7n/trump_has_completely_mishandled_the_coronavirus?limit=2/*/
