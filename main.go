package main

import (
	"fmt"
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
	fmt.Println("Hi")
	fmt.Println(cleanInput("        Hello world"))
}

