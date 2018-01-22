package main

import (
	"encoding/csv"
	"fmt"
	// "io"
	"os"
)

type problem struct {
	ques string
	ans  string
}

func main() {

	file, err := os.Open("problems.csv")
	if err != nil {
		fmt.Println("Failed to open csv")
		os.Exit(1)
	}
	defer file.Close()

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		fmt.Println("Failed to open csv reader")
		os.Exit(1)
	}

	problems := make([]problem, len(lines))
	for i, line := range lines {
		problems[i] = problem{
			ques: line[0],
			ans:  line[1],
		}
	}

	var score int //default value is 0
	for i, prob := range problems {
		fmt.Printf("Question #%s\n%s\n", i, prob.ques)
		var user_ans string
		fmt.Scanf("%s\n", &user_ans)
		if user_ans == prob.ans {
			fmt.Println("Correct!")
			score++
		} else {
			fmt.Println("Incorrect!")
		}

	}
	fmt.Printf("You got %d out of %d questions correct", score, len(problems))

}
