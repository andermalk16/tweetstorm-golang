package tweetstorm

import (
	"flag"
	"log"
	"os"

	"github.com/coreos/pkg/flagutil"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

//TwitterCredentials struct
type TwitterCredentials struct {
	consumerKey    *string
	consumerSecret *string
	accessToken    *string
	accessSecret   *string
}

//TwiterAPI struct
type TwiterAPI struct {
}

//TwiterAPIer interface
type TwiterAPIer interface {
	GetTwitterCredentials() TwitterCredentials
	Tweet(text string) bool
}

//Tweet to call api tweet method
func (api TwiterAPI) Tweet(text string) bool {
	cred := api.GetTwitterCredentials()

	if *cred.consumerKey == "" || *cred.consumerSecret == "" || *cred.accessToken == "" || *cred.accessSecret == "" {
		log.Fatal("Consumer key/secret and Access token/secret required")
	}
	config := oauth1.NewConfig(*cred.consumerKey, *cred.consumerSecret)
	token := oauth1.NewToken(*cred.accessToken, *cred.accessSecret)
	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)
	_, resp, err := client.Statuses.Update(text, nil)

	if err != nil {
		log.Println("ERROR!! " + err.Error())
		return false
	} else if resp.StatusCode != 200 {
		log.Println("ERROR!! Status: " + resp.Status)
		return false
	} else {
		return true
	}

}

//GetTwitterCredentials method
func (api TwiterAPI) GetTwitterCredentials() TwitterCredentials {
	flags := flag.NewFlagSet("user-auth", flag.ExitOnError)
	consumerKey := flags.String("consumer-key", "", "Twitter Consumer Key")
	consumerSecret := flags.String("consumer-secret", "", "Twitter Consumer Secret")
	accessToken := flags.String("access-token", "", "Twitter Access Token")
	accessSecret := flags.String("access-secret", "", "Twitter Access Secret")

	flags.Parse(os.Args[1:])
	flagutil.SetFlagsFromEnv(flags, "TWITTER")
	return TwitterCredentials{consumerKey, consumerSecret, accessToken, accessSecret}

}
