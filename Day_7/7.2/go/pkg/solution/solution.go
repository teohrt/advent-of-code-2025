package solution

import (
	"fmt"
	"log"
	"solution/pkg/lineIterator"
	"strings"
)

type Node struct {
	visited    bool
	left       *Node
	right      *Node
	coordinate Coordinate
	paths      *int
}

type Coordinate struct {
	y int
	x int
}

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

	head := Node{
		visited: false,
		left:    nil,
		right:   nil,
		coordinate: Coordinate{
			x: sourceIndex,
			y: 2,
		},
		paths: nil,
	}
	nodes := make(map[Coordinate]*Node)
	nodes[head.coordinate] = &head
	createChildren(&head, grid, nodes)
	traverse(&head)
	return *head.paths
}

func traverse(node *Node) int {
	fmt.Printf("node: %+v\n", node.coordinate)
	isLeafNode := node.left == nil && node.right == nil
	if isLeafNode {
		paths := 2
		node.paths = &paths
		return 2
	}
	leftCount := 1
	if node.left != nil {
		if node.left.paths != nil {
			leftCount = *node.left.paths
		} else {
			leftCount = traverse(node.left)
		}
	}

	rightCount := 1
	if node.right != nil {
		if node.right.paths != nil {
			rightCount = *node.right.paths
		} else {
			rightCount = traverse(node.right)
		}
	}

	totalPaths := leftCount + rightCount
	node.paths = &totalPaths
	return totalPaths
}

func createChildren(node *Node, grid [][]string, nodes map[Coordinate]*Node) {
	currentRow := node.coordinate.y + 2
	for currentRow < len(grid) {
		leftCoordinate := Coordinate{x: node.coordinate.x - 1, y: currentRow}
		if grid[leftCoordinate.y][leftCoordinate.x] == "^" {
			if _, exists := nodes[leftCoordinate]; exists {
				node.left = nodes[leftCoordinate]
				break
			}
			// fmt.Printf("node: %+v\n", leftCoordinate)
			leftNode := Node{
				visited:    false,
				left:       nil,
				right:      nil,
				coordinate: leftCoordinate,
				paths:      nil,
			}
			nodes[leftCoordinate] = &leftNode
			node.left = &leftNode
			break
		}
		currentRow += 2
	}

	currentRow = node.coordinate.y + 2
	for currentRow < len(grid) {
		rightCoordinate := Coordinate{x: node.coordinate.x + 1, y: currentRow}
		if grid[rightCoordinate.y][rightCoordinate.x] == "^" {
			if _, exists := nodes[rightCoordinate]; exists {
				node.right = nodes[rightCoordinate]
				break
			}
			// fmt.Printf("node: %+v\n", rightCoordinate)
			rightNode := Node{
				visited:    false,
				left:       nil,
				right:      nil,
				coordinate: rightCoordinate,
				paths:      nil,
			}
			nodes[rightCoordinate] = &rightNode
			node.right = &rightNode
			break
		}
		currentRow += 2
	}

	if node.left != nil {
		createChildren(node.left, grid, nodes)
	}
	if node.right != nil {
		createChildren(node.right, grid, nodes)
	}
}
