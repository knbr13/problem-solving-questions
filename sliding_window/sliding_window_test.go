package slidingwindow

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{"Empty String", "", 0},
		{"Single Character", "a", 1},
		{"Two Identical Characters", "aa", 1},
		{"String with spaces", "abc def ghi", 7},
		{"String with spaces and duplicates", "abc def abc", 7},
		{"All ASCII characters", "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()", 72},
		{"String with repeated characters", "bbbbbbbbbb", 1},
		{"Substring at the end", "abcabcd", 4},
		{"Substring with duplicates", "abcabcbb", 3},
		{"No Repeated Characters", "abcdef", 6},
		{"String with symbols", "abc!@#def", 9},
		{"Longer Input", "abcdabcdabcdabcd", 4},
		{"Mixed characters", "abc123!@#", 9},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lengthOfLongestSubstring(tt.input)
			if result != tt.expected {
				t.Errorf("For input %q, expected %d, but got %d", tt.input, tt.expected, result)
			}
		})
	}
}

func TestLongestSubstringWithKDistinctCharacter(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		k        int
		expected int
	}{
		{"Example from problem", "aabcdefff", 3, 5},
		{"Empty string", "", 3, 0},
		{"Single character", "a", 1, 1},
		{"String length less than k", "abc", 5, 3},

		{"K=1, single character repeated", "aaaaa", 1, 5},
		{"K=1, multiple characters", "abcde", 1, 1},
		{"K=1, mixed repeated characters", "aabbbccc", 1, 3},

		{"K=2, two characters repeated", "aabb", 2, 4},
		{"K=2, more than two distinct characters", "aabbc", 2, 4},
		{"K=2, alternating characters", "abababa", 2, 7},
		{"K=2, complex pattern", "aabcbaa", 2, 3},

		{"K=3, exactly three distinct characters", "abcabcabc", 3, 9},
		{"K=3, more than three distinct characters", "abcdefg", 3, 3},
		{"K=3, repeated characters", "aaabbbccc", 3, 9},
		{"K=3, complex pattern", "abcbcbade", 3, 7},

		{"K=5, string with exactly 5 distinct chars", "abcdeabcde", 5, 10},
		{"K=10, string with fewer distinct chars", "abcdefgabc", 10, 10},

		{"Distinct characters at the end", "aaaabcde", 3, 6},
		{"Distinct characters at the beginning", "abcdeaaa", 3, 5},
		{"Distinct characters in the middle", "aaabcdeaa", 3, 5},

		{"Complex pattern 1", "eceba", 2, 3},
		{"Complex pattern 2", "ccaabbb", 2, 5},
		{"Complex pattern 3", "aabbcc", 3, 6},
		{"Complex pattern 4", "abaccc", 2, 4},
		{"chatgpt", "abcadcacacaca", 3, 11},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := longestSubstringWithKDistinctCharacter(tt.input, tt.k)
			if result != tt.expected {
				t.Errorf("For input %q with k=%d, expected %d, but got %d", tt.input, tt.k, tt.expected, result)
			}
		})
	}
}
