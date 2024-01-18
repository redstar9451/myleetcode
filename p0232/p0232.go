package p0232

/*
 * @lc app=leetcode.cn id=232 lang=golang
 *
 * [232] 用栈实现队列
 */

// @lc code=start
type MyQueue struct {
	inStack  []int
	outStack []int
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (m *MyQueue) move() {
	if len(m.outStack) != 0 {
		return
	}
	for i := len(m.inStack) - 1; i >= 0; i-- {
		m.outStack = append(m.outStack, m.inStack[i])
	}
	m.inStack = []int{}
}

func (m *MyQueue) Push(x int) {
	m.inStack = append(m.inStack, x)
}

func (m *MyQueue) Pop() int {
	m.move()
	ret := m.outStack[len(m.outStack)-1]
	m.outStack = m.outStack[:len(m.outStack)-1]
	return ret
}

func (m *MyQueue) Peek() int {
	m.move()
	return m.outStack[len(m.outStack)-1]
}

func (m *MyQueue) Empty() bool {
	return len(m.inStack)+len(m.outStack) == 0
}

// @lc code=end
