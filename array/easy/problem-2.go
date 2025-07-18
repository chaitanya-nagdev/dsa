package array_easy

import "math"

// SecondLargestElement returns the second largest unique integer from the given slice.
// If the slice contains fewer than two unique values, it returns -1.
func SecondLargestElement(arr []int) int {
	first := math.MinInt
	second := math.MinInt

	if len(arr) < 2 {
		return -1
	}

	for _, val := range arr {
		if val > first {
			second = first
			first = val
		} else if val < first && val > second {
			second = val
		}
	}

	if second == math.MinInt {
		return -1
	}

	return second
}
