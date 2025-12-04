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
	count, endsAtZero, zeros := solution.Solve(inputPath)
	fmt.Printf("count of instructions: %d\nendsAtZero - solution 1: %d\nzeros - solution 2: %d\n", count, endsAtZero, zeros)
}
