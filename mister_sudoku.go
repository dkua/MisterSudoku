package main

import (
	"bot"
	"fmt"
	"net/url"
	"time"
)

func main() {
	api := bot.GetTwitterApi()
	for {
		puzzle := bot.CreateTweet()
		tweet, err := api.PostTweet(puzzle, url.Values{})
		if err != nil {
			fmt.Printf("%v\n%v", "Something went wrong", err)
		}
		fmt.Println(tweet)
		time.Sleep(600 * time.Second)
	}
}
