package tweetstorm

import (
	"fmt"
	"log"
	"strings"
)

//Storm creates tweetstorm
func Storm(text string, api TwiterAPIer) {

	tweets := GenerateTweetsText(text)
	log.Printf("=====================TWEETS=====================")
	DoTweets(tweets, api)
	log.Printf("=====================DONE=====================")

}

//GenerateTweetsText method
func GenerateTweetsText(text string) []string {

	letterLimit := 135
	words := strings.Split(text, " ")

	var tweet string
	var tweets []string
	var beforeAddWord string
	var afterAddWord string

	for index := 0; index < len(words); index++ {

		beforeAddWord = strings.TrimSpace(tweet)
		afterAddWord = fmt.Sprintf("%s %s ", strings.TrimSpace(tweet), strings.TrimSpace(words[index]))

		size := len([]rune(afterAddWord))
		if size > letterLimit {
			tweets = append(tweets, beforeAddWord)
			tweet = words[index] + " "
		} else {
			tweet += words[index] + " "
		}

	}

	return append(tweets, afterAddWord)
}

//DoTweets send tweet
func DoTweets(tweets []string, api TwiterAPIer) {

	for index := len(tweets) - 1; 0 <= index; index-- {

		tweet := fmt.Sprintf("%d/%d %s", index+1, len(tweets), strings.TrimSpace(tweets[index]))
		result := api.Tweet(tweet)
		if result {
			log.Println(tweet)
		} else {
			log.Println("Not possible send tweet")
			break
		}
	}
}
