package array_easy

import (
	"reflect"
	"testing"
)

type TestCase struct {
	name string
	arr  []int
	want []int
}

func TestRemoveDuplicate(t *testing.T) {
	tests := []TestCase{
		{
			name: "1 element",
			arr:  []int{1},
			want: []int{1},
		},
		{
			name: "No Duplicate",
			arr:  []int{1, 2, 3, 4, 5},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "all Duplicate",
			arr:  []int{1, 1, 1, 1, 1},
			want: []int{1},
		},
		{
			name: "mixed duplicates",
			arr:  []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			want: []int{0, 1, 2, 3, 4},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := RemoveDuplicates(tc.arr)
			result := tc.arr[:got]
			if !reflect.DeepEqual(result, tc.want) {
				t.Errorf("Failed %s : expected : %v, got %v", tc.name, tc.want, got)
			}
		})
	}
}

func TestRemoveDuplicatesOptimal(t *testing.T) {
	tests := []TestCase{
		{
			name: "1 element",
			arr:  []int{1},
			want: []int{1},
		},
		{
			name: "No Duplicate",
			arr:  []int{1, 2, 3, 4, 5},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "all Duplicate",
			arr:  []int{1, 1, 1, 1, 1},
			want: []int{1},
		},
		{
			name: "mixed duplicates",
			arr:  []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			want: []int{0, 1, 2, 3, 4},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := RemoveDuplicatesOptimal(tc.arr)
			result := tc.arr[:got]
			if !reflect.DeepEqual(result, tc.want) {
				t.Errorf("Failed %s : expected : %v, got %v", tc.name, tc.want, got)
			}
		})
	}
}
