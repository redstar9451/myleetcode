package p0088

/*
 * @lc app=leetcode.cn id=88 lang=golang
 *
 * [88] 合并两个有序数组
 */

// @lc code=start
func merge(nums1 []int, m int, nums2 []int, n int) {
	m--
	n--
	// index 不能用m+n来表示，m和n各减了1，m+n导致index多减一次1
	index := len(nums1) - 1
	// 先遍历数组2，只要数组2都插入了，那么数组1剩下的元素小于数组2的所有元素，且无需再排序
	for n >= 0 {
		for m >= 0 && nums1[m] > nums2[n] {
			nums1[index] = nums1[m]
			index--
			m--
		}
		nums1[index] = nums2[n]
		index--
		n--
	}
}

// @lc code=end
