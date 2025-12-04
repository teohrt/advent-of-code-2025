package solution

import (
	"fmt"
	"log"
	"math"
	"solution/pkg/lineIterator"
	"strconv"
)

func mod(a, b int) int {
	remainder := a % b
	if remainder < 0 {
		remainder += b
	}
	return remainder
}

func Solve(filePath string) (int, int, int) {
	iterator, err := lineIterator.NewLineIterator(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer iterator.Close()

	count := 0
	endsAtZero := 0
	zeros := 0

	dialPosition := 50

	for iterator.Next() {
		count++

		line := iterator.Line()
		direction := line[0]
		distance, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Fatal(err)
		}
		nextDialPosition := dialPosition
		switch direction {
		case 'L':
			nextDialPosition -= distance
		case 'R':
			nextDialPosition += distance
		default:
			log.Fatalf("invalid direction: %s", string(direction))
		}

		nextDialPosition = mod(nextDialPosition, 100)
		// solution 1
		if nextDialPosition == 0 {
			endsAtZero++
		}

		// solution 2
		// Full 360 degree rotations always pass 0
		totalDistance := math.Abs(float64(distance) + float64(dialPosition))
		fullRotations := int(math.Floor(totalDistance / 100))
		fmt.Printf("fullRotations: %d\n", fullRotations)
		zeros += fullRotations

		// If we start at 0, we can only pass 0 by a full rotation, not a partial one
		if dialPosition != 0 {
			wentToZero := nextDialPosition == 0
			rightPastZero := direction == 'R' && nextDialPosition < dialPosition
			leftPastZero := direction == 'L' && nextDialPosition > dialPosition
			if wentToZero || rightPastZero || leftPastZero {
				zeros++
			}
		}

		dialPosition = nextDialPosition
		fmt.Printf("dialPosition: %d\n", dialPosition)
	}

	return count, endsAtZero, zeros
}
