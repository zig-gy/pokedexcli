package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	words := strings.Fields(text)
	var finalWords []string
	for _, word := range words {
		finalWords = append(finalWords, strings.ToLower(word))
	}

	return finalWords
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		s.Scan()
		input := s.Text()
		cInput := cleanInput(input)
		fmt.Printf("Your command was: %s\n", cInput[0])
	}
}

