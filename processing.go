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

func (d Configuration) processingTweet(tweet *twitter.Tweet) {
	postBody, err := json.Marshal(CustomTweet{
		Name:     d.Name,
		KeyWords: d.KeyWords,
		Tweet:    *tweet,
	})
	if err != nil {
		log.Println("Can't marshal tweet")
		return
	}
	responseBody := bytes.NewBuffer(postBody)
	resp, err := http.Post(d.CallBackUrl, "application/json", responseBody)
	if err != nil {
		log.Println("Can't post tweet", err)
		return
	}
	log.Println("Call back response : ", resp.StatusCode)
	return
}
