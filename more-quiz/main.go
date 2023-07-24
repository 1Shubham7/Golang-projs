package main

import(
	"fmt"
	"os"
)

func main(){
	file, err := os.Open("quiz.csv")
	if err != nil {
		fmt.Println("Something went wrong, please try again later :", err)
		return
	}
	const bufferSize = 1024

	buffer := make([]byte, bufferSize)
	for {
		bytesRead, err:= file.Read(buffer)
		fmt.Println(string(buffer[:bytesRead]))
		if err != nil{
			fmt.Println("Error reading file", err)
			break
		}
		if bytesRead == 0{
			break
		}
	}

func pullQuestion(fileName string)([]question, error){

}



	defer file.Close()

}

type problem struct{
	question string
	answer string
}

func praseProblem