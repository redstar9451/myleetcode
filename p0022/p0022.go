package p0022

/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 */
// @lc code=start

func backtracking(ans *[]string, left, right int, btTree string, n int) {
	if left+right == 2*n {
		*ans = append(*ans, btTree)
	}
	if left < n {
		backtracking(ans, left+1, right, btTree+"(", n)
	}
	if right < left {
		backtracking(ans, left, right+1, btTree+")", n)
	}
}

func generateParenthesis(n int) []string {
	ans := make([]string, 0)
	btTree := ""
	backtracking(&ans, 0, 0, btTree, n)
	return ans
}

// @lc code=end
