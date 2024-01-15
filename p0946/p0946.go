package p0946

/* func validateStackSequences(pushed []int, popped []int) bool {
	stack := []int{}
	j := 0
	for _, v := range pushed {
		stack = append(stack, v)
		for len(stack) > 0 && stack[len(stack)-1] == popped[j] {
			stack = stack[:len(stack)-1]
			j++
		}
	}
	return len(stack) == 0
} */

func validateStackSequences(pushed []int, popped []int) bool {
	if len(pushed) != len(popped) {
		return false
	}

	stack := make([]int, len(pushed), len(pushed))
	i, j := 0, 0
	for _, v := range pushed {
		stack[i] = v
		for i >= 0 && stack[i] == popped[j] {
			i--
			j++
		}
		i++
	}
	return i == 0
}
