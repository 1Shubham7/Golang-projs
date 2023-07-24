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
	questionAndAnswer,err := pullQuestion(fileName)
	if err != nil{
		exit("Something went wrong, try again later: %w, file: %s", err, fileName)
	}
	correctAnswer := 0
	var answer string
	for i,p := range questionAndAnswer{
		fmt.Println("Q: "+ questionAndAnswer[i].question)
		fmt.Scanln(&answer)
		if answer == questionAndAnswer[i].answer{
			correctAnswer++
		}
	}
}

func pullQuestion(fileName string)([]questionAndAnswer, error){

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

func praseQuestion(allQuestions [][]string) []questionAndAnswer{
	
	r := make ([]questionAndAnswer, len(allQuestions))
	for i := 0; i<len(allQuestions); i++{
		r[i] = questionAndAnswer{question: allQuestions[i][[0], answer: allQuestions[i][1]]}
	}
	return r
}

func exit(msg string){
	fmt.Println(msg)
	os.Exit(1)
}


// what did I learn -
// how to read a csv file using encoding/csv