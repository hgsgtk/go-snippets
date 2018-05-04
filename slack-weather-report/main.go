package main

import (
	"github.com/nlopes/slack"
	"github.com/joho/godotenv"
	"os"
	"fmt"
	"net/url"
	"net/http"
	"encoding/json"
)

var YAHOO_API string = "https://map.yahooapis.jp/weather/V1/place"

type Sample struct {
	Name string
	Age int
}

type Samples []Sample

func main() {
	dog := Sample {
		Name: "わんわん",
		Age: 3,
	}
	cat := Sample {
		Name: "にゃんにゃん",
		Age: 8,
	}
	var animals Samples

	animals = append(animals, dog)
	animals = append(animals, cat)
	fmt.Println(animals)

	b, _ := json.Marshal(animals)
	fmt.Printf("%s\n", b)

	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	// get slack client
	accessToken := os.Getenv("SLACK_TOKEN")
	channel := os.Getenv("SLACK_CHANNEL")
	api := slack.New(accessToken)

	// get whether information
	appID := os.Getenv("YAHOO_APP_ID")
	values := url.Values{}
	values.Add("appid", appID)
	values.Add("coordinates", "35.759411,139.658992")
	values.Add("output", "json")

	res, err := http.Get(YAHOO_API + "?" + values.Encode())
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	fmt.Print(res.Body)

	params := slack.PostMessageParameters{}
	attachment := slack.Attachment{
		Pretext: "test pretext",
		Text: "test text",
	}
	params.Attachments = []slack.Attachment{attachment}
	channelID, timestamp, err := api.PostMessage(channel, "first post", params)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Message successfully sent to channel %s at %s\n", channelID, timestamp)
}