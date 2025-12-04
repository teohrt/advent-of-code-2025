package solution

import (
	"log"
	"math"
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
	dialPosition := 50

	for iterator.Next() {
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
		nextDialPosition = nextDialPosition % 100

		// Full 360 degree rotations always pass 0
		fullRotations := math.Floor(float64(distance) / 100)
		result += int(fullRotations)

		// If we start at 0, we can only pass 0 by a full rotation, not a partial one
		if dialPosition != 0 {
			wentToZero := nextDialPosition == 0
			rightPastZero := direction == 'R' && nextDialPosition < dialPosition
			leftPastZero := direction == 'L' && nextDialPosition > dialPosition
			if wentToZero || rightPastZero || leftPastZero {
				result++
			}
		}

		dialPosition = nextDialPosition
	}

	return result
}
