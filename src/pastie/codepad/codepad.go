package codepad

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/nishitm/RTS-go/config"
)

var urlMap = make(map[string]bool)

type CodepadImplement struct{}

// GetSearchedTerm method implementation for Codepad
func (c CodepadImplement) GetSearchedTerm(configuration config.Config) {
	doc, err := goquery.NewDocument(configuration.Codepad.URL)
	if err != nil {
		log.Print(err)
		return
	}
	newMap := make(map[string]bool)
	doc.Find(".section .label a").Each(func(i int, s *goquery.Selection) {
		Link, _ := s.Attr("href")
		if len(urlMap) == 20 { //Since Codepad is giving 20 entries at a time
			_, ok := urlMap[Link]
			if ok {
				urlMap[Link] = true
			} else {
				newMap[Link] = false
			}
		} else {
			urlMap[Link] = false
		}
	})
	if len(newMap) > 0 {
		for k := range urlMap {
			if urlMap[k] == false {
				delete(urlMap, k)
			}
		}
		for k := range newMap {
			urlMap[k] = false
		}
	}
	for k := range urlMap {
		urlMap[k] = false
		found := CrawlAndSearch(k, configuration)
		if found {
			fmt.Println(k)
		}
	}
}

// CrawlAndSearch method will crawl individual link and search for the term
func CrawlAndSearch(url string, configuration config.Config) bool {
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
	found := false
	for _, term := range configuration.Codepad.SearchTerms {
		if strings.Contains(string(contents), term) {
			found = true
		}
	}
	if found {
		return true
	}
	return false
}
