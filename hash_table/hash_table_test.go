package hashtable

import (
	"reflect"
	"sort"
	"testing"
)

func TestFindDifference(t *testing.T) {
	tests := []struct {
		name     string
		nums1    []int
		nums2    []int
		expected [][]int
	}{
		{
			name:     "Typical case with some overlap",
			nums1:    []int{1, 2, 3, 3},
			nums2:    []int{2, 4, 6, 2},
			expected: [][]int{{1, 3}, {4, 6}},
		},
		{
			name:     "No common elements",
			nums1:    []int{5, 7, 9},
			nums2:    []int{1, 2, 3},
			expected: [][]int{{5, 7, 9}, {1, 2, 3}},
		},
		{
			name:     "One side empty (nums1 empty)",
			nums1:    []int{},
			nums2:    []int{10, 20, 20},
			expected: [][]int{{}, {10, 20}},
		},
		{
			name:     "One side empty (nums2 empty)",
			nums1:    []int{8, 8, 8},
			nums2:    []int{},
			expected: [][]int{{8}, {}},
		},
		{
			name:     "Both empty",
			nums1:    []int{},
			nums2:    []int{},
			expected: [][]int{{}, {}},
		},
		{
			name:     "Identical sets with duplicates",
			nums1:    []int{1, 1, 2, 3, 3},
			nums2:    []int{1, 2, 2, 3},
			expected: [][]int{{}, {}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := findDifference(tt.nums1, tt.nums2)

			sort.Ints(actual[0])
			sort.Ints(actual[1])
			sort.Ints(tt.expected[0])
			sort.Ints(tt.expected[1])

			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("findDifference(%v, %v) = %v, want %v",
					tt.nums1, tt.nums2, actual, tt.expected)
			}
		})
	}
}

func TestUniqueOccurrences(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		expected bool
	}{
		{
			name:     "Typical case with unique occurrence counts",
			arr:      []int{1, 2, 2, 1, 1, 3},
			expected: true,
		},
		{
			name:     "Duplicate occurrence counts",
			arr:      []int{1, 2, 2, 1},
			expected: false,
		},
		{
			name:     "Single element array",
			arr:      []int{42},
			expected: true,
		},
		{
			name:     "Empty array",
			arr:      []int{},
			expected: true,
		},
		{
			name:     "Negative numbers with duplicate counts",
			arr:      []int{-1, -1, 2, 2, 3},
			expected: false,
		},
		{
			name:     "Negative numbers with unique counts",
			arr:      []int{-5, -5, -5, 0, 0, 7},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := uniqueOccurrences(tt.arr)
			if actual != tt.expected {
				t.Errorf("uniqueOccurrences(%v) = %v, want %v", tt.arr, actual, tt.expected)
			}
		})
	}
}

func TestCloseStrings(t *testing.T) {
	tests := []struct {
		name   string
		word1  string
		word2  string
		expect bool
	}{
		{
			name:   "identical single char",
			word1:  "a",
			word2:  "a",
			expect: true,
		},
		{
			name:   "simple anagram",
			word1:  "abc",
			word2:  "bca",
			expect: true,
		},
		{
			name:   "different length",
			word1:  "a",
			word2:  "aa",
			expect: false,
		},
		{
			name:   "same chars different freqs but transformable",
			word1:  "aabbcc",
			word2:  "bbccaa",
			expect: true,
		},
		{
			name:   "same freq multichar groups",
			word1:  "abbccc",
			word2:  "cccbba",
			expect: true,
		},
		{
			name:   "mismatched char sets",
			word1:  "aabb",
			word2:  "aabc",
			expect: false,
		},
		{
			name:   "disjoint char sets",
			word1:  "abc",
			word2:  "def",
			expect: false,
		},
		{
			name:   "complex transform",
			word1:  "cabbba",
			word2:  "abbccc",
			expect: true,
		},
		{
			name:   "single char to multichar not transformable",
			word1:  "a",
			word2:  "b",
			expect: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := closeStrings(tc.word1, tc.word2)
			if got != tc.expect {
				t.Errorf("closeStrings(%q, %q) = %v; want %v", tc.word1, tc.word2, got, tc.expect)
			}
		})
	}
}
