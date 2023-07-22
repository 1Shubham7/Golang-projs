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
	fmt.Println((question(a,b)))

}

func question(a,b int) int{
	result := a+b
	return result
}

func answer(a,b,result int)bool{
	if (a+b)==result {
		return true
	} else{
		return false
	}
}