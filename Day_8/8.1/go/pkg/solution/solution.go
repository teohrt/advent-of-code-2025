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

const CONNECTION_COUNT = 10

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

	// generate all coordinate combinations
	coordinateCombinations := [][]Coordinate{}
	for i := 0; i < len(coords); i++ {
		for j := i + 1; j < len(coords); j++ {
			if i == j {
				continue
			}
			coordinateCombinations = append(coordinateCombinations, []Coordinate{coords[i], coords[j]})
		}
	}

	// sort the coordinate combinations by distance
	sort.Slice(coordinateCombinations, func(i, j int) bool {
		return getDistance(coordinateCombinations[i][0], coordinateCombinations[i][1]) < getDistance(coordinateCombinations[j][0], coordinateCombinations[j][1])
	})

	// generate circuits using the first CONNECTION_COUNT coordinate combinations
	coordsTocircuits := make(map[Coordinate]*map[Coordinate]struct{})
	sliced := coordinateCombinations[:CONNECTION_COUNT]
	fmt.Printf("sliced: %d\n", len(sliced))
	for _, combination := range coordinateCombinations[:CONNECTION_COUNT] {
		c1 := combination[0]
		c2 := combination[1]
		_, c1Exists := coordsTocircuits[c1]
		_, c2Exists := coordsTocircuits[c2]
		sharedSet := make(map[Coordinate]struct{})
		if !c1Exists && !c2Exists {
			sharedSet[c1] = struct{}{}
			sharedSet[c2] = struct{}{}
			coordsTocircuits[c1] = &sharedSet
			coordsTocircuits[c2] = &sharedSet
		} else if c1Exists && !c2Exists {
			// base the shared set on the prexisting set for c1 and add c2
			sharedSet := coordsTocircuits[c1]
			(*sharedSet)[c2] = struct{}{}
			coordsTocircuits[c2] = sharedSet
		} else if !c1Exists && c2Exists {
			// base the shared set on the prexisting set for c2 and add c1
			sharedSet := coordsTocircuits[c2]
			(*sharedSet)[c1] = struct{}{}
			coordsTocircuits[c1] = sharedSet
		} else {
			// merge the two pre-existing sets
			newSharedSet := make(map[Coordinate]struct{})
			for coord := range *coordsTocircuits[c1] {
				newSharedSet[coord] = struct{}{}
			}
			for coord := range *coordsTocircuits[c2] {
				newSharedSet[coord] = struct{}{}
			}
			*coordsTocircuits[c1] = newSharedSet
			*coordsTocircuits[c2] = newSharedSet
		}
	}

	// initialize circuit list and sort by list length
	seen := make(map[*map[Coordinate]struct{}]struct{})
	for _, circuit := range coordsTocircuits {
		if _, ok := seen[circuit]; !ok {
			seen[circuit] = struct{}{}
		}
	}
	circuits := [][]Coordinate{}
	for circuit := range seen {
		circuitSlice := []Coordinate{}
		for coord := range *circuit {
			circuitSlice = append(circuitSlice, coord)
		}
		circuits = append(circuits, circuitSlice)
	}
	sort.Slice(circuits, func(i, j int) bool {
		return len(circuits[i]) > len(circuits[j])
	})

	fmt.Printf("number of circuits: %d\n", len(circuits))
	fmt.Println("Expecting: 11")

	for _, circuit := range circuits {
		fmt.Println(circuit)
	}

	return 0
}

// d = sqrt((x2-x1)^2 + (y2-y1)^2 + (z2-z1)^2)
func getDistance(coord1 Coordinate, coord2 Coordinate) float64 {
	return math.Sqrt(float64(math.Pow(float64(coord2.x-coord1.x), 2) + math.Pow(float64(coord2.y-coord1.y), 2) + math.Pow(float64(coord2.z-coord1.z), 2)))
}
