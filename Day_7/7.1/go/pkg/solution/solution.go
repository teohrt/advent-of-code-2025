package solution

import (
	"log"
	"solution/pkg/lineIterator"
	"strings"
)

func Solve(filePath string) int {
	iterator, err := lineIterator.NewLineIterator(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer iterator.Close()

	var grid [][]string
	for iterator.Next() {
		line := iterator.Line()
		grid = append(grid, strings.Split(line, ""))
	}

	sourceIndex := len(grid[0]) / 2

	splits := make(map[int]struct{})
	splits[sourceIndex] = struct{}{}

	splitCount := 0
	currentRow := 2
	for currentRow < len(grid) {
		nextSplits := make(map[int]struct{})
		for idx := range splits {
			value := grid[currentRow][idx]
			if value == "^" {
				nextSplits[idx+1] = struct{}{}
				nextSplits[idx-1] = struct{}{}
				splitCount += 1
			}
		}
		currentRow += 2 // skip the next row
		splits = nextSplits
	}

	return splitCount
}
