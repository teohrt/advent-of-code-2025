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

	result := 0

	changesMade := true
	for changesMade {
		count := 0
		for y := 0; y < len(grid); y++ {
			for x := 0; x < len(grid[y]); x++ {
				if grid[y][x] == "@" && canAccess(grid, x, y) {
					count++
					grid[y][x] = "x"
				}
			}
		}
		changesMade = count > 0
		result += count
	}
	return result
}

func canAccess(grid [][]string, x int, y int) bool {
	X := len(grid[0])
	Y := len(grid)

	totalNeighbors := 0

	// top left
	if x > 0 && y > 0 && grid[y-1][x-1] == "@" {
		totalNeighbors++
	}

	// top middle
	if y > 0 && grid[y-1][x] == "@" {
		totalNeighbors++
	}

	// top right
	if x < X-1 && y > 0 && grid[y-1][x+1] == "@" {
		totalNeighbors++
	}

	// left
	if x > 0 && grid[y][x-1] == "@" {
		totalNeighbors++
	}

	// right
	if x < X-1 && grid[y][x+1] == "@" {
		totalNeighbors++
	}

	// bottom left
	if x > 0 && y < Y-1 && grid[y+1][x-1] == "@" {
		totalNeighbors++
	}

	// bottom middle
	if y < Y-1 && grid[y+1][x] == "@" {
		totalNeighbors++
	}

	// bottom right
	if x < X-1 && y < Y-1 && grid[y+1][x+1] == "@" {
		totalNeighbors++
	}

	return totalNeighbors < 4
}
