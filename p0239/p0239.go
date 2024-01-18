package p0239

func maxSlidingWindow(nums []int, k int) []int {
	deque := make([]int, 0, k)
	for i := 0; i < k; i++ {
		for len(deque) > 0 && nums[deque[len(deque)-1]] <= nums[i] {
			deque = deque[:len(deque)-1]
		}
		deque = append(deque, i)
	}

	res := make([]int, 1, len(nums)-k+1)
	res[0] = nums[deque[0]]
	for j := k; j < len(nums); j++ {
		for len(deque) > 0 && nums[deque[len(deque)-1]] <= nums[j] {
			deque = deque[:len(deque)-1]
		}
		for len(deque) > 0 && deque[0] <= j-k {
			deque = deque[1:]
		}
		deque = append(deque, j)
		res = append(res, nums[deque[0]])
	}
	return res
}
