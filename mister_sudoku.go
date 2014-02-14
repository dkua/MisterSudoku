package main

import (
	"bot"
	"fmt"
	"net/url"
)

func main() {
	api := bot.GetTwitterApi()
	puzzle := bot.CreateTweet()
	tweet, err := api.PostTweet(puzzle, url.Values{})
	if err != nil {
		fmt.Printf("%v\n%v", "Something went wrong", err)
	}
	fmt.Println(tweet)
}
