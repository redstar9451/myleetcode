package p138

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

/*
 * @lc app=leetcode.cn id=138 lang=golang
 *
 * [138] 随机链表的复制
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	nodeMap := make(map[*Node]*Node)
	dummyHead := &Node{}
	newCur := dummyHead
	for cur := head; cur != nil; cur = cur.Next {
		if nodeMap[cur] == nil {
			nodeMap[cur] = &Node{cur.Val, nil, nil}
		}
		if cur.Random != nil && nodeMap[cur.Random] == nil {
			nodeMap[cur.Random] = &Node{cur.Random.Val, nil, nil}
		}
		newCur.Next = nodeMap[cur]
		newCur = newCur.Next
		newCur.Random = nodeMap[cur.Random]
	}

	return dummyHead.Next
}

// @lc code=end
