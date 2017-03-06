package main

import (
	"os"

	tweetstorm "github.com/andermalk16/tweetstorm/core"
)

func main() {

	text := os.Args[1]
	api := tweetstorm.TwiterAPI{}
	tweetstorm.Storm(text, api)

}
