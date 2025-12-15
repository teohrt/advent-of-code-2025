package solution

import (
	"log"
	"math"
	"solution/pkg/lineIterator"
	"sort"
	"strconv"
	"strings"
)

type Coordinate struct {
	x int
	y int
}

func Solve(filePath string) int {
	iterator, err := lineIterator.NewLineIterator(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer iterator.Close()

	var coords []Coordinate
	for iterator.Next() {
		line := iterator.Line()
		rawValues := strings.Split(line, ",")
		x, err := strconv.Atoi(rawValues[0])
		if err != nil {
			log.Fatal(err)
		}
		y, err := strconv.Atoi(rawValues[1])
		if err != nil {
			log.Fatal(err)
		}
		coordinate := Coordinate{
			x: x,
			y: y,
		}
		coords = append(coords, coordinate)
	}

	// generate all coordinate combinations
	coordinateCombinations := [][]Coordinate{}
	for i := 0; i < len(coords); i++ {
		for j := i + 1; j < len(coords); j++ {
			coordinateCombinations = append(coordinateCombinations, []Coordinate{coords[i], coords[j]})
		}
	}

	// sort the coordinate combinations by area, largest first
	sort.Slice(coordinateCombinations, func(i, j int) bool {
		return getRectangleArea(coordinateCombinations[i][0], coordinateCombinations[i][1]) > getRectangleArea(coordinateCombinations[j][0], coordinateCombinations[j][1])
	})

	// return the area of the largest rectangle
	largestRectangle := coordinateCombinations[0]
	return getRectangleArea(largestRectangle[0], largestRectangle[1])
}

// The two coordinates are the opposite corners of the rectangle
func getRectangleArea(coord1 Coordinate, coord2 Coordinate) int {
	width := math.Abs(float64(coord2.x-coord1.x)) + 1
	height := math.Abs(float64(coord2.y-coord1.y)) + 1
	return int(width * height)
}
