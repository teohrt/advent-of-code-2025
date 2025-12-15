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

	coordsTocircuits := make(map[Coordinate]*[]Coordinate)

	for i := 0; i < len(coords); i++ {
		ci := coords[i]
		closestDistance := math.MaxFloat64
		closestIdx := -1
		for j := i + 1; j < len(coords); j++ {
			cj := coords[j]
			distance := getDistance(ci, cj)
			if distance < closestDistance {
				closestDistance = distance
				closestIdx = j
			}
		}
		cj := coords[closestIdx]

		// join the two circuits
		var newSharedSlice []Coordinate
		if coordsTocircuits[ci] != nil {
			newSharedSlice = append(newSharedSlice, *coordsTocircuits[ci]...)
		} else {
			newSharedSlice = append(newSharedSlice, ci)
		}
		if coordsTocircuits[cj] != nil {
			newSharedSlice = append(newSharedSlice, *coordsTocircuits[cj]...)
		} else {
			newSharedSlice = append(newSharedSlice, cj)
		}
		coordsTocircuits[ci] = &newSharedSlice
		coordsTocircuits[cj] = &newSharedSlice
	}

	seen := make(map[*[]Coordinate]struct{})
	for _, circuit := range coordsTocircuits {
		if _, ok := seen[circuit]; !ok {
			seen[circuit] = struct{}{}
		}
	}

	var circuits [][]Coordinate
	for circuit := range seen {
		circuits = append(circuits, *circuit)
	}

	sort.Slice(circuits, func(i, j int) bool {
		return len(circuits[i]) > len(circuits[j])
	})

	result := 0
	// result = len(circuits[0]) * len(circuits[1]) * len(circuits[2])

	return result
}

// d = sqrt((x2-x1)^2 + (y2-y1)^2 + (z2-z1)^2)
func getDistance(coord1 Coordinate, coord2 Coordinate) float64 {
	return math.Sqrt(float64(math.Pow(float64(coord2.x-coord1.x), 2) + math.Pow(float64(coord2.y-coord1.y), 2) + math.Pow(float64(coord2.z-coord1.z), 2)))
}
