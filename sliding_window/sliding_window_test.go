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
