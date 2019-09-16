package util

type ResultStruct struct {
	Kind string     `json: kind`
	Data DataStruct `json: data`
}

type DataStruct struct {
	After    string               `json: after`
	Dist     int                  `json: dist`
	Children []ChildrenDataStruct `json: children`
}

type ChildrenStruct struct {
	Data []ChildrenDataStruct
}

type ChildrenDataStruct struct {
	Data []RedditPostStruct `json: data`
}

type RedditPostStruct struct {
	SelfText       string `json: selftext`
	AuthorFullName string `json: author_fullname`
	Title          string `json: title`
	Author         string `json: author`
	Url            string `json: url`
}
