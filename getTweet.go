package main

import (
    "encoding/json"
    "fmt"

    "github.com/aws/aws-lambda-go/lambda"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/guregu/dynamo"

    // デバッグ用
    // spew.Dump(value)
    "github.com/davecgh/go-spew/spew"
    // reflect.TypeOf(value)
    //"reflect"
)

type Tweet struct {
    Tweets []struct {
        Contributors     interface{} `json:"contributors,string"`
        Coordinates      interface{} `json:"coordinates,string"`
        CreatedAt        string      `json:"created_at"`
        DisplayTextRange []int       `json:"display_text_range,string"`
        Entities         struct {
            Urls     []interface{} `json:"urls,string"`
            Hashtags []interface{} `json:"hashtags,string"`
            URL      struct {
                Urls interface{} `json:"urls,string"`
            } `json:"url,string"`
            UserMentions []interface{} `json:"user_mentions,string"`
            Media        []struct {
                ID            int64  `json:"id,string"`
                IDStr         string `json:"id_str"`
                MediaURL      string `json:"media_url"`
                MediaURLHTTPS string `json:"media_url_https"`
                URL           string `json:"url"`
                DisplayURL    string `json:"display_url"`
                ExpandedURL   string `json:"expanded_url"`
                Sizes         struct {
                    Medium struct {
                        W      int    `json:"w,string"`
                        H      int    `json:"h,string"`
                        Resize string `json:"resize"`
                    } `json:"medium,string"`
                    Thumb struct {
                        W      int    `json:"w,string"`
                        H      int    `json:"h,string"`
                        Resize string `json:"resize"`
                    } `json:"thumb,string"`
                    Small struct {
                        W      int    `json:"w,string"`
                        H      int    `json:"h,string"`
                        Resize string `json:"resize"`
                    } `json:"small,string"`
                    Large struct {
                        W      int    `json:"w,string"`
                        H      int    `json:"h,string"`
                        Resize string `json:"resize"`
                    } `json:"large,string"`
                } `json:"sizes,string"`
                SourceStatusID    int    `json:"source_status_id,string"`
                SourceStatusIDStr string `json:"source_status_id_str"`
                Type              string `json:"type"`
                Indices           []int  `json:"indices,string"`
                VideoInfo         struct {
                    AspectRatio    interface{} `json:"aspect_ratio,string"`
                    DurationMillis int         `json:"duration_millis,string"`
                    Variants       interface{} `json:"variants,string"`
                } `json:"video_info,string"`
                ExtAltText string `json:"ext_alt_text"`
            } `json:"media,string"`
        } `json:"entities,string"`
        ExtendedEntities struct {
            Urls     interface{} `json:"urls,string"`
            Hashtags interface{} `json:"hashtags,string"`
            URL      struct {
                Urls interface{} `json:"urls,string"`
            } `json:"url,string"`
            UserMentions interface{} `json:"user_mentions,string"`
            Media        []struct {
                ID            int64  `json:"id,string"`
                IDStr         string `json:"id_str"`
                MediaURL      string `json:"media_url"`
                MediaURLHTTPS string `json:"media_url_https"`
                URL           string `json:"url"`
                DisplayURL    string `json:"display_url"`
                ExpandedURL   string `json:"expanded_url"`
                Sizes         struct {
                    Medium struct {
                        W      int    `json:"w,string"`
                        H      int    `json:"h,string"`
                        Resize string `json:"resize"`
                    } `json:"medium,string"`
                    Thumb struct {
                        W      int    `json:"w,string"`
                        H      int    `json:"h,string"`
                        Resize string `json:"resize"`
                    } `json:"thumb,string"`
                    Small struct {
                        W      int    `json:"w,string"`
                        H      int    `json:"h,string"`
                        Resize string `json:"resize"`
                    } `json:"small,string"`
                    Large struct {
                        W      int    `json:"w,string"`
                        H      int    `json:"h,string"`
                        Resize string `json:"resize"`
                    } `json:"large,string"`
                } `json:"sizes,string"`
                SourceStatusID    int    `json:"source_status_id,string"`
                SourceStatusIDStr string `json:"source_status_id_str"`
                Type              string `json:"type"`
                Indices           []int  `json:"indices,string"`
                VideoInfo         struct {
                    AspectRatio    []int `json:"aspect_ratio,string"`
                    DurationMillis int   `json:"duration_millis,string"`
                    Variants       []struct {
                        Bitrate     int    `json:"bitrate,string"`
                        ContentType string `json:"content_type"`
                        URL         string `json:"url"`
                    } `json:"variants,string"`
                } `json:"video_info,string"`
                ExtAltText string `json:"ext_alt_text"`
            } `json:"media,string"`
        } `json:"extended_entities,string"`
        ExtendedTweet struct {
            FullText         string      `json:"full_text"`
            DisplayTextRange interface{} `json:"display_text_range,string"`
            Entities         struct {
                Urls     interface{} `json:"urls,string"`
                Hashtags interface{} `json:"hashtags,string"`
                URL      struct {
                    Urls interface{} `json:"urls,string"`
                } `json:"url,string"`
                UserMentions interface{} `json:"user_mentions,string"`
                Media        interface{} `json:"media,string"`
            } `json:"entities,string"`
            ExtendedEntities struct {
                Urls     interface{} `json:"urls,string"`
                Hashtags interface{} `json:"hashtags,string"`
                URL      struct {
                    Urls interface{} `json:"urls,string"`
                } `json:"url,string"`
                UserMentions interface{} `json:"user_mentions,string"`
                Media        interface{} `json:"media,string"`
            } `json:"extended_entities,string"`
        } `json:"extended_tweet,string"`
        FavoriteCount        int    `json:"favorite_count,string"`
        Favorited            bool   `json:"favorited,string"`
        FilterLevel          string `json:"filter_level"`
        FullText             string `json:"full_text"`
        HasExtendedProfile   bool   `json:"has_extended_profile,string"`
        ID                   int64  `json:"id,string"`
        IDStr                string `json:"id_str"`
        InReplyToScreenName  string `json:"in_reply_to_screen_name"`
        InReplyToStatusID    int    `json:"in_reply_to_status_id,string"`
        InReplyToStatusIDStr string `json:"in_reply_to_status_id_str"`
        InReplyToUserID      int    `json:"in_reply_to_user_id,string"`
        InReplyToUserIDStr   string `json:"in_reply_to_user_id_str"`
        IsTranslationEnabled bool   `json:"is_translation_enabled,string"`
        Lang                 string `json:"lang"`
        Place                struct {
            Attributes  interface{} `json:"attributes,string"`
            BoundingBox struct {
                Coordinates interface{} `json:"coordinates,string"`
                Type        string      `json:"type"`
            } `json:"bounding_box,string"`
            ContainedWithin interface{} `json:"contained_within,string"`
            Country         string      `json:"country"`
            CountryCode     string      `json:"country_code"`
            FullName        string      `json:"full_name"`
            Geometry        struct {
                Coordinates interface{} `json:"coordinates,string"`
                Type        string      `json:"type"`
            } `json:"geometry,string"`
            ID        string      `json:"id"`
            Name      string      `json:"name"`
            PlaceType string      `json:"place_type"`
            Polylines interface{} `json:"polylines,string"`
            URL       string      `json:"url"`
        } `json:"place,string"`
        QuotedStatusID              int         `json:"quoted_status_id,string"`
        QuotedStatusIDStr           string      `json:"quoted_status_id_str"`
        QuotedStatus                interface{} `json:"quoted_status,string"`
        PossiblySensitive           bool        `json:"possibly_sensitive,string"`
        PossiblySensitiveAppealable bool        `json:"possibly_sensitive_appealable,string"`
        RetweetCount                int         `json:"retweet_count,string"`
        Retweeted                   bool        `json:"retweeted,string"`
        RetweetedStatus             interface{} `json:"retweeted_status,string"`
        Source                      string      `json:"source"`
        Scopes                      interface{} `json:"scopes,string"`
        Text                        string      `json:"text"`
        User                        struct {
            ContributorsEnabled bool   `json:"contributors_enabled,string"`
            CreatedAt           string `json:"created_at"`
            DefaultProfile      bool   `json:"default_profile,string"`
            DefaultProfileImage bool   `json:"default_profile_image,string"`
            Description         string `json:"description"`
            Email               string `json:"email"`
            Entities            struct {
                Urls     interface{} `json:"urls,string"`
                Hashtags interface{} `json:"hashtags,string"`
                URL      struct {
                    Urls interface{} `json:"urls,string"`
                } `json:"url,string"`
                UserMentions interface{} `json:"user_mentions,string"`
                Media        interface{} `json:"media,string"`
            } `json:"entities,string"`
            FavouritesCount                int         `json:"favourites_count,string"`
            FollowRequestSent              bool        `json:"follow_request_sent,string"`
            FollowersCount                 int         `json:"followers_count,string"`
            Following                      bool        `json:"following,string"`
            FriendsCount                   int         `json:"friends_count,string"`
            GeoEnabled                     bool        `json:"geo_enabled,string"`
            HasExtendedProfile             bool        `json:"has_extended_profile,string"`
            ID                             int64       `json:"id,string"`
            IDStr                          string      `json:"id_str"`
            IsTranslator                   bool        `json:"is_translator,string"`
            IsTranslationEnabled           bool        `json:"is_translation_enabled,string"`
            Lang                           string      `json:"lang"`
            ListedCount                    int         `json:"listed_count,string"`
            Location                       string      `json:"location"`
            Name                           string      `json:"name"`
            Notifications                  bool        `json:"notifications,string"`
            ProfileBackgroundColor         string      `json:"profile_background_color"`
            ProfileBackgroundImageURL      string      `json:"profile_background_image_url"`
            ProfileBackgroundImageURLHTTPS string      `json:"profile_background_image_url_https"`
            ProfileBackgroundTile          bool        `json:"profile_background_tile,string"`
            ProfileBannerURL               string      `json:"profile_banner_url"`
            ProfileImageURL                string      `json:"profile_image_url"`
            ProfileImageURLHTTPS           string      `json:"profile_image_url_https"`
            ProfileLinkColor               string      `json:"profile_link_color"`
            ProfileSidebarBorderColor      string      `json:"profile_sidebar_border_color"`
            ProfileSidebarFillColor        string      `json:"profile_sidebar_fill_color"`
            ProfileTextColor               string      `json:"profile_text_color"`
            ProfileUseBackgroundImage      bool        `json:"profile_use_background_image,string"`
            Protected                      bool        `json:"protected,string"`
            ScreenName                     string      `json:"screen_name"`
            ShowAllInlineMedia             bool        `json:"show_all_inline_media,string"`
            Status                         interface{} `json:"status,string"`
            StatusesCount                  int         `json:"statuses_count,string"`
            TimeZone                       string      `json:"time_zone"`
            URL                            string      `json:"url"`
            UtcOffset                      int         `json:"utc_offset,string"`
            Verified                       bool        `json:"verified,string"`
            WithheldInCountries            interface{} `json:"withheld_in_countries,string"`
            WithheldScope                  string      `json:"withheld_scope"`
        } `json:"user,string"`
        WithheldCopyright   bool        `json:"withheld_copyright,string"`
        WithheldInCountries interface{} `json:"withheld_in_countries,string"`
        WithheldScope       string      `json:"withheld_scope"`
    } `json:"tweets,string"`
}

type Data struct {
  Id int `dynamo:"id"`
  SearchWord string `dynamo:"search_word"`
  TweetData string `dynamo:"tweet_data"`
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
        word = "cat"
    }

    // DynamoDBへ接続
    db := dynamo.New(session.New(), &aws.Config{
        Region: aws.String("us-east-2"), // "ap-northeast-1"等
    })

    table := db.Table("Tweet")

    var datas []Data
    err := table.Get("search_word", word).Index("search_word_index").All(&datas)
    if err != nil {
        panic(err.Error())
    }

    var responses []Response
    var tweets []Tweet

    for _, data := range datas {
        err := json.Unmarshal([]byte(data.TweetData), &tweets)
        if err != nil {
            spew.Dump(tweets)
            fmt.Println(err)
        }

        for _, tweet := range tweets {
            spew.Dump(tweet)
            response := Response {
                //Id: tweet.IDStr,
                //UserId: tweet.User.ScreenName,
            }

            responses = append(responses, response)
        }
    }

    return responses, nil
}

func main() {
    lambda.Start(Handler)
}
