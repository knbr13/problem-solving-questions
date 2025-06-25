package stack

import "testing"

func TestIsValid(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{"empty string", "", true},
		{"single pair parentheses", "()", true},
		{"multiple pairs mixed", "()[]{}", true},
		{"mismatched brackets", "(]", false},
		{"nested valid", "([])", true},
		{"only opening brackets", "(((", false},
		{"only closing brackets", ")))", false},
		{"interleaved mismatched", "([)]", false},
		{"long valid string", "({[]}){[()]()}", true},
		{"long invalid string", "({[)]}", false},
		{"single opening bracket", "(", false},
		{"single closing bracket", ")", false},
		{"long repeated valid", "()" + "[]" + "{}" + "()" + "[]" + "{}", true},
		{"long repeated invalid", "([)]" + "([)]", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isValid(tt.s)
			if got != tt.want {
				t.Errorf("isValid(%q) = %v; want %v", tt.s, got, tt.want)
			}
		})
	}
}
