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
