package solution

import (
	"fmt"
	"log"
	"solution/pkg/lineIterator"
	"strconv"
	"strings"
)

type Coordinate struct {
	x int
	y int
	z int
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
		z, err := strconv.Atoi(rawValues[2])
		if err != nil {
			log.Fatal(err)
		}
		coordinate := Coordinate{
			x: x,
			y: y,
			z: z,
		}
		coords = append(coords, coordinate)
	}

	result := 0
	for _, coord := range coords {
		fmt.Printf("coord: %+v\n", coord)
	}

	return result
}
