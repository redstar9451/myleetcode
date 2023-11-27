/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 */

// @lc code=start
package p20

type Stack []rune

func (s *Stack) Push(r rune) {
	*s = append(*s, r)
}

func (s *Stack) Pop() (rune, bool) {
	if len(*s) == 0 {
		return 0, false
	} else {
		res := (*s)[len(*s)-1]
		*s = (*s)[:len(*s)-1]
		return res, true
	}
}

func (s *Stack) Peek() (rune, bool) {
	if len(*s) == 0 {
		return 0, false
	} else {
		res := (*s)[len(*s)-1]
		return res, true
	}
}

func (s *Stack) Len() int {
	return len(*s)
}

func index(s []rune, r rune) int {
	for i, a := range s {
		if a == r {
			return i
		}
	}
	return -1
}

func isValid(s string) bool {
	left := []rune{'(', '{', '['}
	right := []rune{')', '}', ']'}
	stack := new(Stack)
	for _, x := range s {
		if index(left, x) == -1 && index(right, x) == -1 {
			return false
		}
		if index(left, x) != -1 {
			stack.Push(x)
		} else {
			if stack.Len() == 0 {
				return false
			}
			cur, _ := stack.Pop()
			le := index(left, cur)
			ri := index(right, x)
			if le != ri {
				return false
			}
		}
	}
	return stack.Len() == 0
}

// @lc code=end
