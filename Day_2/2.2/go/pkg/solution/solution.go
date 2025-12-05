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
			if isInvalid(s) {
				result += id
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

func isInvalid(id string) bool {
	// check if the id is made of a sequence of digits repeated at least twice
	// try all possible pattern lengths from 1 to half the string length
	for patternLen := 1; patternLen <= len(id)/2; patternLen++ {
		// Check if the string length is divisible by the pattern length
		if len(id)%patternLen != 0 {
			continue
		}

		// Extract the pattern (first patternLen characters)
		pattern := id[:patternLen]

		// Check if the entire string is just this pattern repeated
		repetitions := len(id) / patternLen
		if repetitions < 2 {
			continue
		}

		// Build the expected string by repeating the pattern
		expected := strings.Repeat(pattern, repetitions)

		// If it matches, the ID is invalid
		if id == expected {
			return true
		}
	}

	return false
}
