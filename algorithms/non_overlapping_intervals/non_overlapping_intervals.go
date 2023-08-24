package non_overlapping_intervals

import (
	"sort"
)

/*
|-----------|           |          |-----------|
x           y           |          x           y
     |-----------|      |     |-----------|
     a           b      |     a           b

|----------------|
x                y
     |-----|
     a     b
*/

func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) <= 1 {
		return 0
	}

	sort.Slice(intervals, func(a, b int) bool {
		return intervals[a][1] < intervals[b][1]
	})

	current := intervals[0]
	var result int = 1

	for _, interval := range intervals[1:] {
		if current[1] <= interval[0] {
			current = interval
			result++
		}
	}

	return len(intervals) - result
}
