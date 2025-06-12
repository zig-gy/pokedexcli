package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"

	pokecache "github.com/zig-gy/pokedexcli/internal/pokeCache"
)

func FetchPokemon(url string, cache *pokecache.Cache) (Pokemon, error) {
	body, ok := cache.Get(url)
	if !ok {
		res, err := http.Get(url)
		if err != nil {
			return Pokemon{}, err
		}
		defer res.Body.Close()
		body, err = io.ReadAll(res.Body)
		if err != nil {
			return Pokemon{}, err
		}
		cache.Add(url, body)
	}

	var pokemon Pokemon
	err := json.Unmarshal(body, &pokemon)
	return pokemon, err
}

type Pokemon struct {
	Name string 	`json:"name"`
	BaseExp int 	`json:"base_experience"`
	Height int		`json:"height"`
	Weight int		`json:"weight"`
	Stats []Stat	`json:"stats"`
	Types []Type	`json:"types"`
}

type Stat struct {
	Value int		`json:"base_stat"`
	Info StatInfo	`json:"stat"`
}

type StatInfo struct {
	Name string	`json:"name"`
}

type Type struct {
	Info TypeInfo	`json:"type"`
}

type TypeInfo struct {
	Name string	`json:"name"`
}