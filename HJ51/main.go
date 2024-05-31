package main

import "fmt"

type ListNode struct {
	m_nKey  int
	m_pNext *ListNode
}

func AddListNode(val int) *ListNode {
	return &ListNode{
		m_nKey:  val,
		m_pNext: nil,
	}
}

func main() {
	var count int
	fmt.Scan(&count)
	listNode := &ListNode{}
	cur := listNode

	for i := 0; i < count; i++ {
		var node int
		fmt.Scan(&node)
		cur.m_pNext = AddListNode(node)
		cur = cur.m_pNext
	}
	var key int
	fmt.Scan(&key)
	cur = listNode.m_pNext
	//val := 0
	ints := make([]int, 0)
	for cur != nil {
		ints = append(ints, cur.m_nKey)
		cur = cur.m_pNext
	}
	//if key == 1 {
	//	fmt.Println(cur.m_nKey)
	//} else {
	//	for i := 0; i < key; i++ {
	//		if i == key {
	//			val = cur.m_nKey
	//			break
	//		}
	//		cur = cur.m_pNext
	//	}
	//	fmt.Println(val)
	//}
}
