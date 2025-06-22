package hashtable

import (
	"slices"
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
func closeStrings(word1, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}

	m1 := map[byte]int{}
	m2 := map[byte]int{}

	for i := 0; i < len(word1); i++ {
		m1[word1[i]]++
		m2[word2[i]]++
	}

	if len(m1) != len(m2) {
		return false
	}

	if equalMaps(m1, m2) {
		return true
	}

	for k := range m1 {
		if _, ok := m2[k]; !ok {
			return false
		}
	}

	arr1 := make([]int, 0, len(m1))
	arr2 := make([]int, 0, len(m2))

	for k := range m1 {
		arr1 = append(arr1, m1[k])
		arr2 = append(arr2, m2[k])
	}

	slices.Sort(arr1)
	slices.Sort(arr2)

	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}

	return true
}

// equalMaps is a simple helper function to compare maps
func equalMaps(m1, m2 map[byte]int) bool {
	for k, v := range m1 {
		v2 := m2[k]
		if v != v2 {
			return false
		}
	}

	return true
}
