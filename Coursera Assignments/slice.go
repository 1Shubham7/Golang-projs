package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	slice := make([]int, 0, 3) // Create an empty integer slice with capacity 3

	for {
		fmt.Println("Enter an integer (or 'X' to exit):")
		var input string
		fmt.Scanln(&input)

		if input == "X" {
			break
		}

		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter an integer or 'X'.")
			continue
		}

		slice = append(slice, num)
		sort.Ints(slice)

		fmt.Println("Sorted slice:", slice)
	}
}
