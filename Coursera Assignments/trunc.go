package main

import (
	"fmt"
	"math"
)

func main() {
	num1 := 12.3456789
	num2 := 9.87654321

	fmt.Printf("Original Number 1: %.8f\n", num1)
	fmt.Printf("Original Number 2: %.8f\n", num2)

	// Truncating the numbers to two decimal places
	truncatedNum1 := math.Trunc(num1*100) / 100
	truncatedNum2 := math.Trunc(num2*100) / 100

	fmt.Printf("Truncated Number 1: %.2f\n", truncatedNum1)
	fmt.Printf("Truncated Number 2: %.2f\n", truncatedNum2)
}
