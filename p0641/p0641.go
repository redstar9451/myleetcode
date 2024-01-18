package p0641

/*
 * @lc app=leetcode.cn id=641 lang=golang
 *
 * [641] 设计循环双端队列
 */

// @lc code=start
type MyCircularDeque struct {
	q      []int
	front  int
	rear   int
	length int
}

func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{
		q:      make([]int, k+1),
		front:  0,
		rear:   0,
		length: k + 1,
	}
}

func (m *MyCircularDeque) InsertFront(value int) bool {
	if m.IsFull() {
		return false
	}
	m.front = (m.front - 1 + m.length) % m.length
	m.q[m.front] = value
	return true
}

func (m *MyCircularDeque) InsertLast(value int) bool {
	if m.IsFull() {
		return false
	}
	m.q[m.rear] = value
	m.rear = (m.rear + 1) % m.length
	return true
}

func (m *MyCircularDeque) DeleteFront() bool {
	if m.IsEmpty() {
		return false
	}
	m.front = (m.front + 1 + m.length) % m.length
	return true
}

func (m *MyCircularDeque) DeleteLast() bool {
	if m.IsEmpty() {
		return false
	}
	m.rear = (m.rear - 1 + m.length) % m.length
	return true
}

func (m *MyCircularDeque) GetFront() int {
	if m.IsEmpty() {
		return -1
	}
	return m.q[m.front]
}

func (m *MyCircularDeque) GetRear() int {
	if m.IsEmpty() {
		return -1
	}
	return m.q[(m.rear-1+m.length)%m.length]
}

func (m *MyCircularDeque) IsEmpty() bool {
	return m.front == m.rear
}

func (m *MyCircularDeque) IsFull() bool {
	return (m.rear+1)%m.length == m.front
}

// @lc code=end
