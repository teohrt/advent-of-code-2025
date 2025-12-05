// Package lineIterator provides a streaming iterator for comma-separated values
// in a single line file. It reads the file character-by-character without loading
// the entire line into memory, making it memory-efficient for very long lines.
//
// The iterator reads only the first line from the file and splits it by commas,
// yielding each comma-separated value one at a time. Only the current value is
// kept in memory, not the entire line.
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
//		value := iterator.Line()
//		// Process value...
//	}
package lineIterator

import (
	"bufio"
	"io"
	"os"
)

type LineIterator struct {
	reader      *bufio.Reader
	file        *os.File
	currentLine string
	err         error
	done        bool
}

func NewLineIterator(path string) (*LineIterator, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	reader := bufio.NewReader(f)

	return &LineIterator{
		reader: reader,
		file:   f,
		done:   false,
	}, nil
}

func (it *LineIterator) Next() bool {
	if it.err != nil || it.done {
		return false
	}

	// Build up the current value character by character until we hit a comma or newline
	var currentValue []byte

	for {
		b, err := it.reader.ReadByte()
		if err == io.EOF {
			// If we have a value, return it (last value in the line)
			if len(currentValue) > 0 {
				it.currentLine = string(currentValue)
				it.done = true
				return true
			}
			it.done = true
			return false
		}
		if err != nil {
			it.err = err
			return false
		}

		// Stop at comma or newline
		if b == ',' {
			// Hit a comma - return the current value
			it.currentLine = string(currentValue)
			return true
		}

		if b == '\n' || b == '\r' {
			// If we hit a newline, we're done with the first line
			// Handle \r\n by peeking ahead and consuming the \n if present
			if b == '\r' {
				nextByte, err := it.reader.Peek(1)
				if err == nil && len(nextByte) > 0 && nextByte[0] == '\n' {
					it.reader.ReadByte() // Consume the \n
				}
			}
			it.done = true
			// If we have a value before the newline, return it
			if len(currentValue) > 0 {
				it.currentLine = string(currentValue)
				return true
			}
			return false
		}

		// Append to current value
		currentValue = append(currentValue, b)
	}
}

func (it *LineIterator) Line() string {
	return it.currentLine
}

func (it *LineIterator) Close() error {
	return it.file.Close()
}
