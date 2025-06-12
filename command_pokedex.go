package main

import (
	"fmt"
)

func commandPokedex(cfg *config, arguments []string) (err error) {
	if len(cfg.pokedex) < 1 {
		fmt.Println("Your Pokedex is empty")
		return
	}

	fmt.Println("Your Pokedex")
	for _, pokemon := range cfg.pokedex {
		fmt.Printf("  - %s\n", pokemon.Name)
	}

	return
}