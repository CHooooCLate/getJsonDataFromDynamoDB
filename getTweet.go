package main

import (
    "github.com/aws/aws-lambda-go/lambda"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/guregu/dynamo"

    // デバッグ用
    // spew.Dump(value)
    //"github.com/davecgh/go-spew/spew"
    // reflect.TypeOf(value)
    //"reflect"
)

type Tweet struct {
    Tweets []struct {
        Contributors     interface{} `json:"contributors"`
        Coordinates      interface{} `json:"coordinates"`
        CreatedAt        string      `json:"created_at"`
        DisplayTextRange []int       `json:"display_text_range"`
        Entities         struct {
            Urls     []interface{} `json:"urls"`
            Hashtags []interface{} `json:"hashtags"`
            URL      struct {
                Urls interface{} `json:"urls"`
            } `json:"url"`
            UserMentions []interface{} `json:"user_mentions"`
            Media        []struct {
                ID            int64  `json:"id"`
                IDStr         string `json:"id_str"`
                MediaURL      string `json:"media_url"`
                MediaURLHTTPS string `json:"media_url_https"`
                URL           string `json:"url"`
                DisplayURL    string `json:"display_url"`
                ExpandedURL   string `json:"expanded_url"`
                Sizes         struct {
                    Medium struct {
                        W      int    `json:"w"`
                        H      int    `json:"h"`
                        Resize string `json:"resize"`
                    } `json:"medium"`
                    Thumb struct {
                        W      int    `json:"w"`
                        H      int    `json:"h"`
                        Resize string `json:"resize"`
                    } `json:"thumb"`
                    Small struct {
                        W      int    `json:"w"`
                        H      int    `json:"h"`
                        Resize string `json:"resize"`
                    } `json:"small"`
                    Large struct {
                        W      int    `json:"w"`
                        H      int    `json:"h"`
                        Resize string `json:"resize"`
                    } `json:"large"`
                } `json:"sizes"`
                SourceStatusID    int    `json:"source_status_id"`
                SourceStatusIDStr string `json:"source_status_id_str"`
                Type              string `json:"type"`
                Indices           []int  `json:"indices"`
                VideoInfo         struct {
                    AspectRatio    interface{} `json:"aspect_ratio"`
                    DurationMillis int         `json:"duration_millis"`
                    Variants       interface{} `json:"variants"`
                } `json:"video_info"`
                ExtAltText string `json:"ext_alt_text"`
            } `json:"media"`
        } `json:"entities"`
        ExtendedEntities struct {
            Urls     interface{} `json:"urls"`
            Hashtags interface{} `json:"hashtags"`
            URL      struct {
                Urls interface{} `json:"urls"`
            } `json:"url"`
            UserMentions interface{} `json:"user_mentions"`
            Media        []struct {
                ID            int64  `json:"id"`
                IDStr         string `json:"id_str"`
                MediaURL      string `json:"media_url"`
                MediaURLHTTPS string `json:"media_url_https"`
                URL           string `json:"url"`
                DisplayURL    string `json:"display_url"`
                ExpandedURL   string `json:"expanded_url"`
                Sizes         struct {
                    Medium struct {
                        W      int    `json:"w"`
                        H      int    `json:"h"`
                        Resize string `json:"resize"`
                    } `json:"medium"`
                    Thumb struct {
                        W      int    `json:"w"`
                        H      int    `json:"h"`
                        Resize string `json:"resize"`
                    } `json:"thumb"`
                    Small struct {
                        W      int    `json:"w"`
                        H      int    `json:"h"`
                        Resize string `json:"resize"`
                    } `json:"small"`
                    Large struct {
                        W      int    `json:"w"`
                        H      int    `json:"h"`
                        Resize string `json:"resize"`
                    } `json:"large"`
                } `json:"sizes"`
                SourceStatusID    int    `json:"source_status_id"`
                SourceStatusIDStr string `json:"source_status_id_str"`
                Type              string `json:"type"`
                Indices           []int  `json:"indices"`
                VideoInfo         struct {
                    AspectRatio    []int `json:"aspect_ratio"`
                    DurationMillis int   `json:"duration_millis"`
                    Variants       []struct {
                        Bitrate     int    `json:"bitrate"`
                        ContentType string `json:"content_type"`
                        URL         string `json:"url"`
                    } `json:"variants"`
                } `json:"video_info"`
                ExtAltText string `json:"ext_alt_text"`
            } `json:"media"`
        } `json:"extended_entities"`
        ExtendedTweet struct {
            FullText         string      `json:"full_text"`
            DisplayTextRange interface{} `json:"display_text_range"`
            Entities         struct {
                Urls     interface{} `json:"urls"`
                Hashtags interface{} `json:"hashtags"`
                URL      struct {
                    Urls interface{} `json:"urls"`
                } `json:"url"`
                UserMentions interface{} `json:"user_mentions"`
                Media        interface{} `json:"media"`
            } `json:"entities"`
            ExtendedEntities struct {
                Urls     interface{} `json:"urls"`
                Hashtags interface{} `json:"hashtags"`
                URL      struct {
                    Urls interface{} `json:"urls"`
                } `json:"url"`
                UserMentions interface{} `json:"user_mentions"`
                Media        interface{} `json:"media"`
            } `json:"extended_entities"`
        } `json:"extended_tweet"`
        FavoriteCount        int    `json:"favorite_count"`
        Favorited            bool   `json:"favorited"`
        FilterLevel          string `json:"filter_level"`
        FullText             string `json:"full_text"`
        HasExtendedProfile   bool   `json:"has_extended_profile"`
        ID                   int64  `json:"id"`
        IDStr                string `json:"id_str"`
        InReplyToScreenName  string `json:"in_reply_to_screen_name"`
        InReplyToStatusID    int    `json:"in_reply_to_status_id"`
        InReplyToStatusIDStr string `json:"in_reply_to_status_id_str"`
        InReplyToUserID      int    `json:"in_reply_to_user_id"`
        InReplyToUserIDStr   string `json:"in_reply_to_user_id_str"`
        IsTranslationEnabled bool   `json:"is_translation_enabled"`
        Lang                 string `json:"lang"`
        Place                struct {
            Attributes  interface{} `json:"attributes"`
            BoundingBox struct {
                Coordinates interface{} `json:"coordinates"`
                Type        string      `json:"type"`
            } `json:"bounding_box"`
            ContainedWithin interface{} `json:"contained_within"`
            Country         string      `json:"country"`
            CountryCode     string      `json:"country_code"`
            FullName        string      `json:"full_name"`
            Geometry        struct {
                Coordinates interface{} `json:"coordinates"`
                Type        string      `json:"type"`
            } `json:"geometry"`
            ID        string      `json:"id"`
            Name      string      `json:"name"`
            PlaceType string      `json:"place_type"`
            Polylines interface{} `json:"polylines"`
            URL       string      `json:"url"`
        } `json:"place"`
        QuotedStatusID              int         `json:"quoted_status_id"`
        QuotedStatusIDStr           string      `json:"quoted_status_id_str"`
        QuotedStatus                interface{} `json:"quoted_status"`
        PossiblySensitive           bool        `json:"possibly_sensitive"`
        PossiblySensitiveAppealable bool        `json:"possibly_sensitive_appealable"`
        RetweetCount                int         `json:"retweet_count"`
        Retweeted                   bool        `json:"retweeted"`
        RetweetedStatus             interface{} `json:"retweeted_status"`
        Source                      string      `json:"source"`
        Scopes                      interface{} `json:"scopes"`
        Text                        string      `json:"text"`
        User                        struct {
            ContributorsEnabled bool   `json:"contributors_enabled"`
            CreatedAt           string `json:"created_at"`
            DefaultProfile      bool   `json:"default_profile"`
            DefaultProfileImage bool   `json:"default_profile_image"`
            Description         string `json:"description"`
            Email               string `json:"email"`
            Entities            struct {
                Urls     interface{} `json:"urls"`
                Hashtags interface{} `json:"hashtags"`
                URL      struct {
                    Urls interface{} `json:"urls"`
                } `json:"url"`
                UserMentions interface{} `json:"user_mentions"`
                Media        interface{} `json:"media"`
            } `json:"entities"`
            FavouritesCount                int         `json:"favourites_count"`
            FollowRequestSent              bool        `json:"follow_request_sent"`
            FollowersCount                 int         `json:"followers_count"`
            Following                      bool        `json:"following"`
            FriendsCount                   int         `json:"friends_count"`
            GeoEnabled                     bool        `json:"geo_enabled"`
            HasExtendedProfile             bool        `json:"has_extended_profile"`
            ID                             int64       `json:"id"`
            IDStr                          string      `json:"id_str"`
            IsTranslator                   bool        `json:"is_translator"`
            IsTranslationEnabled           bool        `json:"is_translation_enabled"`
            Lang                           string      `json:"lang"`
            ListedCount                    int         `json:"listed_count"`
            Location                       string      `json:"location"`
            Name                           string      `json:"name"`
            Notifications                  bool        `json:"notifications"`
            ProfileBackgroundColor         string      `json:"profile_background_color"`
            ProfileBackgroundImageURL      string      `json:"profile_background_image_url"`
            ProfileBackgroundImageURLHTTPS string      `json:"profile_background_image_url_https"`
            ProfileBackgroundTile          bool        `json:"profile_background_tile"`
            ProfileBannerURL               string      `json:"profile_banner_url"`
            ProfileImageURL                string      `json:"profile_image_url"`
            ProfileImageURLHTTPS           string      `json:"profile_image_url_https"`
            ProfileLinkColor               string      `json:"profile_link_color"`
            ProfileSidebarBorderColor      string      `json:"profile_sidebar_border_color"`
            ProfileSidebarFillColor        string      `json:"profile_sidebar_fill_color"`
            ProfileTextColor               string      `json:"profile_text_color"`
            ProfileUseBackgroundImage      bool        `json:"profile_use_background_image"`
            Protected                      bool        `json:"protected"`
            ScreenName                     string      `json:"screen_name"`
            ShowAllInlineMedia             bool        `json:"show_all_inline_media"`
            Status                         interface{} `json:"status"`
            StatusesCount                  int         `json:"statuses_count"`
            TimeZone                       string      `json:"time_zone"`
            URL                            string      `json:"url"`
            UtcOffset                      int         `json:"utc_offset"`
            Verified                       bool        `json:"verified"`
            WithheldInCountries            interface{} `json:"withheld_in_countries"`
            WithheldScope                  string      `json:"withheld_scope"`
        } `json:"user"`
        WithheldCopyright   bool        `json:"withheld_copyright"`
        WithheldInCountries interface{} `json:"withheld_in_countries"`
        WithheldScope       string      `json:"withheld_scope"`
    } `json:"tweets"`
}

type Data struct {
  Id int `dynamo:"id"`
  SearchWord string `dynamo:"search_word"`
  TweetData Tweet `dynamo:"tweet_data"`
  CreatedAt string `dynamo:"created_at"`
}

type Request struct {
    Animal string `json:"animal"`
}

type Response struct {
    Id string `json:"id"`
    UserId string `json:"user_id"`
}

func Handler(request Request) ([]Response, error) {
    var word string

    switch request.Animal {
    case "dog":
        word = "dog"
    case "cat":
        word = "cat"
    case "fish":
        word = "fish"
    default:
        word = "dog"
    }

    // DynamoDBへ接続
    db := dynamo.New(session.New(), &aws.Config{
        Region: aws.String("us-east-2"), // "ap-northeast-1"等
    })

    table := db.Table("Tweet")

    var datas []Data
    err := table.Get("word", word).All(&datas)
    if err != nil {
        panic(err.Error())
    }

    var responses []Response

    for _, data := range datas {
        for _, tweet := range data.TweetData.Tweets {
            response := Response {
                Id: tweet.IDStr,
                UserId: tweet.User.ScreenName,
            }

            responses = append(responses, response)
        }
    }

    return responses, nil
}

func main() {
    lambda.Start(Handler)
}
