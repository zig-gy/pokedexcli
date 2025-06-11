package main

import (
	"fmt"

	pokeapi "github.com/zig-gy/pokedexcli/internal/pokeAPI"
)

func commandExplore(cfg *config, arguments []string) (err error) {
	argument := arguments[0]
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s/", argument)

	pokemon, err := pokeapi.FetchPokemon(url, cfg.cache)
	if err!= nil {
		return
	}

	for _, name := range pokemon {
		fmt.Println(name)
	}
	
	return
}
