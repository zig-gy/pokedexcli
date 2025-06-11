package main

import (
	"fmt"
	"os"
)

func commandExit(cfg *config, arguments []string) (err error) {
	_, err = fmt.Println("Closing the Pokedex... Goodbye!")
	if err != nil {
		return
	}
	os.Exit(0)
	return
}