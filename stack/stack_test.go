package stack

import (
	"reflect"
	"testing"
)

func TestAsteroidCollision(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "No collision, all asteroids move forward",
			input:    []int{5, 10, 20},
			expected: []int{5, 10, 20},
		},
		{
			name:     "Simple collision: Smaller asteroid destroyed",
			input:    []int{8, -5},
			expected: []int{8},
		},
		{
			name:     "Multiple asteroids with one explosion",
			input:    []int{10, 2, -5},
			expected: []int{10},
		},
		{
			name:     "Equal asteroids collide and both are destroyed",
			input:    []int{10, -10},
			expected: []int{},
		},
		{
			name:     "Alternating collisions: Some asteroids survive",
			input:    []int{5, 10, -5},
			expected: []int{5, 10},
		},
		{
			name:     "No asteroids in the input",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "Single asteroid, no collision",
			input:    []int{5},
			expected: []int{5},
		},
		{
			name:     "Complex collision scenario with multiple asteroids",
			input:    []int{8, -8, 10, -5, -8},
			expected: []int{10},
		},
		{
			name:     "No collisions in large sequence of asteroids",
			input:    []int{100, 200, 300},
			expected: []int{100, 200, 300},
		},
		{
			name:     "Consecutive collisions with surviving asteroids",
			input:    []int{10, -5, 20, -10},
			expected: []int{10, 20},
		},
		{
			name:     "Asteroids with equal magnitude in opposite directions",
			input:    []int{-5, 5},
			expected: []int{-5, 5},
		},
		{
			name:     "Alternating directions, no asteroid destruction",
			input:    []int{3, -2, 1, -4, 5},
			expected: []int{-4, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := asteroidCollision(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("For input %v, expected %v but got %v", tt.input, tt.expected, result)
			}
		})
	}
}
