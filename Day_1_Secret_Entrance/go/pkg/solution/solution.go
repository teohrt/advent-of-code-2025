package solution

import (
	"log"
	"solution/pkg/lineiterator"
	"strconv"
)

func Solve(filePath string) int {
	iterator, err := lineiterator.NewLineIterator(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer iterator.Close()

	result := 0
	dialPosition := 50

	for iterator.Next() {
		line := iterator.Line()
		direction := line[0]
		distance, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Fatal(err)
		}
		switch direction {
		case 'L':
			dialPosition -= distance
		case 'R':
			dialPosition += distance
		default:
			log.Fatalf("invalid direction: %s", string(direction))
		}

		dialPosition = dialPosition % 100
		if dialPosition == 0 {
			result++
		}
	}

	return result
}
