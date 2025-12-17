package solution

import (
	"fmt"
)

func Solve(filePath string) int {
	machines := parseInput(filePath)

	for _, machine := range machines {
		fmt.Printf("machine: %+v\n", machine)
	}

	return 0
}
