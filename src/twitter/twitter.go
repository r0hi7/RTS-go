package twitter

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/nishitm/RTS-go/config"
)

type TwitterImplement struct{}

// GetSearchedTerm method implementation for Twitter
func (t TwitterImplement) GetSearchedTerm(configuration config.Config) {
	config := oauth1.NewConfig(configuration.Twitter.ConsumerKey, configuration.Twitter.ConsumerSecret)
	token := oauth1.NewToken(configuration.Twitter.AccessToken, configuration.Twitter.AccessSecret)
	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)
	params := &twitter.StreamFilterParams{
		Track:         configuration.Twitter.SearchTerms,
		StallWarnings: twitter.Bool(true),
	}
	stream, err := client.Streams.Filter(params)
	if err != nil {
		log.Print(err)
		return
	}
	demux := twitter.NewSwitchDemux()
	demux.Tweet = func(tweet *twitter.Tweet) {
		fmt.Println(tweet.Text)
		//		fmt.Println(reflect.TypeOf(tweet.Place.URL))
	}
	go demux.HandleChan(stream.Messages)
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	log.Println(<-ch)
	stream.Stop()
}
