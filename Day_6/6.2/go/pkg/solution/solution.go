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
		grid = append(grid, strings.Split(line, ""))
	}
	result := 0
	operatorAndColumnRanges := getOperatorAndColumnRanges(grid)
	for _, operatorAndColumnRange := range operatorAndColumnRanges {
		op, start, end := operatorAndColumnRange.operator, operatorAndColumnRange.startColumn, operatorAndColumnRange.endColumn
		var total int
		if op == "+" {
			total = 0
		} else {
			total = 1
		}
		for i := end; i >= start; i-- {
			number := getNumberFromGrid(grid, i)
			if op == "+" {
				total += number
			} else {
				total *= number
			}
		}
		result += total
	}

	return result
}

type OperatorAndColumnRange struct {
	operator    string
	startColumn int
	endColumn   int
}

func getOperatorAndColumnRanges(grid [][]string) []OperatorAndColumnRange {
	operatorAndColumnRanges := []OperatorAndColumnRange{}
	i := 1
	Y := len(grid)
	X := len(grid[Y-1])
	last := OperatorAndColumnRange{
		operator:    grid[Y-1][0],
		startColumn: 0,
	}
	for i < X {
		value := grid[Y-1][i]
		if value != " " {
			last.endColumn = i - 2 // -2 because we want to exclude the operator and the space after it
			operatorAndColumnRanges = append(operatorAndColumnRanges, last)
			last = OperatorAndColumnRange{
				operator:    value,
				startColumn: i,
			}
		}
		i++
	}
	last.endColumn = i - 1 // no space to exclude since it's the last column
	operatorAndColumnRanges = append(operatorAndColumnRanges, last)
	return operatorAndColumnRanges
}

func getNumberFromGrid(grid [][]string, column int) int {
	result := 0
	multiplier := 1
	for i := len(grid) - 2; i >= 0; i-- {
		number := grid[i][column]
		if number != " " {
			numberInt, err := strconv.Atoi(number)
			if err != nil {
				log.Fatal(err)
			}
			result += numberInt * multiplier
			multiplier *= 10
		}
	}
	return result
}
