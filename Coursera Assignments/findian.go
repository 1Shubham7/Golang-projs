package main

import (
	"fmt"
	"strings"
)

func main() {
	input := getInput() // Get user input
	if containsCharacters(input, "ian") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}

func getInput() string {
	var input string
	fmt.Print("Enter a string: ")
	fmt.Scanln(&input)
	return input
}

func containsCharacters(input, characters string) bool {
	for _, char := range characters {
		if !strings.ContainsRune(input, char) {
			return false
		}
	}
	return true
}
