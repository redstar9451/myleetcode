package p042

func trap(height []int) int {
	maxLeft := make([]int, len(height))
	for i := 1; i < len(height); i++ {
		maxLeft[i] = max(maxLeft[i-1], height[i-1])
	}

	maxRight := make([]int, len(height))
	for i := len(height) - 2; i >= 0; i-- {
		maxRight[i] = max(maxRight[i+1], height[i+1])
	}

	minHeight := 0
	sum := 0
	for i := 1; i < len(height)-1; i++ {
		minHeight = min(maxLeft[i], maxRight[i])
		if height[i] < minHeight {
			sum += minHeight - height[i]
		}
	}
	return sum
}
