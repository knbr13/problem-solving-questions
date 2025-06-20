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

// Hi, here's your problem today. This problem was recently asked by Amazon:
// You are given a string s, and an integer k. Return the length of the longest substring in s that contains at most k distinct characters.
// For instance, given the string:
// aabcdefff and k = 3, then the longest substring with 3 distinct characters would be defff. The answer should be 5.
func longestSubstringWithKDistinctCharacter(s string, k int) int {
	if k == 0 || len(s) == 0 {
		return 0
	}

	m := make(map[byte]int)
	ml := 0

	for l, r := 0, 0; r < len(s); r++ {
		m[s[r]]++

		for len(m) > k {
			m[s[l]]--
			if m[s[l]] == 0 {
				delete(m, s[l])
			}
			l++
		}

		ml = max(ml, r-l+1)
	}

	return ml
}
