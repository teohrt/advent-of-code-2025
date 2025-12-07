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

	splits := make(map[int]struct{}) // need to use a set to avoid duplicates
	splits[sourceIndex] = struct{}{}

	splitCount := 0
	currentRow := 2
	for currentRow < len(grid) {
		for idx := range splits {
			value := grid[currentRow][idx]
			if value == "^" {
				delete(splits, idx)
				splits[idx+1] = struct{}{}
				splits[idx-1] = struct{}{}
				splitCount += 1
			}
		}
		currentRow += 2 // skip the next row
	}

	return splitCount
}
