package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		s.Scan()
		input := s.Text()
		cInput := cleanInput(input)[0]
		val, ok := GetCommands()[cInput]
		if ok {
			val.callback()
		} else {
			fmt.Println("Unknown command")
		}
	}
}

type cliCommand struct {
	name		string
	description	string
	callback	func() error
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
	}
	return supportedCommands
}
