package main

import (
	"fmt"
	"os"
)

func commandExit(next, prev string) (outNext, outPrev string, err error) {
	outNext = next
	outPrev = prev

	_, err = fmt.Println("Closing the Pokedex... Goodbye!")
	if err != nil {
		return
	}
	os.Exit(0)
	return
}