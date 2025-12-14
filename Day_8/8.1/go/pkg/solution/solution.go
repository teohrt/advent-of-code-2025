package solution

import (
	"fmt"
	"log"
	"math"
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

	c1 := coords[0]
	c2 := coords[1]
	fmt.Printf("c1: %+v\n", c1)
	fmt.Printf("c2: %+v\n", c2)
	fmt.Printf("distance: %f\n", getDistance(c1, c2))

	return result
}

// d = sqrt((x2-x1)^2 + (y2-y1)^2 + (z2-z1)^2)
func getDistance(coord1 Coordinate, coord2 Coordinate) float64 {
	return math.Sqrt(float64(math.Pow(float64(coord2.x-coord1.x), 2) + math.Pow(float64(coord2.y-coord1.y), 2) + math.Pow(float64(coord2.z-coord1.z), 2)))
}
