package main
import (
	"fmt"
	"first-app/operations"
)

func main() {
		fmt.Println("This is a calculator")
		fmt.Println("Enter the first number : ")
		var numberOne int
		fmt.Scan(&numberOne)
		var numberTwo int
		fmt.Println("Enter the second number : ")
		fmt.Scan(&numberTwo)
		fmt.Printf("The first number is %v and the second number is %v", numberOne, numberTwo)

		var operation string
		fmt.Print("choose among the following operations: add, sub, div, mul : ")
		fmt.Scan(&operation)
		var result int

		switch operation {
		case "add" :
			result = operations.Add(numberOne, numberTwo)
		
		case "sub" :
			result = operations.Sub(numberOne, numberTwo)
		
		case "mul" :
			result = operations.Mul(numberOne, numberTwo)
		
		case "div" :
			result = operations.Div(numberOne, numberTwo)
		default :
			result = 0
		}
		fmt.Println("The result is :",result)
}