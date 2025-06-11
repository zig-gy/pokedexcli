package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"

	pokecache "github.com/zig-gy/pokedexcli/internal/pokeCache"
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

func FetchLocations(url string, cache *pokecache.Cache) (loc Locations, err error) {
	body, ok := cache.Get(url)
	if !ok {
		res, err := http.Get(url)
		if err != nil {
			return Locations{}, err
		}
		defer res.Body.Close()
		body, err = io.ReadAll(res.Body)
		if err != nil {
			return Locations{}, err
		}
		cache.Add(url, body)

	}
	
	if err = json.Unmarshal(body, &loc); err != nil {
		return Locations{}, err
	}
	return
}
