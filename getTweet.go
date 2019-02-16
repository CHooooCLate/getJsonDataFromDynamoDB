package main

import (
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
  Id int `dynamo:"id"`
  UserId string `dynamo:"user_id"`
  TweetId string `dynamo:"tweet_id"`
  WordId string `dynamo:"word_id"`
}

type Response struct {
    Id int `json:"id"`
    UserId string `json:"user_id"`
    WordId string `json:"word_id"`
}

func Handler() (Response, error) {
    // DynamoDBへ接続
    db := dynamo.New(session.New(), &aws.Config{
        Region: aws.String("us-east-2"), // "ap-northeast-1"等
    })

    table := db.Table("Tweet")

    var tweet []Tweet
    err := table.Scan().All(&tweet)
    if err != nil {
        panic(err.Error())
    }

    spew.Dump(tweet)

    var response Response

    return response, nil
}

func main() {
    lambda.Start(Handler)
}
