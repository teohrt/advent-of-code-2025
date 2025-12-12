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

	splits := make(map[int]int, 0)
	sourceIndex := len(grid[0]) / 2
	splits[sourceIndex] = 1

	for _, row := range grid[1:] {
		for k := range splits {
			if row[k] == "^" {
				splits[k-1] += splits[k]
				splits[k+1] += splits[k]
				delete(splits, k)
			}
		}
	}

	sum := 0
	for _, b := range splits {
		sum += b
	}
	return sum
}
