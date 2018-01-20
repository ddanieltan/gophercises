package main

import (
	"encoding/csv"
	"fmt"
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

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println(record)
	}

}
