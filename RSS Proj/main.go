package main

import (
	"fmt"
	"os"
	"log"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load(".env")
	

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not found in the enviornment")
	}
	fmt.Println("Port: ", portString)

}