package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	pokeapi "github.com/zig-gy/pokedexcli/internal/pokeAPI"
	pokecache "github.com/zig-gy/pokedexcli/internal/pokeCache"
)

func main() {
	cache := pokecache.NewCache(5 * time.Second)

	cfg := &config{
		nextLocationsURL: "https://pokeapi.co/api/v2/location-area/",
		cache: cache,
		catchTreshold: 50,
		pokedex: make(map[string]pokeapi.Pokemon),
	}

	s := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		s.Scan()
		input := s.Text()
		words := cleanInput(input)
		if len(words) < 1 {
			continue
		}
		cInput := words[0]
		val, ok := GetCommands()[cInput]
		if ok {
			val.callback(cfg, words[1:])
		} else {
			fmt.Println("Unknown command")
		}
	}
}

type cliCommand struct {
	name		string
	description	string
	callback	func(*config, []string) error
}

type config struct {
	nextLocationsURL	string
	prevLocationsURL	string
	cache				*pokecache.Cache
	catchTreshold		int
	pokedex				map[string]pokeapi.Pokemon
}

func cleanInput(text string) []string {
	words := strings.Fields(text)
	var finalWords []string
	for _, word := range words {
		finalWords = append(finalWords, strings.ToLower(word))
	}

	return finalWords
}

func GetCommands() map[string]cliCommand {
	supportedCommands := map[string]cliCommand {
	"exit": {
		name: 			"exit",
		description:	"Exit the Pokedex",
		callback:		commandExit,
		},
	"help": {
		name: "help",
		description: "Displays a help message",
		callback: commandHelp,
		},
	"map" : {
		name: "map",
		description: "Displays the name of 20 locations",
		callback: commandMap,
		},
	"mapb" : {
		name: "mapb",
		description: "Displays the name of the previous 20 locations",
		callback: commandMapb,
		},
	"explore" : {
		name: "explore",
		description: "Displays the names of the possible encounters in the area passed as argument",
		callback: commandExplore,
		},
	"catch" : {
		name: "catch",
		description: "Catch a Pokemon passed as an argument",
		callback: commandCatch,
		},
	"inspect" : {
		name: "inspect",
		description: "Get information about a pokemon you have in your pokedex",
		callback: commandInspect,
		},
	"pokedex" : {
		name: "pokedex",
		description: "Get all of the pokemon stored in the pokedex",
		callback: commandPokedex,
		},
	}
	return supportedCommands
}
