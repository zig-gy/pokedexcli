package main

import (
	"fmt"

	pokeapi "github.com/zig-gy/pokedexcli/internal/pokeAPI"
)

func commandMap(cfg *config, arguments []string) (err error) {
	next := cfg.nextLocationsURL
	cache := cfg.cache

	var zeroString string
	if next == zeroString {
		fmt.Println("Already on the last page")
		return fmt.Errorf("already on the last page")
	}

	locs, err := pokeapi.FetchLocations(next, cache)
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

func commandMapb(cfg *config, arguments []string) (err error) {
	prev := cfg.prevLocationsURL
	cache := cfg.cache

	var zeroString string
	if prev == zeroString {
		fmt.Println("Already on the first page")
		return fmt.Errorf("already on the first page")
	}

	locs, err := pokeapi.FetchLocations(prev, cache)
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