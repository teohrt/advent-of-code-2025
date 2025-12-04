package solution

import (
	"log"
	"math"
	"solution/pkg/lineIterator"
	"strconv"
)

func mod(a, b int) int {
	// This is a custom modulo function that handles negative numbers differently than the built-in modulo operator.
	// The built-in modulo operator returns a negative remainder if the dividend is negative,
	// while this function returns a positive remainder.
	// This is useful because we need to handle negative numbers when calculating the dial position.
	// For example, if the dial position is -1, we need to wrap it around to the other side of the dial.
	// If we use the built-in modulo operator, we would get -1, which is not what we want.
	// By using this custom modulo function, we can get the correct remainder.
	remainder := a % b
	if remainder < 0 {
		remainder += b
	}
	return remainder
}

func Solve(filePath string) (int, int) {
	iterator, err := lineIterator.NewLineIterator(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer iterator.Close()

	endsAtZero := 0
	zeros := 0

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

		nextDialPosition = mod(nextDialPosition, 100)
		// solution 1
		if nextDialPosition == 0 {
			endsAtZero++
		}

		// solution 2
		// Full 360 degree rotations always pass 0
		fullRotations := int(math.Floor(float64(distance) / 100))
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
	}

	return endsAtZero, zeros
}
