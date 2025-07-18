package array_easy

import "testing"

func TestSecondLargestNumber(t *testing.T) {
	type TestCase struct {
		name string
		arr  []int
		want int
	}

	tests := []TestCase{
		{
			name: "empty array",
			arr:  []int{},
			want: -1,
		},
		{
			name: "invalid size",
			arr:  []int{1},
			want: -1,
		},
		{
			name: "minimum size",
			arr:  []int{-1, 3},
			want: -1,
		},
		{
			name: "all duplicates values",
			arr:  []int{3, 3, 3, 3, 3},
			want: -1,
		},
		{
			name: "normal values",
			arr:  []int{100, -3, 13, -13, 3, 4},
			want: 13,
		},
		{
			name: "second largest in negative",
			arr:  []int{-1, -19, -10},
			want: -10,
		},
		{
			name: "descending values",
			arr:  []int{100, 90, 88, 77, 76, 50},
			want: 90,
		},
		{
			name: "duplicate scenario",
			arr:  []int{100, 100, 76, 77, 77, 50},
			want: 77,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := SecondLargestElement(tc.arr)
			if got != tc.want {
				t.Errorf("Failed %s : expected : %d, got %d", tc.name, tc.want, got)
			}
		})
	}
}
