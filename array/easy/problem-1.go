package array_easy

import "errors"

// LargestNumber returns the maximum value in arr.
// It takes an integer slice arr and its length n.
// Returns an error if n ≤ 0.
func LargestNumber(arr []int, n int) (int, error) {
	if n <= 0 {
		return 0, errors.New("invalid length")
	}

	maxVal := arr[0]
	for _, val := range arr {
		if val > maxVal {
			maxVal = val
		}
	}
	return maxVal, nil
}
