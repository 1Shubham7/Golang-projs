package main

import "fmt"

func main() {
	type circle struct {
		radius int
	}

	func (c circle) area() int {
		return 2 * 3 * c.radius
	}

	c := circle{
		radius: 10,
	}

	fmt.Println(c.area())

}

