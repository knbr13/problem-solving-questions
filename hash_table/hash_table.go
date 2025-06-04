package hashtable

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
