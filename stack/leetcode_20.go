package stack

// isValid solves the problem described at: https://leetcode.com/problems/valid-parentheses/
func isValid(s string) bool {
	arr := []rune{}
	for _, v := range s {
		switch v {
		case '{', '[', '(':
			arr = append(arr, v)
		case '}', ']', ')':
			n := len(arr)
			if n == 0 || !((v == ')' && v-1 == arr[n-1]) || (v != ')' && v-2 == arr[n-1])) {
				return false
			}
			arr = arr[:n-1]
		}
	}
	return len(arr) == 0
}
