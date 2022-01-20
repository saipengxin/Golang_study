package main

import "fmt"

// 链表翻转

type ListNode struct {
	Val  int
	Next *ListNode
}

// 翻转链表
func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	return pre
}

func main() {
	var l = &ListNode{
		1, &ListNode{
			2, &ListNode{
				3, &ListNode{
					4, &ListNode{
						5, nil,
					},
				},
			},
		},
	}
	newL := reverseList(l)
	fmt.Println(newL)
}
