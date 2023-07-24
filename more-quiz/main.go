package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	fileName := flag.String("f", "quiz.csv", "path of the file")
	timer := flag.Int("t", 30, "timer for the quiz")
	flag.Parse()
	problems,err := pullQuestion(fileName)
	if err != nil{
		exit("Something went wrong, try again later: %w, file: %s", err, fileName)
	}
	correctAnswer := 0
}

func pullQuestion(fileName string)([]question, error){

	if fObj, err := os.Open(fileName); err == nil {
		csvR := csv.NewReader(fObj)
		if cLines, err := csvR.ReadAll(); err == nil{
			return parseQuestion(cLines), nil
		} else {
			return nil, fmt.Errorf("Error reading the file %s: %w", fileName, err)
		}
	} else {
		return nil, fmt.Errorf("Error opening file %s", fileName);
	}


}

type problem struct{
	question string
	answer string
}

func praseQuestion(lines [][]string) []question{
	
	return []
}


// what did I learn -
// how to read a csv file using encoding/csv