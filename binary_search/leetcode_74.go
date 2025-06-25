package binarysearch

// searchMatrix solves the problem described at: https://leetcode.com/problems/search-a-2d-matrix/
func searchMatrix(matrix [][]int, target int) bool {
	n := len(matrix) * len(matrix[0])
	left, right := 0, n-1
	mid := left + (right-left)/2
	firstIndex := mid / len(matrix[0])
	secondIndex := mid % len(matrix[0])

	for left <= right {
		v := matrix[firstIndex][secondIndex]
		if v == target {
			return true
		} else if v > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
		mid = left + (right-left)/2
		firstIndex = mid / len(matrix[0])
		secondIndex = mid % len(matrix[0])
	}
	return false
}
