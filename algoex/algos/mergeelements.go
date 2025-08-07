package algos

import (
	"fmt"
	"sort"
)

func mergeIntervals(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	merged := [][]int{intervals[0]}

	for i := 1; i < len(intervals); i++ {
		// Get the last merged interval.
		lastMerged := merged[len(merged)-1]
		current := intervals[i]

		// Check for an overlap. An overlap occurs if the current interval's start
		// time is less than or equal to the last merged interval's end time.
		if current[0] <= lastMerged[1] {
			// If there's an overlap, merge them by updating the end time of the
			// last merged interval to be the maximum of the two end times.
			if current[1] > lastMerged[1] {
				merged[len(merged)-1][1] = current[1]
			}
		} else {
			// If there's no overlap, the current interval is a new, non-overlapping
			// interval. Add it to the merged list.
			merged = append(merged, current)
		}
	}

	return merged
}

func RunMergeElements() {
	// The problem's example input.
	input := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 17}}
	fmt.Printf("Input: %v\n", input)

	output := mergeIntervals(input)
	fmt.Printf("Output: %v\n", output)

	fmt.Println("---------------------------------")

	// A more complex example with multiple overlaps.
	input2 := [][]int{{1, 4}, {4, 5}, {6, 8}, {7, 9}}
	fmt.Printf("Input: %v\n", input2)
	output2 := mergeIntervals(input2)
	fmt.Printf("Output: %v\n", output2)
}
