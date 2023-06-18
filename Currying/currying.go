package main

import "fmt"

func main() {
	squareFunc := selfMath(multiply)
	doubleFunc := selfMath(add)
	
	fmt.Println(squareFunc(5))
	fmt.Println(doubleFunc(5))
  }
  
  func multiply(x, y int) int {
	return x * y
  }
  
  func add(x, y int) int {
	return x + y
  }
  
  func selfMath(mathFunc func(int, int) int) func (int) int {
	return func(x int) int {
	  return mathFunc(x, x)
	}
  }