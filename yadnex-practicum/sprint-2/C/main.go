package main

import "fmt"

func main() {
	n3 := &ListNode{data: "Third", next: nil}
	n2 := &ListNode{data: "Second", next: n3}
	n1 := &ListNode{data: "First", next: n2}

	head := n1
	_ = head


	printList(head)
	idx := 3
	head = deleteNode(head, idx)
	printList(head)
}

func deleteNode(head *ListNode, idx int) *ListNode {
	if idx == 0 {
		head = head.next
		return head
	}

	prev := head
	for idx-1 > 0 {
		prev = prev.next
		idx--
	}

	tmp := prev.next
	prev.next = tmp.next
	tmp.next = nil

	return head
}

func printList(head *ListNode) {
	curr := head
	for curr != nil {
		fmt.Printf("%s -> ", curr.data)
		curr = curr.next
	}
	fmt.Println(nil)
}

type ListNode struct {
	data string
	next *ListNode
}