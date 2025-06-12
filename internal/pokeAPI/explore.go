package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"

	pokecache "github.com/zig-gy/pokedexcli/internal/pokeCache"
)

func FetchPokemonFromLocation(url string, cache *pokecache.Cache) ([]string, error) {
	body, ok := cache.Get(url)
	if !ok {
		res, err := http.Get(url)
		if err != nil {
			return []string{}, err
		}
		defer res.Body.Close()
		body, err = io.ReadAll(res.Body)
		if err != nil {
			return []string{}, err
		}
		cache.Add(url, body)
	}

	var locInfo LocationInfo
	if err := json.Unmarshal(body, &locInfo); err != nil {
		return []string{}, err
	}

	var names []string
	for _, encounter := range locInfo.PokemonEncounters {
		names = append(names, encounter.Pokemon.Name)
	}
	return names, nil
}

type LocationInfo struct {
    PokemonEncounters []PokemonEncounter `json:"pokemon_encounters"`
}

type PokemonEncounter struct {
    Pokemon PokemonLoc `json:"pokemon"`
}

type PokemonLoc struct {
    Name string `json:"name"`
    URL  string `json:"url"`
}