package solution

import (
	"fmt"
	"log"
	"solution/pkg/lineIterator"
)

func Solve(filePath string) int {
	iterator, err := lineIterator.NewLineIterator(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer iterator.Close()

	for iterator.Next() {
		line := iterator.Line()
		fmt.Println(line)
	}

	result := 0

	return result
}
