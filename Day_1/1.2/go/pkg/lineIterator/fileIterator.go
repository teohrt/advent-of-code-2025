// Package lineIterator provides an iterator for reading lines from a file.
// It reads the file line-by-line, yielding each line one at a time.
// This approach is memory-efficient as it only keeps one line in memory at a time,
// rather than loading the entire file.
//
// Example usage:
//
//	iterator, err := lineIterator.NewLineIterator("input.txt")
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer iterator.Close()
//
//	for iterator.Next() {
//		line := iterator.Line()
//		// Process line...
//	}
package lineIterator

import (
	"bufio"
	"os"
)

type LineIterator struct {
	scanner *bufio.Scanner
	file    *os.File
	line    string
	err     error
}

func NewLineIterator(path string) (*LineIterator, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	sc := bufio.NewScanner(f)

	return &LineIterator{
		scanner: sc,
		file:    f,
	}, nil
}

func (it *LineIterator) Next() bool {
	if it.err != nil {
		return false
	}

	ok := it.scanner.Scan()
	if !ok {
		it.err = it.scanner.Err()
		return false
	}

	it.line = it.scanner.Text()
	return true
}

func (it *LineIterator) Line() string {
	return it.line
}

func (it *LineIterator) Close() error {
	return it.file.Close()
}
