package array_easy

func IsOriginalSorted(arr []int) bool {
	n := len(arr)

	// empty array, one element is ascending sorted
	if n == 0 || n == 1 {
		return true
	}

	count := 0
	for i := 1; i <= n-1; i++ {
		// previous element is greater than next, increase count
		if arr[i-1] > arr[i] {
			count++
		}
	}

	if arr[n-1] > arr[0] {
		count++
	}

	// if we find violation atmost 1 or equal to n-1
	// we know it could have been rotated by x where x range from [0,n-1]
	return count <= 1
}
