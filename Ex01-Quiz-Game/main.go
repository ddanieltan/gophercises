package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func main() {

	file, err := os.Open("problems.csv")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	r := csv.NewReader(file)

	for {
		record, err := r.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		ques := record[0]
		ans := record[1]
		fmt.Printf("\nQuestion: %v Answer: %v", ques, ans)
	}

}
