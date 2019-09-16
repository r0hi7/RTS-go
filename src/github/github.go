package github

import (
	"context"
	"fmt"

	"github.com/google/go-github/github"
	"github.com/nishitm/RTS-go/config"
)

//GitImplement struct is to implement interface
type GitImplement struct{}

var gitMap = make(map[string]bool)

// GetSearchedTerm method implementation for Github
func (g GitImplement) GetSearchedTerm(configuration config.Config) {
	client := github.NewClient(nil)
	opt := &github.SearchOptions{}
	for _, term := range configuration.Github.SearchTerms {
		repos, _, _ := client.Search.Repositories(context.Background(), term, opt)
		for _, repo := range repos.Repositories {
			_, ok := gitMap[*repo.GitURL]
			if !ok {
				fmt.Printf("Name: %s\nUrl: %s\n\n", *repo.FullName, *repo.GitURL)
				gitMap[*repo.GitURL] = true
			}
		}
	}
}
