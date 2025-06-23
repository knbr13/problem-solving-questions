package stack

// asteroidCollision solves the problem described at: https://leetcode.com/problems/asteroid-collision/
func asteroidCollision(asteroids []int) []int {
	stack := []int{}

	for _, ast := range asteroids {
		collision := false
		for len(stack) > 0 && ast < 0 && stack[len(stack)-1] > 0 {
			top := stack[len(stack)-1]
			if top < -ast {
				stack = stack[:len(stack)-1]
				continue
			} else if top == -ast {
				stack = stack[:len(stack)-1]
			}
			collision = true
			break
		}
		if !collision {
			stack = append(stack, ast)
		}
	}

	return stack
}
