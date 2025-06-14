package hashtable

import (
	"reflect"
	"sort"
)

// findDifference solves the problem described at: https://leetcode.com/problems/find-the-difference-of-two-arrays/
func findDifference(nums1 []int, nums2 []int) [][]int {
	set1 := map[int]struct{}{}
	set2 := map[int]struct{}{}

	for _, v := range nums1 {
		set1[v] = struct{}{}
	}
	for _, v := range nums2 {
		set2[v] = struct{}{}
	}

	res1 := []int{}
	res2 := []int{}

	for k := range set1 {
		if _, ok := set2[k]; !ok {
			res1 = append(res1, k)
		}
	}

	for k := range set2 {
		if _, ok := set1[k]; !ok {
			res2 = append(res2, k)
		}
	}

	return [][]int{res1, res2}
}

// uniqueOccurrences solves the problem described at: https://leetcode.com/problems/unique-number-of-occurrences/
func uniqueOccurrences(arr []int) bool {
	countMap := map[int]int{}

	for _, v := range arr {
		countMap[v]++
	}

	seenCounts := map[int]struct{}{}
	for _, v := range countMap {
		if _, ok := seenCounts[v]; ok {
			return false
		}
		seenCounts[v] = struct{}{}
	}

	return true
}

// closeStrings solves the problem described at: https://leetcode.com/problems/determine-if-two-strings-are-close/
func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}
	m1 := map[rune]int{}
	for _, v := range word1 {
		m1[v]++
	}
	m2 := map[rune]int{}
	for _, v := range word2 {
		m2[v]++
	}
	if reflect.DeepEqual(m1, m2) {
		return true
	}
	if len(m1) != len(m2) {
		return false
	}

	arr1 := make([]int, 0, len(m1))
	for _, v := range m1 {
		arr1 = append(arr1, v)
	}
	arr2 := make([]int, 0, len(m2))
	for _, v := range m2 {
		arr2 = append(arr2, v)
	}

	sort.Ints(arr1)
	sort.Ints(arr2)

	for i := range arr1 {
		if arr1[i] != arr2[i] {
			return false
		}
	}

	arr3 := make([]int, 0, len(m1))
	for k := range m1 {
		arr3 = append(arr3, int(k))
	}
	arr4 := make([]int, 0, len(m2))
	for k := range m2 {
		arr4 = append(arr4, int(k))
	}

	sort.Ints(arr3)
	sort.Ints(arr4)

	return reflect.DeepEqual(arr3, arr4)
}
