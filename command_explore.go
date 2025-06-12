package main

import (
	"fmt"

	pokeapi "github.com/zig-gy/pokedexcli/internal/pokeAPI"
)

func commandExplore(cfg *config, arguments []string) (err error) {
	if len(arguments) < 1 {
		fmt.Println("You need to pass an argument for this command")
		return fmt.Errorf("you need to pass an argument for this comamnd")
	}
	
	argument := arguments[0]
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s/", argument)

	pokemon, err := pokeapi.FetchPokemonFromLocation(url, cfg.cache)
	if err!= nil {
		return
	}

	for _, name := range pokemon {
		fmt.Println(name)
	}
	
	return
}
