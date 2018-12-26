package main

import "fmt"

func main() {
	ln1 := &ListNode{Val: 3}
	ln2 := &ListNode{Val: 2}
	ln3 := &ListNode{Val: 0}
	ln4 := &ListNode{Val: -4}
	ln1.Next = ln2
	ln2.Next = ln3
	ln3.Next = ln4
	ln4.Next = ln2

	ll := detectCycle(ln1)
	fmt.Printf("=====> %#v", ll)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil {
		slow = slow.Next
		if fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		if fast == slow {
			break
		}
	}

	if fast == nil {
		return nil
	}

	fast = head
	// 再次交汇的点就应该是begin point
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}

	return fast
}
