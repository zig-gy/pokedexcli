package main

import "fmt"

func commandInspect(cfg *config, arguments []string) (err error) {
	if len(arguments) < 1 {
		fmt.Println("You need to pass an argument for this command")
		return fmt.Errorf("you need to pass an argument for this command")
	}

	pokemonName := arguments[0]
	pokemon, ok := cfg.pokedex[pokemonName]
	if !ok {
		fmt.Printf("Pokemon '%s' not found in your pokedex\n", pokemonName)
		return
	}

	fmt.Printf("Name:\t%s\n", pokemon.Name)
	fmt.Printf("Height:\t%d cm\n", pokemon.Height * 10)
	fmt.Printf("Weight:\t%d g\n",pokemon.Weight * 100)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  - %s: %d\n", stat.Info.Name, stat.Value)
	}
	fmt.Println("Types:")
	for _, pType := range pokemon.Types {
		fmt.Printf("  - %s\n", pType.Info.Name)
	}

	return
}