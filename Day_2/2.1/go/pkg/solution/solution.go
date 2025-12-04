package solution

import (
	"log"
	"solution/pkg/lineIterator"
)

func Solve(filePath string) int {
	iterator, err := lineIterator.NewLineIterator(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer iterator.Close()

	count := 0

	for iterator.Next() {
		count++
	}

	return count
}
