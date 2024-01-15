package p0739

func dailyTemperatures(temperatures []int) []int {
	stack := []int{}
	res := make([]int, len(temperatures))
	for i := len(temperatures) - 1; i >= 0; i-- {
		if len(stack) == 0 {
			res[i] = 0
			stack = append(stack, i)
			continue
		}
		for len(stack) > 0 && temperatures[i] >= temperatures[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			res[i] = 0
		} else {
			res[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}
	ans := make([]int, len(res))
	for i := 0; i < len(res); i++ {
		if res[i] == 0 {
			ans[i] = 0
		} else {
			ans[i] = res[i] - i
		}
	}
	return ans
}
