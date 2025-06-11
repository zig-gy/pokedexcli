package main

import (
	"fmt"

	pokeapi "github.com/zig-gy/pokedexcli/internal/PokeAPI"
)

func commandMap(next, prev string) (outNext, outPrev string, err error) {
	locs, err := pokeapi.FetchLocations(next)
	if err != nil {
		return
	}
	for _, loc := range locs.Results {
		fmt.Println(loc.Name)
	}
	return locs.Next, locs.Previous, err
}

func commandMapb(next, prev string) (outNext, outPrev string, err error) {
	var zeroString string
	if prev == zeroString {
		fmt.Println("Already on the first page")
		return "", "", fmt.Errorf("already on the first page")
	}

	locs, err := pokeapi.FetchLocations(prev)
	if err != nil {
		return
	}

	for _, loc := range locs.Results {
		fmt.Println(loc.Name)
	}

	return locs.Next, locs.Previous, err
}