package pokeapi

import (
	"encoding/json"
	"net/http"
)

const (
	baseURL = "https://pokeapi.co/api/v2"
)

type Locations struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous,omitempty"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func FetchLocations(url string) (loc Locations, err error) {
	res, err := http.Get(url)
	if err != nil {
		return
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&loc)
	return
}
