package array_easy

import (
	"reflect"
	"testing"
)

func TestRotateArray(t *testing.T) {
	tests := []struct {
		input    []int
		k        int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7}, 3, []int{5, 6, 7, 1, 2, 3, 4}},
		{[]int{1, 2, 3, 4, 5}, 2, []int{4, 5, 1, 2, 3}},
		{[]int{1, 2, 3}, 4, []int{3, 1, 2}},       // k > len(nums)
		{[]int{1}, 10, []int{1}},                  // single element
		{[]int{}, 5, []int{}},                     // empty array
		{[]int{1, 2}, 2, []int{1, 2}},             // k == len(nums)
		{[]int{1, 2, 3, 4}, 0, []int{1, 2, 3, 4}}, // k == 0
	}

	for _, tt := range tests {
		nums := make([]int, len(tt.input))
		copy(nums, tt.input)

		RotateArray(nums, tt.k)

		if !reflect.DeepEqual(nums, tt.expected) {
			t.Errorf("RotateArray(%v, %d) = %v; expected %v",
				tt.input, tt.k, nums, tt.expected)
		}
	}
}

