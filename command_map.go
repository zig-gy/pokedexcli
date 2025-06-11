package main

import (
	"fmt"

	pokeapi "github.com/zig-gy/pokedexcli/internal/pokeAPI"
)

func commandMap(cfg *config) (err error) {
	next := cfg.nextLocationsURL

	var zeroString string
	if next == zeroString {
		fmt.Println("Already on the last page")
		return fmt.Errorf("already on the last page")
	}

	locs, err := pokeapi.FetchLocations(next)
	if err != nil {
		return
	}
	for _, loc := range locs.Results {
		fmt.Println(loc.Name)
	}
	cfg.nextLocationsURL = locs.Next
	cfg.prevLocationsURL = locs.Previous

	return
}

func commandMapb(cfg *config) (err error) {
	prev := cfg.prevLocationsURL
	var zeroString string
	if prev == zeroString {
		fmt.Println("Already on the first page")
		return fmt.Errorf("already on the first page")
	}

	locs, err := pokeapi.FetchLocations(prev)
	if err != nil {
		return
	}

	for _, loc := range locs.Results {
		fmt.Println(loc.Name)
	}
	cfg.nextLocationsURL = locs.Next
	cfg.prevLocationsURL = locs.Previous

	return err
}