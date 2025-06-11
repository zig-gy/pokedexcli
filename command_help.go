package main

import "fmt"

func commandHelp(next, prev string) (outNext, outPrev string, err error) {
	outNext = next
	outPrev = prev
	
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