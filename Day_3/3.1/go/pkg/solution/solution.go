package solution

import (
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

	for iterator.Next() {
		largestLeft := 0
		leftIndex := 0
		largestRight := 0
		line := iterator.Line()

		// find largest left digit
		// exclude the last character - the  last character can only be the right digit
		for i := 0; i < len(line)-1; i++ {
			digit := line[i]
			digitInt, err := strconv.Atoi(string(digit))
			if err != nil {
				log.Fatal(err)
			}
			if digitInt > largestLeft {
				largestLeft = digitInt
				leftIndex = i
			}
		}

		// find largest right digit
		for i := leftIndex + 1; i < len(line); i++ {
			digit := line[i]
			digitInt, err := strconv.Atoi(string(digit))
			if err != nil {
				log.Fatal(err)
			}
			if digitInt > largestRight {
				largestRight = digitInt
			}
		}

		// avoids string conversions
		result += (largestLeft * 10) + largestRight
	}

	return result
}
