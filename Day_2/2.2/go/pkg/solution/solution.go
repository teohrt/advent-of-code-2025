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

	result := 0

	for iterator.Next() {
		line := iterator.Line()
		firstId, lastId := getIds(line)
		for id := firstId; id <= lastId; id++ {
			s := strconv.Itoa(id)
			// check if the id is made of a sequence of digits repeated twice
			// this is only possible if the length of the id is even
			if len(s)%2 == 0 {
				firstHalf := s[:len(s)/2]
				secondHalf := s[len(s)/2:]
				if firstHalf == secondHalf {
					result += id
				}
			}
		}
	}

	return result
}

func getIds(line string) (int, int) {
	split := strings.Split(line, "-")
	if len(split) != 2 {
		log.Fatalf("invalid line: %s", line)
	}
	firstId, err := strconv.Atoi(split[0])
	if err != nil {
		log.Fatal(err)
	}
	lastId, err := strconv.Atoi(split[1])
	if err != nil {
		log.Fatal(err)
	}
	return firstId, lastId
}
