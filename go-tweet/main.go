package main

import (
    "github.com/ChimeraCoder/anaconda"
    "fmt"
    "github.com/joho/godotenv"
    "os"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        panic(err)
    }

    accessToken := os.Getenv("ACCESS_TOKEN")
    accessTokenSecret := os.Getenv("ACCESS_TOKEN_SECRET")
    consumerKey := os.Getenv("CONSUMER_KEY")
    consumerSecret := os.Getenv("CONSUMER_SECRET")
    api := anaconda.NewTwitterApiWithCredentials(accessToken, accessTokenSecret, consumerKey, consumerSecret)

    text := "First testing tweet via golang api client."
    tweet, err := api.PostTweet(text, nil)
    if (err != nil) {
        panic(err)
    }
    fmt.Println(tweet.Text)
}
