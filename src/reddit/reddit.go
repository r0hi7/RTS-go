package reddit

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/nishitm/RTS-go/config"
	"github.com/nishitm/RTS-go/util"
)

type RedditImplement struct{}

// GetSearchedTerm method implementation for Reddit
func (r RedditImplement) GetSearchedTerm(configuration config.Config) {
	const UserAgent = "Golang Reddit Reader v1.0"
	for _, term := range configuration.Reddit.SearchTerms {
		req, err := http.NewRequest("GET", configuration.Reddit.URL, nil)
		if err != nil {
			log.Print(err)
			return
		}
		req.Header.Set("User-Agent", UserAgent)
		q := req.URL.Query()
		q.Add("q", term)
		q.Add("after", configuration.Reddit.Interval)
		req.URL.RawQuery = q.Encode()
		tr := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
		client := &http.Client{Transport: tr}
		resp, err := client.Do(req)
		if err != nil {
			log.Print("Errored when sending request to the server")
			return
		}
		defer resp.Body.Close()
		respBody, _ := ioutil.ReadAll(resp.Body)
		var resultStruct util.ChildrenDataStruct
		json.Unmarshal(respBody, &resultStruct)
		for _, result := range resultStruct.Data {
			fmt.Printf("Url: %s \n Selftext: %s\n\n", result.Url, result.SelfText)
		}
	}
}
