package main

import (
	"fmt"
	"log"
	"os"
	"solution/pkg/solution"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: solution <input-file-path>")
	}

	inputPath := os.Args[1]
	input, err := os.ReadFile(inputPath)
	if err != nil {
		log.Fatal(err)
	}

	result := solution.Solve(string(input))
	fmt.Println(result)
}
