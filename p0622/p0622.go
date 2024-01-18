package p0622

/*
 * @lc app=leetcode.cn id=622 lang=golang
 *
 * [622] 设计循环队列
 */

// @lc code=start
type MyCircularQueue struct {
	arr    []int
	front  int
	rear   int
	length int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		arr:    make([]int, k+1),
		front:  0,
		rear:   0,
		length: k + 1,
	}
}

func (m *MyCircularQueue) EnQueue(value int) bool {
	if m.IsFull() {
		return false
	}
	m.arr[m.rear] = value
	m.rear++
	m.rear = m.rear % m.length
	return true
}

func (m *MyCircularQueue) DeQueue() bool {
	if m.IsEmpty() {
		return false
	}
	m.front++
	m.front = m.front % m.length
	return true
}

func (m *MyCircularQueue) Front() int {
	if m.IsEmpty() {
		return -1
	}
	return m.arr[m.front]
}

func (m *MyCircularQueue) Rear() int {
	if m.IsEmpty() {
		return -1
	}
	return m.arr[(m.rear-1+m.length)%m.length]
}

func (m *MyCircularQueue) IsEmpty() bool {
	return m.front == m.rear
}

func (m *MyCircularQueue) IsFull() bool {
	return (m.rear+1)%m.length == m.front
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
// @lc code=end
