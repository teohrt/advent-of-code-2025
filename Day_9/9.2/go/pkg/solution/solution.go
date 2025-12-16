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

/*
https://www.reddit.com/r/adventofcode/comments/1pichj2/comment/nt5guy3

1. compress 2d points so you can grid represent in array. this is an awesome trick and i am super happy to have in my toolbag now. Search "compress 2d coordinates", YouTube videos helped me understand it. Basically, create a map of each sorted unique value for x values, and one for y sorted unique values. So lowest x=1, second lowest =2, etc. Then create a new list of points using the mapped values for x and y. Then you have a lossless compressed representation, with size of grid < number of points instead of size of grid being highest x/y values. So cool

2. rasterize the polygon (fill in the edges). This is pretty easy cause we know point1 is connected to point2, etc. don't forget last point is connected to first point.

3. use raycast point in polygon to find a single inside point. You need to be careful implementing this as the grid representation has non zero width boundy. Find first ".", if there are an odd number of ".#" or "#." transitions, that point is inside.

4. Flood fill polygon using the found inside point as starting. Dfs works great

5. Calculate areas where the rectangle describe by the two corners have all borders as inside the polygon. Easy cause it's been filled!
*/

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

	// Step 1: Compress 2D coordinates 🤯
	// Get unique x and y values
	xSet := make(map[int]bool)
	ySet := make(map[int]bool)
	for _, coord := range coords {
		xSet[coord.x] = true
		ySet[coord.y] = true
	}

	// Convert to slices and sort
	xs := make([]int, 0, len(xSet))
	for x := range xSet {
		xs = append(xs, x)
	}
	sort.Ints(xs)

	ys := make([]int, 0, len(ySet))
	for y := range ySet {
		ys = append(ys, y)
	}
	sort.Ints(ys)

	// Create maps from original values to compressed indices
	xMap := make(map[int]int)
	for i, x := range xs {
		xMap[x] = i
	}

	yMap := make(map[int]int)
	for i, y := range ys {
		yMap[y] = i
	}

	// Compressed red tiles
	compressedRed := make([]Coordinate, len(coords))
	for i, coord := range coords {
		compressedRed[i] = Coordinate{
			x: xMap[coord.x],
			y: yMap[coord.y],
		}
	}

	// Create grid
	width := len(xs)
	height := len(ys)
	grid := make([][]rune, height)
	for i := range grid {
		grid[i] = make([]rune, width)
		for j := range grid[i] {
			grid[i][j] = '.'
		}
	}

	// Mark red tiles in the grid
	for _, coord := range compressedRed {
		grid[coord.y][coord.x] = '#'
	}

	// Step 2: "Rasterize" the polygon (fill in the edges)
	rasterize(grid, compressedRed)

	// Step 3: Find an inside point using raycast
	insidePoint := getInsidePointFromRaycast(grid, width, height)

	// Step 4: DFS flood fill from the inside point
	floodFill(grid, insidePoint, width, height)

	// Step 5: Find the largest rectangle with red corners where the perimeter is inside polygon
	largestArea := findLargestAreaWithinPolygon(grid, coords, xMap, yMap)

	printGrid(grid)
	return largestArea
}

func findLargestAreaWithinPolygon(grid [][]rune, inputCoords []Coordinate, xMap map[int]int, yMap map[int]int) int {
	largestArea := 0
	for i, inputCoord1 := range inputCoords {
		cx1, cy1 := xMap[inputCoord1.x], yMap[inputCoord1.y]
		for _, inputCoord2 := range inputCoords[i+1:] {
			cx2, cy2 := xMap[inputCoord2.x], yMap[inputCoord2.y]

			// Calculate actual area, not the compressed area
			area := getRectangleArea(inputCoord1, inputCoord2)
			// Skip the rectangle rest if it's not a contender
			if area <= largestArea {
				continue
			}

			// Check if the perimeter is inside the polygon
			minCx := min(cx1, cx2)
			maxCx := max(cx1, cx2)
			minCy := min(cy1, cy2)
			maxCy := max(cy1, cy2)
			enclosed := true

			// Check top and bottom edges
			for cx := minCx; cx <= maxCx; cx++ {
				if grid[minCy][cx] == '.' || grid[maxCy][cx] == '.' {
					enclosed = false
					break
				}
			}

			// Check left and right edges
			if enclosed {
				for cy := minCy; cy <= maxCy; cy++ {
					if grid[cy][minCx] == '.' || grid[cy][maxCx] == '.' {
						enclosed = false
						break
					}
				}
			}

			if enclosed {
				largestArea = area
			}
		}
	}
	return largestArea
}

