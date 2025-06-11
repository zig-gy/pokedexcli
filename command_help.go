package main

import "fmt"

func commandHelp(cfg *config) (err error) {
	message := `
Welcome to the Pokedex!
Usage:

`

	for _, v := range GetCommands() {
		message = fmt.Sprintf("%s%s: %s\n", message, v.name, v.description)
	}

	_, err = fmt.Println(message)
	return
}