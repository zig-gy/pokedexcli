package main

import (
	"fmt"
	"os"
)

func commandExit(cfg *config) (err error) {
	_, err = fmt.Println("Closing the Pokedex... Goodbye!")
	if err != nil {
		return
	}
	os.Exit(0)
	return
}