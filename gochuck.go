package gochuck

import (
	"encoding/json"
	"net/http"
	"net/url"
	"time"
)

const (
	baseUrl            = "https://api.chucknorris.io"
	randomEndpoint     = "/jokes/random"
	categoriesEndpoint = "/jokes/categories"
	searchEndpoint     = "/jokes/search"
)

var client = http.Client{Timeout: 5 * time.Second}

type Fact struct {
	IconUrl  string   `json:"icon_url"`
	Id       string   `json:"id"`
	Url      string   `json:"url"`
	Value    string   `json:"value"`
	Category []string `json:"category"`
}

type FactCollection struct {
	Total   int    `json:"total"`
	Results []Fact `json:"result"`
}

func GetRandom() (*Fact, error) {
	reqUrl := baseUrl + randomEndpoint
	req, _ := http.NewRequest("GET", reqUrl, nil)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var fact Fact
	if err := json.NewDecoder(resp.Body).Decode(&fact); err != nil {
		return nil, err
	}

	return &fact, nil
}

func GetCategories() ([]string, error) {
	reqUrl := baseUrl + categoriesEndpoint
	req, _ := http.NewRequest("GET", reqUrl, nil)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var categories []string

	if err := json.NewDecoder(resp.Body).Decode(&categories); err != nil {
		return nil, err
	}

	return categories, nil
}

func GetByQuery(query string) (*FactCollection, error) {
	reqUrl := baseUrl + searchEndpoint
	req, _ := http.NewRequest("GET", reqUrl, nil)

	reqBody := url.Values{}
	reqBody.Set("query", query)
	req.URL.RawQuery = reqBody.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var collection FactCollection
	if err := json.NewDecoder(resp.Body).Decode(&collection); err != nil {
		return nil, err
	}

	return &collection, nil
}

func GetRandomByCategory(category string) (*Fact, error) {
	reqUrl := baseUrl + randomEndpoint
	req, _ := http.NewRequest("GET", reqUrl, nil)

	reqBody := url.Values{}
	reqBody.Set("category", category)
	req.URL.RawQuery = reqBody.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var fact Fact
	if err := json.NewDecoder(resp.Body).Decode(&fact); err != nil {
		return nil, err
	}

	return &fact, nil
}
