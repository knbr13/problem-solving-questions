package slidingwindow

// lengthOfLongestSubstring solves the problem described at: https://leetcode.com/problems/longest-substring-without-repeating-characters/
func lengthOfLongestSubstring(s string) int {
	n := len(s)
	if n < 2 {
		return n
	}

	type entry struct {
		occupied bool
		idx      int
	}

	arr := make([]entry, 95)
	arr[int(rune(s[0])-32)] = entry{occupied: true}
	maxlen := 1
	for i, j := 0, 1; j < len(s); j++ {
		idx := int(rune(s[j]) - 32)
		if arr[idx].occupied && arr[idx].idx >= i {
			i = arr[idx].idx + 1
		} else {
			maxlen = max(maxlen, j-i+1)
		}
		arr[idx] = entry{occupied: true, idx: j}
	}

	return maxlen
}
