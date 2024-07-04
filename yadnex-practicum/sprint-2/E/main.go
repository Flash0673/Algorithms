package main

import "fmt"

func main() {
	n3 := &ListNode{data: "Third", next: nil}
	n2 := &ListNode{data: "Second", next: n3}
	n1 := &ListNode{data: "First", next: n2}
	n3.prev = n2
	n2.prev = n1
	n1.prev = nil

	head := n1
	_ = head


	printList(head)

	head = reverseList(head)
	printList(head)
}

func reverseList(head *ListNode) *ListNode {
	curr := head
	var newHead *ListNode
	for curr != nil {
		if curr.next == nil {
			newHead = curr
		}
		prev := curr.prev
		curr.prev = curr.next
		curr.next = prev
		curr = curr.prev
	}

	return newHead
}

func printList(head *ListNode) {
	curr := head
	fmt.Print(nil, " <-> ")
	for curr != nil {
		fmt.Printf("%s <-> ", curr.data)
		curr = curr.next
	}
	fmt.Println(nil)
}

type ListNode struct {
	data string
	next *ListNode
	prev *ListNode
}