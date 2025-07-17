package array_easy

import (
	"errors"
	"testing"
)

func TestLargestNumber(t *testing.T) {
	type testCase struct {
		name    string
		arr     []int
		n       int
		want    int
		wantErr error
	}

	tests := []testCase{
		{
			name:    "normal case",
			arr:     []int{3, 1, 4, 1, 5, 9},
			n:       6,
			want:    9,
			wantErr: nil,
		},
		{
			name:    "single element",
			arr:     []int{42},
			n:       1,
			want:    42,
			wantErr: nil,
		},
		{
			name:    "all equal elements",
			arr:     []int{5, 5, 5, 5},
			n:       4,
			want:    5,
			wantErr: nil,
		},
		{
			name:    "sorted ascend",
			arr:     []int{1, 2, 3, 4, 5},
			n:       5,
			want:    5,
			wantErr: nil,
		},
		{
			name:    "invalid length zero",
			arr:     []int{1, 2, 3},
			n:       0,
			want:    0,
			wantErr: errors.New("invalid length"),
		},
		{
			name:    "invalid length negative",
			arr:     []int{1, 2, 3},
			n:       -5,
			want:    0,
			wantErr: errors.New("invalid length"),
		},
	}

	for _, tc := range tests {
		tc := tc // capture range variable
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got, err := LargestNumber(tc.arr, tc.n)

			if tc.wantErr != nil {
				if err == nil || err.Error() != tc.wantErr.Error() {
					t.Fatalf("expected error %q, got %v", tc.wantErr.Error(), err)
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if got != tc.want {
				t.Errorf("got %d, want %d", got, tc.want)
			}
		})
	}
}
