package solution

import (
	"log"
	"solution/pkg/lineIterator"
	"sort"
	"strconv"
	"strings"
)

func Solve(filePath string) int {
	iterator, err := lineIterator.NewLineIterator(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer iterator.Close()

	initialIntervals := [][]int{}
	// parse the initial intervals
	for iterator.Next() {
		line := iterator.Line()
		if line == "" {
			break // end of intervals, continuuing would enter into the available ingredient id portion of the input
		}
		intervals := strings.Split(line, "-")
		if len(intervals) != 2 {
			log.Fatalf("invalid line: %s", line)
		}
		start, err := strconv.Atoi(intervals[0])
		if err != nil {
			log.Fatal(err)
		}
		end, err := strconv.Atoi(intervals[1])
		if err != nil {
			log.Fatal(err)
		}
		initialIntervals = append(initialIntervals, []int{start, end})
	}

	// sort and merge the intervals
	mergedIntervals := mergeOverlappingIntervals(initialIntervals)

	result := 0
	// parse the ids and binary search to see if they are in any of the merged intervals
	for iterator.Next() {
		line := iterator.Line()
		id, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		if isInInterval(mergedIntervals, id) {
			result++
		}
	}

	return result
}

func mergeOverlappingIntervals(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	mergedIntervals := [][]int{}
	for _, interval := range intervals {
		if len(mergedIntervals) == 0 || mergedIntervals[len(mergedIntervals)-1][1] < interval[0] {
			mergedIntervals = append(mergedIntervals, interval)
		} else {
			mergedIntervals[len(mergedIntervals)-1][1] = max(mergedIntervals[len(mergedIntervals)-1][1], interval[1])
		}
	}
	return mergedIntervals
}

func isInInterval(intervals [][]int, id int) bool {
	// binary search for the interval that contains the id
	low := 0
	high := len(intervals) - 1
	for low <= high {
		mid := (low + high) / 2
		if intervals[mid][0] <= id && intervals[mid][1] >= id {
			return true
		}
		if intervals[mid][0] > id {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return false
}
