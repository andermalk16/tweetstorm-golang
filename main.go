package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/coreos/pkg/flagutil"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

func main() {

	text := os.Args[1]

	storm(text)
}

func storm(text string) {

	tweets := generateTweetsText(text)
	log.Printf("=====================TWEETS=====================")
	doTweets(tweets)

}

func generateTweetsText(text string) []string {
	letterLimit := 135
	words := strings.Split(text, " ")
	var tweet string
	var tweets []string
	var beforeAddWord string
	var afterAddWord string

	for index := 0; index < len(words); index++ {
		beforeAddWord = strings.TrimSpace(tweet)
		afterAddWord = fmt.Sprintf(" %s %s ", strings.TrimSpace(tweet), strings.TrimSpace(words[index]))

		if len(afterAddWord) > letterLimit {
			tweets = append(tweets, beforeAddWord)
			tweet = words[index] + " "
		} else {
			tweet += words[index] + " "
		}
	}
	return append(tweets, afterAddWord)
}

func doTweets(tweets []string) {

	flags := flag.NewFlagSet("user-auth", flag.ExitOnError)
	consumerKey := flags.String("consumer-key", "", "Twitter Consumer Key")
	consumerSecret := flags.String("consumer-secret", "", "Twitter Consumer Secret")
	accessToken := flags.String("access-token", "", "Twitter Access Token")
	accessSecret := flags.String("access-secret", "", "Twitter Access Secret")

	flags.Parse(os.Args[1:])
	flagutil.SetFlagsFromEnv(flags, "TWITTER")

	if *consumerKey == "" || *consumerSecret == "" || *accessToken == "" || *accessSecret == "" {
		log.Fatal("Consumer key/secret and Access token/secret required")
	}

	config := oauth1.NewConfig(*consumerKey, *consumerSecret)
	token := oauth1.NewToken(*accessToken, *accessSecret)
	// OAuth1 http.Client will automatically authorize Requests
	httpClient := config.Client(oauth1.NoContext, token)
	// Twitter client
	client := twitter.NewClient(httpClient)

	for index := len(tweets) - 1; 0 <= index; index-- {
		tweet := fmt.Sprintf("%d/%d %s", index+1, len(tweets), strings.TrimSpace(tweets[index]))
		_, resp, err := client.Statuses.Update(tweet, nil)
		if err != nil {
			log.Fatal("ERROR!! " + err.Error())
		} else if resp.StatusCode != 200 {
			log.Fatal("ERROR!! Status: " + resp.Status)
		} else {
			log.Println(tweet)
		}
	}
}
