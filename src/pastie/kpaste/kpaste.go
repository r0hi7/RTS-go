package kpaste

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/nishitm/RTS-go/config"
)

var urlMap = make(map[string]bool)

//KpasteImplement struct is used to implement the interface GetSearchedTerm
type KpasteImplement struct{}

// GetSearchedTerm method implementation for Kpaste
func (c KpasteImplement) GetSearchedTerm(configuration config.Config) {

	resp, err := http.Get(configuration.Kpaste.URL)
	if err != nil {
		log.Print(err)
		return
	}
	defer resp.Body.Close()
	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Print(err)
		return
	}
	r := regexp.MustCompile(configuration.Kpaste.Regex)
	matches := r.FindAllString(string(html), -1)
	matches = matches[:len(matches)-1]

	for i, str := range matches {
		matches[i] = configuration.Kpaste.URLBase + str[8:len(str)-2]
	}

	newMap := make(map[string]bool)
	for _, link := range matches {
		if len(urlMap) == 10 { //Since Kpaste is giving 10 entries at a time
			_, ok := urlMap[link]
			if ok {
				urlMap[link] = true
			} else {
				newMap[link] = false
			}
		} else {
			urlMap[link] = false
		}
	}

	if len(newMap) > 0 {
		for k := range urlMap {
			if urlMap[k] == false {
				delete(urlMap, k)
			}
		}
		for k := range newMap {
			urlMap[k] = false
		}
	} else {
		newMap = urlMap
	}

	if len(newMap) > 0 {
		for k := range newMap {
			if newMap[k] != true {
				found := CrawlAndSearch(k, configuration)
				if found {
					fmt.Println(k)
				}

			}
		}
	} else {
		for k := range urlMap {
			if urlMap[k] != true {
				found := CrawlAndSearch(k, configuration)
				if found {
					fmt.Println(k)
				}

			}
		}
	}

	for k := range urlMap {
		urlMap[k] = false
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
