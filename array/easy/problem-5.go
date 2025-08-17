package array_easy

func RotateArray(nums []int, k int) {
	if k <= 0 || len(nums) <= 1 || k == len(nums) {
		return
	}

	k = k % len(nums)
	ReverseArray(nums, 0, len(nums)-1)
	ReverseArray(nums, k, len(nums)-1)
	ReverseArray(nums, 0, k-1)
}

func ReverseArray(nums []int, x int, y int) {
	for x < y {
		nums[x], nums[y] = nums[y], nums[x]
		x++
		y--
	}
}
