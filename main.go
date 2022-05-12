package main

import (
	"log"
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/joho/godotenv"
)

var CONFIG = readConfig()

func main() {
	importEnv()
	client := twitterConnexion()
	stream := createStreaming(client, CONFIG.Principal)
	demux := twitter.NewSwitchDemux()
	demux.Tweet = processingTweet
	for message := range stream.Messages {
		log.Println("Received message")
		go demux.Handle(message)
	}

}

func importEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func twitterConnexion() *twitter.Client {
	config := oauth1.NewConfig(os.Getenv("consumerKey"), os.Getenv("consumerSecret"))
	token := oauth1.NewToken(os.Getenv("accessToken"), os.Getenv("accessSecret"))
	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)
	verifyParams := &twitter.AccountVerifyParams{
		SkipStatus:   twitter.Bool(true),
		IncludeEmail: twitter.Bool(true),
	}
	user, _, err := client.Accounts.VerifyCredentials(verifyParams)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Conected to user's account: %+v\n", user.ScreenName)
	return client
}

func createStreaming(client *twitter.Client, track []string) *twitter.Stream {
	params := &twitter.StreamFilterParams{
		Track:    track,
		Language: []string{"en"},
	}
	stream, err := client.Streams.Filter(params)
	if err != nil {
		log.Fatal("Use of the stream is impossible")
	}
	return stream
}

func processing() error {

	return nil
}
