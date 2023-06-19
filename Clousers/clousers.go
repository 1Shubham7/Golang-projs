package main

import "fmt"

func concatter() func(string) string {
	doc := ""
	return func(word string) string {
		doc += word + " "
		return doc
	}
}

func main() {
	join := concatter()
	join("Hello")
	join("Ji")
	join("Good")
	join("Morning")
	fmt.Println(join(""))
}