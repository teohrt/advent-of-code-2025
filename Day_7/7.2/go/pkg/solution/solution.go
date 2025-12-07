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
	count := traverse(grid, sourceIndex, 2) + 1
	return count
}

func traverse(grid [][]string, x int, y int) int {
	depth := y
	for depth < len(grid) {
		if grid[depth][x] == "^" {
			left := traverse(grid, x-1, depth)
			right := traverse(grid, x+1, depth)
			return left + right + 1
		}
		depth++
	}
	return 0
}
