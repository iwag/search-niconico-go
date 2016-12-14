package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"log"
	"os"
)

const SEARCH_API = "http://api.search.nicovideo.jp/"

var logger *log.Logger = log.New(os.Stdout, "search ", log.Lshortfile)

var HTTPClient = &http.Client{}

func parseJson(body io.ReadCloser, intf *interface{}, debug bool) error {
	response, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}

	if debug {
		logger.Printf("parseJson: %s\n", string(response))
	}

	if err = json.Unmarshal(response, &intf); err != nil {
		return err
	}

	return nil
}

func post(path string, values url.Values, intf interface{}, debug bool) error {
	resp, err := HTTPClient.PostForm(SEARCH_API + path, values)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return parseJson(resp.Body, &intf, debug)
}
