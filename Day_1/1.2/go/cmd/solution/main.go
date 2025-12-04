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
	endsAtZero, zeros := solution.Solve(inputPath)
	fmt.Printf("endsAtZero - solution 1: %d\nzeros - solution 2: %d\n", endsAtZero, zeros)
}
