package github

import "time"

const IssuesURL = "https://api.github.com/search/issues"

type IssuesSearchResult struct {
   TotalCount int `json:"total_count"`
   Items   []*Issues

}

type Issue struct {
   Number int
   HTMLURL string `json:"html_url"`
   Title   string
   State   string
   User    *User
   CreateAt  time.Time `json:"created_at"`
   Body    string   // in markdown format
}

type User struct {
   Login string
   HTMLURL string `json:"html_url"`
}
