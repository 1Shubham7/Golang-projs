package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	
	rand.Seed(time.Now().UnixNano())
	a := rand.Intn(1000)
	b := rand.Intn(1000)
	fmt.Printf("Enter the result of = %v + %v\n",a,b )
	var myAnswer int 
	fmt.Scanln(&myAnswer)
	result := question(a,b)
	fmt.Println((answer(myAnswer,result)))

}

func question(a,b int) int{
	result := a+b
	return result
}

func answer(result, myAnswer int) bool{
	if myAnswer==result {
		return true
	} else{
		return false
	}
}