package solution

import (
	"fmt"
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

	coordsTocircuits := make(map[Coordinate]*map[Coordinate]struct{})

	for i := 0; i < len(coords); i++ {
		ci := coords[i]
		closestDistance := math.MaxFloat64
		closestIdx := -1
		for j := 0; j < len(coords); j++ {
			if j == i {
				continue
			}
			cj := coords[j]
			distance := getDistance(ci, cj)
			if math.Abs(distance) < closestDistance {
				closestDistance = math.Abs(distance)
				closestIdx = j
			}
		}
		cj := coords[closestIdx]

		fmt.Printf("ci: %v\n", ci)
		fmt.Printf("cj: %v\n", cj)

		// join the two circuits
		newSharedSet := make(map[Coordinate]struct{})
		if coordsTocircuits[ci] != nil {
			for coord := range *coordsTocircuits[ci] {
				newSharedSet[coord] = struct{}{}
			}
		} else {
			newSharedSet[ci] = struct{}{}
		}
		if coordsTocircuits[cj] != nil {
			for coord := range *coordsTocircuits[cj] {
				newSharedSet[coord] = struct{}{}
			}
		} else {
			newSharedSet[cj] = struct{}{}
		}
		fmt.Printf("newSharedSet: %v\n", newSharedSet)
		coordsTocircuits[ci] = &newSharedSet
		coordsTocircuits[cj] = &newSharedSet
	}

	seen := make(map[*map[Coordinate]struct{}]struct{})
	for _, circuit := range coordsTocircuits {
		if _, ok := seen[circuit]; !ok {
			seen[circuit] = struct{}{}
		}
	}

	var circuits [][]Coordinate
	for circuit := range seen {
		circuitSlice := make([]Coordinate, 0, len(*circuit))
		for coord := range *circuit {
			circuitSlice = append(circuitSlice, coord)
		}
		circuits = append(circuits, circuitSlice)
	}

	sort.Slice(circuits, func(i, j int) bool {
		return len(circuits[i]) > len(circuits[j])
	})
	for _, circuit := range circuits {
		fmt.Println(len(circuit))
	}

	result := 0
	result = len(circuits[0]) * len(circuits[1]) * len(circuits[2])

	return result
}

// d = sqrt((x2-x1)^2 + (y2-y1)^2 + (z2-z1)^2)
func getDistance(coord1 Coordinate, coord2 Coordinate) float64 {
	return math.Sqrt(float64(math.Pow(float64(coord2.x-coord1.x), 2) + math.Pow(float64(coord2.y-coord1.y), 2) + math.Pow(float64(coord2.z-coord1.z), 2)))
}
