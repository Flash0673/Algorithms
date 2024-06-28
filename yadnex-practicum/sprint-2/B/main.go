package main

import "fmt"

func main() {
	n3 := ListNode{data: "Third", next: nil}
	n2 := ListNode{data: "Second", next: &n3}
	n1 := ListNode{data: "First", next: &n2}

	head := n1
	for head.next != nil {
		fmt.Printf("%s -> ", head.data)
		head = *head.next
	}
	fmt.Printf("%s\n", head.data)
}

type ListNode struct {
	data string
	next *ListNode
}