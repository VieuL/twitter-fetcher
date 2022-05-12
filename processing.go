package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/dghubble/go-twitter/twitter"
)

type CustomTweet struct {
	Name     string
	KeyWords []string
	twitter.Tweet
}

func processingTweet(tweet *twitter.Tweet) {
	postBody, err := json.Marshal(CustomTweet{
		Name:     CONFIG.Name,
		KeyWords: CONFIG.KeyWords,
		Tweet:    *tweet,
	})
	if err != nil {
		log.Println("Can't marshal tweet")
		return
	}
	responseBody := bytes.NewBuffer(postBody)
	resp, err := http.Post(CONFIG.CallBackUrl, "application/json", responseBody)
	if err != nil {
		log.Println("Can't post tweet", err)
		return
	}
	log.Println("Call back response : ", resp.StatusCode)
	return
}
