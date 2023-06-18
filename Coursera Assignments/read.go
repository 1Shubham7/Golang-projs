package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Name struct {
	Fname string
	Lname string
}

func main() {
	fmt.Print("Enter the name of the text file: ")
	var filename string
	fmt.Scanln(&filename)

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Failed to open the file:", err)
	}
	defer file.Close()

	names := make([]Name, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		if len(fields) != 2 {
			continue // Skip lines that do not have two fields
		}
		name := Name{
			Fname: fields[0],
			Lname: fields[1],
		}
		names = append(names, name)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Failed to read the file:", err)
	}

	for _, name := range names {
		fmt.Println("First Name:", name.Fname)
		fmt.Println("Last Name:", name.Lname)
	}
}
