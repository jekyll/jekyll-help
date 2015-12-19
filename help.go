package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

var (
	apiKey      = "FIX-ME"
	apiUrl      = "https://talk.jekyllrb.com"
	apiUsername = "jekyllbot"
)

type Issue struct {
	Id        int64     `json:"id"`
	Url       string    `json:"url"`
	Title     string    `json:"title"`
	CreatedBy string    `json:"created_by"`
	CreatedAt string    `json:"created_at"`
	Body      string    `json:"body"`
	State     string    `json:"state"`
	Comments  []Comment `json:"comments"`
	Quicklink string    `json:"quicklink"`
}

func (i Issue) String() string {
	return formattedQuotable(i)
}

func (i Issue) RefUrl() string {
	return i.Url
}

func (i Issue) CreationUser() string {
	return i.CreatedBy
}

func (i Issue) CreationTime() string {
	return i.CreatedAt
}

func (i Issue) Message() string {
	return i.Body
}

type Comment struct {
	Url       string `json:"url"`
	HtmlUrl   string `json:"html_url"`
	Id        int64  `json:"id"`
	User      User   `json:"user"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Body      string `json:"body"`
}

func (c Comment) String() string {
	return formattedQuotable(c)
}

func (c Comment) RefUrl() string {
	return c.HtmlUrl
}

func (c Comment) CreationUser() string {
	return c.User.Login
}

func (c Comment) CreationTime() string {
	return c.CreatedAt
}

func (c Comment) Message() string {
	return c.Body
}

type User struct {
	Login   string `json:"login"`
	HtmlUrl string `json"html_url"`
}

func (u User) String() string {
	return fmt.Sprintf("[%s](%s)", u.Login, u.HtmlUrl)
}

type Quotable interface {
	CreationUser() string
	CreationTime() string
	RefUrl() string
	Message() string
}

func formattedQuotable(q Quotable) string {
	return fmt.Sprintf(
		"[%s](%s) wrote on [%s](%s):\n\n%s",
		q.CreationUser(),
		"https://github.com/"+q.CreationUser(),
		q.CreationTime(),
		q.RefUrl(),
		quotedBody(q))
}

func quotedBody(q Quotable) string {
	insideBacktick := false
	outputLines := []string{}
	lines := strings.Split(q.Message(), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "```") {
			log.Println("oh, i see ur backticks", line)
			if insideBacktick {
				insideBacktick = false
				outputLines = append(outputLines, line+"\n")
			} else {
				insideBacktick = true
				outputLines = append(outputLines, "\n"+line)
			}
			continue
		}
		if insideBacktick {
			outputLines = append(outputLines, line)
		} else {
			outputLines = append(outputLines, "> "+line)
		}
	}
	return strings.Join(outputLines, "\n")
}

type TopicCreateResponse struct {
	TopicId int64 `json:"topic_id"`
}

func createTopic(i Issue) (*TopicCreateResponse, error) {
	params := url.Values{
		"api_key":      {apiKey},
		"api_username": {apiUsername},
		"title":        {i.Title},
		"raw":          {i.String()},
		"category":     {"jekyll-help-import"},
	}
	fmt.Println(params)
	resp, err := http.PostForm(apiUrl+"/posts", params)

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	log.Println(string(body))

	topicInfo := TopicCreateResponse{}
	err = json.Unmarshal(body, &topicInfo)
	if err != nil {
		return nil, err
	}
	log.Println(topicInfo)

	return &topicInfo, err
}

func createComment(c Comment, topicId int64) error {
	params := url.Values{
		"api_key":      {apiKey},
		"api_username": {apiUsername},
		"raw":          {c.String()},
		"topic_id":     {strconv.FormatInt(topicId, 10)},
	}
	fmt.Println(params)
	resp, err := http.PostForm(apiUrl+"/posts", params)

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return err
	}
	log.Println(string(body))

	return nil
}

func main() {
	var err error
	b, err := ioutil.ReadFile("comments.json")
	if err != nil {
		log.Fatalf("error reading comments.json: %v", err)
	}

	var issues []Issue
	err = json.Unmarshal(b, &issues)
	if err != nil {
		log.Fatalf("error unmarshalling issues: %v", err)
	}

	for _, i := range issues {
		log.Printf("Posting issue %d", i.Id)
		log.Printf("%s", i.RefUrl())
		topicInfo, err := createTopic(i)
		if topicInfo != nil {
			if topicInfo.TopicId == 0 {
				log.Fatalln("Got a 0 topic ID back. Flipping out.")
			} else {
				log.Printf("Created topic #%s", strconv.FormatInt(topicInfo.TopicId, 10))
			}
		}
		if err != nil {
			log.Printf("Got some error creating the topic (%v): %v", topicInfo, err)
			return
		}
		for _, c := range i.Comments {
			log.Printf("Posting comment %d for issue %d", c.Id, i.Id)
			log.Printf("%s", c.RefUrl())
			createComment(c, topicInfo.TopicId)
		}
		log.Println("--------------------------------------------------")
	}
}
