package gochuck

import (
	"encoding/json"
	"net/http"
	"net/url"
)

const (
	baseUrl            = "https://api.chucknorris.io"
	randomEndpoint     = "/jokes/random"
	categoriesEndpoint = "/jokes/categories"
	searchEndpoint     = "/jokes/search"
)

func init() {
	SetClient(http.DefaultClient)
}

var client *http.Client

// SetClient allows injecting an *http.Client if needed,
// http.DefaultClient is used by default.
func SetClient(c *http.Client) {
	client = c
}

// Fact describes a chuck norris fact.
type Fact struct {
	IconUrl  string   `json:"icon_url"`
	Id       string   `json:"id"`
	Url      string   `json:"url"`
	Value    string   `json:"value"`
	Category []string `json:"category"`
}

// FactCollection is a collection of one or more facts with a total number.
type FactCollection struct {
	Total   int    `json:"total"`
	Results []Fact `json:"result"`
}

// GetRandom returns a random fact.
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

// GetCategories returns a list of fact categories.
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

// GetByQuery returns a fact with a text query.
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

// GetRandomByCategory returns a random fact by category.
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
