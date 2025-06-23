package binarysearch

import (
	"testing"
)

func TestSearchInsert(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected int
	}{
		{
			name:     "target present in the middle",
			nums:     []int{1, 3, 5, 6},
			target:   5,
			expected: 2,
		},
		{
			name:     "target smaller than all elements",
			nums:     []int{1, 3, 5, 6},
			target:   0,
			expected: 0,
		},
		{
			name:     "target greater than all elements",
			nums:     []int{1, 3, 5, 6},
			target:   7,
			expected: 4,
		},
		{
			name:     "target should be inserted at index 1",
			nums:     []int{1, 3, 5, 6},
			target:   2,
			expected: 1,
		},
		{
			name:     "target at the beginning",
			nums:     []int{1, 3, 5, 6},
			target:   1,
			expected: 0,
		},
		{
			name:     "target at the end",
			nums:     []int{1, 3, 5, 6},
			target:   6,
			expected: 3,
		},
		{
			name:     "single element, target matches",
			nums:     []int{3},
			target:   3,
			expected: 0,
		},
		{
			name:     "single element, target smaller",
			nums:     []int{3},
			target:   2,
			expected: 0,
		},
		{
			name:     "single element, target greater",
			nums:     []int{3},
			target:   4,
			expected: 1,
		},
		{
			name:     "empty array",
			nums:     []int{},
			target:   10,
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := searchInsert(tt.nums, tt.target)
			if result != tt.expected {
				t.Errorf("searchInsert(%v, %d) = %d; want %d",
					tt.nums, tt.target, result, tt.expected)
			}
		})
	}
}
