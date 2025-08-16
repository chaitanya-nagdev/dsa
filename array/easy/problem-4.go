package array_easy

func RemoveDuplicates(arr []int) int {
	m := make(map[int]bool)
	i := 0
	j := 0

	for j < len(arr) {
		_, isPresent := m[arr[j]]

		if isPresent {
			j++
		} else {
			m[arr[j]] = true
			arr[i], arr[j] = arr[j], arr[i]
			i++
			j++
		}
	}

	return len(m)
}

// We compare from poisition one
// We compare curr & prev element, if it doesn't match
// we know it is unique we increase j pointer where next unique element
// should go & i is increase, if it matches we just increase i , so util we find
// next unique element we do this & once unique element found we replace
func RemoveDuplicatesOptimal(nums []int) int {
	// arr:  []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
	j := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[j] = nums[i]
			j++
		}
	}
	return j
}
