package main

import "fmt"

func main() {
	n3 := &ListNode{data: "Third", next: nil}
	n2 := &ListNode{data: "Second", next: n3}
	n1 := &ListNode{data: "First", next: n2}

	head := n1
	_ = head


	printList(head)
	val := "Third"

	fmt.Println(findVal(head, val))
	
}

func findVal(head *ListNode, val string) int {
	curr := head
	idx := 0

	for curr != nil {
		if curr.data == val {
			return idx
		}
		curr = curr.next
		idx++
	}
	return -1
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