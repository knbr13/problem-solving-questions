package binarysearch

import (
	"math/rand/v2"
	"testing"
)

func TestGuessNumber(t *testing.T) {
	tests := []struct {
		name string
		n    int
		p    int
	}{
		{"n = 1", 1, 1},
		{"n = 2", 2, 2},
		{"small range", 10, 4},
		{"medium range", 100, 28},
		{"large range", 100000, 99865},
		{"very large range", 100000000, 876465},
		{"max int", 2147483647, 1139828923},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var pick int
			guess, pick = buildGuess(tt.n)
			if v := guessNumber(tt.n); v != pick {
				t.Errorf("guessNumber(%d), result = %d, expected: %d\n", tt.n, v, pick)
			}
		})
	}
}

func buildGuess(n int) (func(int) int, int) {
	pick := rand.IntN(n) + 1
	return func(i int) int {
		if i == pick {
			return 0
		} else if i > pick {
			return -1
		}
		return 1
	}, pick
}
