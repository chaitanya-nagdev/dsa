package array_easy

import "testing"

func TestIsOriginalSorted(t *testing.T) {
	type TestCase struct {
		name string
		arr  []int
		want bool
	}

	tests := []TestCase{
		{
			name: "empty array",
			arr:  []int{},
			want: true,
		},
		{
			name: "invalid size",
			arr:  []int{1},
			want: true,
		},
		{
			name: "x=3 rotated, normal case",
			arr:  []int{3, 4, 5, 1, 2},
			want: true,
		},
		{
			name: "x=1 rotated, normal case",
			arr:  []int{2, 1},
			want: true,
		},
		{
			name: "rotated, normal case",
			arr:  []int{10, 1, 1, 10},
			want: true,
		},

		{
			name: "Original not sorted",
			arr:  []int{2, 1, 3, 4},
			want: false,
		},
		{
			name: "x=3 rotated, normal case with duplicates",
			arr:  []int{3, 4, 5, 5, 1, 2, 2},
			want: true,
		},

		{
			name: "x=0,No Rotation",
			arr:  []int{1, 2, 3},
			want: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := IsOriginalSorted(tc.arr)
			if got != tc.want {
				t.Errorf("Failed %s : expected : %v, got %v", tc.name, tc.want, got)
			}
		})
	}
}