func rasterize(grid [][]rune, compressedRed []Coordinate) {
	// Rasterizing - vector graphics term that basically means to convert points into lines
	// This only works because the prompt guarentees the input to be a closed polygon
	for i, c1 := range compressedRed {
		c2 := compressedRed[(i+1)%len(compressedRed)]
		// wraps to 0 when i+1 equals the length.
		// This connects each point to the next, and the last point to the first
		// Required because the last point in the list is connected to the first point in the list
		//  - guarenteed by the prompt
		if c1.x == c2.x { // vertical line
			yMin := min(c1.y, c2.y)
			yMax := max(c1.y, c2.y)
			for y := yMin; y <= yMax; y++ {
				grid[y][c1.x] = 'X'
			}
		} else if c1.y == c2.y { // horizontal line
			xMin := min(c1.x, c2.x)
			xMax := max(c1.x, c2.x)
			for x := xMin; x <= xMax; x++ {
				grid[c1.y][x] = 'X'
			}
		}
	}
}

func getInsidePointFromRaycast(grid [][]rune, width int, height int) Coordinate {
	// This is a common computer graphics algorithm called "raycasting point-in-polygon"
	// The name was more intimidating than the implementation.
	// For each coordinate in the graph, count the number of "transitions", i.e. changes in boundaries.
	// If that number is odd, the point is inside the polygon.
	// Requires that the polygon is closed and non-self-intersecting.
	// Since this is guaranteed by the input prompt, we can use this method.
	insidePoint := Coordinate{-1, -1}
	for y := range height {
		for x := range width {
			if grid[y][x] != '.' {
				continue
			}
			// Count transitions from this point to the left
			transitions := 0
			prev := '.'
			for i := x - 1; i >= 0; i-- {
				curr := grid[y][i]
				if curr != prev {
					transitions++
				}
				prev = curr
			}

			// Odd number of transitions means inside
			if transitions%2 == 1 {
				insidePoint = Coordinate{x, y}
				break
			}
		}
		if insidePoint.x != -1 {
			break
		}
	}
	return insidePoint
}

func floodFill(grid [][]rune, insidePoint Coordinate, width int, height int) {
	stack := []Coordinate{insidePoint}
	for len(stack) > 0 {
		var popped Coordinate
		popped, stack = stack[0], stack[1:]
		isInGrid := 0 <= popped.x && popped.x < width && 0 <= popped.y && popped.y < height
		shouldFill := isInGrid && grid[popped.y][popped.x] == '.'
		if shouldFill {
			grid[popped.y][popped.x] = '0'
			up := Coordinate{popped.x, popped.y - 1}
			down := Coordinate{popped.x, popped.y + 1}
			left := Coordinate{popped.x - 1, popped.y}
			right := Coordinate{popped.x + 1, popped.y}
			stack = append(stack, up, down, left, right)
		}
	}
}

// The two coordinates are the opposite corners of the rectangle
func getRectangleArea(coord1 Coordinate, coord2 Coordinate) int {
	width := math.Abs(float64(coord2.x-coord1.x)) + 1
	height := math.Abs(float64(coord2.y-coord1.y)) + 1
	return int(width * height)
}

func printGrid(grid [][]rune) {
	for _, row := range grid {
		for _, cell := range row {
			fmt.Print(string(cell))
		}
		fmt.Println()
	}
}
