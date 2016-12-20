package main

import (
	"errors"
	"net/url"
	"strconv"
	"fmt"
)

const (
	DEFAULT_SEARCH_SORT = "startTime"
	DEFAULT_SEARCH_TARGETS = "tags,title"
	DEFAULT_SEARCH_FIELDS = "contentId,title,tags,startTime"
	DEFAULT_SEARCH_LIMIT = 25
)

type SearchParameters struct {
	Targets string
	Fields  string
	Sort    string
	Limit   int
}

type MetaResponse struct {
	Code       int    `json:"status"`
	TotalCount int    `json:"totalCount"`
	Id         string `json:"id"`
}

type SearchHit struct {
	ContentId string     `json:"contentId"`
	Title     string     `json:"title"`
	tags      string     `json:"tags"`
	startTime string     `json:"startTime"`
}

type SearchResponse struct {
	MetaResponse    `json:"meta"`
	Hits			[]SearchHit `json:"data"`
}

func CreateSearchParameters() SearchParameters {
	return SearchParameters{
		Sort:     DEFAULT_SEARCH_SORT,
		Targets:  DEFAULT_SEARCH_TARGETS,
		Fields:   DEFAULT_SEARCH_FIELDS,
		Limit:    DEFAULT_SEARCH_LIMIT,
	}
}

type Client struct {
	debug bool
}

func New() *Client {
	s := &Client{}
	return s
}

func (client *Client) _search(path, query string, params SearchParameters) (response *SearchResponse, error error) {
	queries := url.Values{
		"q": {query},
	}
	queries.Add("_sort", params.Sort)
	queries.Add("_limit", strconv.Itoa(params.Limit))
	queries.Add("fields", params.Fields)
	queries.Add("targets", params.Targets)

	response = &SearchResponse{}
	if err := post(path, queries, response, client.debug); err != nil {
		return nil, err
	}
	if response.MetaResponse.Code != 200 {
		logger.Printf("%v\n", response.MetaResponse)
		return nil, errors.New("something wrong")
	}
	return response, nil
}

func (client *Client) Search(path, query string, params SearchParameters) (*SearchResponse, error) {
	if response, err := client._search(path, query, params); err != nil {
		return nil, err
	} else {
		return response, nil
	}
}

func main() {
	client := New()
	res, err := client.Search("api/v2/video/contents/search", "test", CreateSearchParameters())
	if err != nil {
		fmt.Printf("client err:%s\n", err)
		return
	}
	fmt.Printf("response: %v\n", res)
}
