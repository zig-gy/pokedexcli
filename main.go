package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	var prev string
	next := "https://pokeapi.co/api/v2/location-area/"
	for {
		fmt.Print("Pokedex > ")
		s.Scan()
		input := s.Text()
		cInput := cleanInput(input)[0]
		val, ok := GetCommands()[cInput]
		if ok {
			next, prev, _ = val.callback(next, prev)
		} else {
			fmt.Println("Unknown command")
		}
	}
}

type cliCommand struct {
	name		string
	description	string
	callback	func(string, string) (string, string, error)
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
	}
	return supportedCommands
}
