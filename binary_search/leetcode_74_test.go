package binarysearch

import "testing"

func TestSearchMatrix(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		target int
		want   bool
	}{
		{
			name:   "Target exists",
			matrix: [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
			target: 3,
			want:   true,
		},
		{
			name:   "Target does not exist",
			matrix: [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
			target: 13,
			want:   false,
		},
		{
			name:   "Single element matrix, target exists",
			matrix: [][]int{{5}},
			target: 5,
			want:   true,
		},
		{
			name:   "Single element matrix, target does not exist",
			matrix: [][]int{{5}},
			target: 10,
			want:   false,
		},
		{
			name:   "Matrix with one row, target exists",
			matrix: [][]int{{1, 2, 3, 4, 5}},
			target: 3,
			want:   true,
		},
		{
			name:   "Matrix with one row, target does not exist",
			matrix: [][]int{{1, 2, 3, 4, 5}},
			target: 6,
			want:   false,
		},
		{
			name:   "Matrix with one column, target exists",
			matrix: [][]int{{1}, {3}, {5}, {7}, {9}},
			target: 5,
			want:   true,
		},
		{
			name:   "Matrix with one column, target does not exist",
			matrix: [][]int{{1}, {3}, {5}, {7}, {9}},
			target: 6,
			want:   false,
		},
		{
			name:   "Matrix with negative and large values, target exists",
			matrix: [][]int{{-1000, -500, 0, 500, 1000}, {2000, 2500, 3000, 3500, 4000}},
			target: 0,
			want:   true,
		},
		{
			name:   "Matrix with negative and large values, target does not exist",
			matrix: [][]int{{-1000, -500, 0, 500, 1000}, {2000, 2500, 3000, 3500, 4000}},
			target: 1500,
			want:   false,
		},
		{
			name:   "Matrix with negative values only, target exists",
			matrix: [][]int{{-10, -5, 0}, {-4, 1, 6}},
			target: -5,
			want:   true,
		},
		{
			name:   "Matrix with negative values only, target does not exist",
			matrix: [][]int{{-10, -5, 0}, {-4, 1, 6}},
			target: -7,
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := searchMatrix(tt.matrix, tt.target)
			if got != tt.want {
				t.Errorf("searchMatrix(%v, %d) = %v; want %v", tt.matrix, tt.target, got, tt.want)
			}
		})
	}
}
