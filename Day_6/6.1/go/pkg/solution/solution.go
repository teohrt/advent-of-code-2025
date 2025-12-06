package solution

import (
	"log"
	"solution/pkg/lineIterator"
	"strconv"
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
		elements := strings.Fields(line)
		grid = append(grid, elements)
	}

	totals := []int{}

	for x := 0; x < len(grid[0]); x++ {
		operator := grid[len(grid)-1][x] // the operator is the last row of the grid

		var total int
		if operator == "+" {
			total = 0
		} else {
			total = 1
		}

		for y := 0; y < len(grid)-1; y++ {
			number, err := strconv.Atoi(grid[y][x])
			if err != nil {
				log.Fatal(err)
			}
			if operator == "+" {
				total += number
			} else {
				total *= number
			}
		}
		totals = append(totals, total)
	}

	result := 0
	for _, total := range totals {
		result += total
	}
	return result
}
