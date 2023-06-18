package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	var name string
	var address string

	fmt.Println("Enter a name:")
	fmt.Scanln(&name)

	fmt.Println("Enter an address:")
	fmt.Scanln(&address)

	data := map[string]string{
		"name":    name,
		"address": address,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatal("Failed to marshal data to JSON:", err)
	}

	fmt.Println("JSON object:", string(jsonData))
}
