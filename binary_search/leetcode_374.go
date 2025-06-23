package binarysearch

var (
	// to be initialized before use
	guess func(int) int
)

// guessNumber solves the problem described at: https://leetcode.com/problems/guess-number-higher-or-lower/
func guessNumber(n int) int {
	left, right := 1, n

	for left <= right {
		mid := left + (right-left)/2
		res := guess(mid)

		switch res {
		case 0:
			return mid
		case -1:
			right = mid - 1
		case 1:
			left = mid + 1
		}
	}
	return -1
}
