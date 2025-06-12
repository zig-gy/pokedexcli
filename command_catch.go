package main

import (
	"fmt"
	"math/rand"
	"time"

	pokeapi "github.com/zig-gy/pokedexcli/internal/pokeAPI"
)

func commandCatch(cfg *config, arguments []string) (err error) {
	if len(arguments) < 1 {
		fmt.Println("You need to pass an argument for this command")
		return fmt.Errorf("you need to pass an argument for this command")
	}
	pokemonName := arguments[0]
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s/", pokemonName)
	fmt.Printf("Throwing a Pokeball at %s", pokemonName)
	time.Sleep(250 * time.Millisecond)
	fmt.Print(".")
	time.Sleep(250 * time.Millisecond)
	fmt.Print(".")
	time.Sleep(250 * time.Millisecond)
	fmt.Print(".\n")
	time.Sleep(500 * time.Millisecond)

	pokemon, err := pokeapi.FetchPokemon(url, cfg.cache)
	if err != nil {
		fmt.Printf("Error found fetching: %v", err)
		return
	}

	if rand.Intn(pokemon.BaseExp + cfg.catchTreshold) < cfg.catchTreshold {
		fmt.Printf("%s was caught!\n", pokemonName)
		cfg.pokedex[pokemonName] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", pokemonName)
	}

	return
}