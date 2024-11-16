package p0027

/*
 * @lc app=leetcode.cn id=27 lang=golang
 *
 * [27] 移除元素
 */

// @lc code=start
func removeElement(nums []int, val int) int {
	k := 0
	var tmp int
	val_index := make([]int, 0, len(nums))
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			k++
			val_index = append(val_index, i)
		} else {
			if len(val_index) != 0 {
				tmp = nums[i]
				nums[i] = nums[val_index[0]]
				nums[val_index[0]] = tmp
				val_index = val_index[1:]
				val_index = append(val_index, i)
			}
		}
	}
	return len(nums) - k
}

// @lc code=end
