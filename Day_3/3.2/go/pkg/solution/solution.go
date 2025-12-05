package solution

import (
	"fmt"
	"log"
	"solution/pkg/lineIterator"
	"strconv"
)

func Solve(filePath string) int {
	iterator, err := lineIterator.NewLineIterator(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer iterator.Close()

	result := 0
	const SIZE_OF_NUMBER = 12

	for iterator.Next() {
		lastIndex := -1
		indices := []int{}
		line := iterator.Line()
		for i := SIZE_OF_NUMBER - 1; i >= 0; i-- {
			idx := getLargestValueIndex(line, lastIndex, len(line)-i)
			indices = append(indices, idx)
			lastIndex = idx
		}
		number := getNumberFromIndices(line, indices)
		fmt.Println(number)
		result += number
	}

	return result
}

func getLargestValueIndex(line string, startIndex int, endIndex int) int {
	largest := 0
	largestIdx := 0
	for i := startIndex + 1; i < endIndex; i++ {
		digit := line[i]
		digitInt, err := strconv.Atoi(string(digit))
		if err != nil {
			log.Fatal(err)
		}
		if digitInt > largest {
			largest = digitInt
			largestIdx = i
		}
	}
	return largestIdx
}

func getNumberFromIndices(line string, indices []int) int {
	number := ""
	for _, idx := range indices {
		number += string(line[idx])
	}
	numberInt, err := strconv.Atoi(number)
	if err != nil {
		log.Fatal(err)
	}
	return numberInt
}
