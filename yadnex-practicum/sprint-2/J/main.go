package main

import (
	"bufio"
	"fmt"
	"os"
	// "slices"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n int

	fmt.Scan(&n)

	q := Queue{}

	// fmt.Println(q.queue)
	// q.printQ()

	for i := 0; i < n; i++ {
		line, _ := reader.ReadString('\n')
		splitted := strings.Split(strings.TrimSpace(line), " ")
		var command string
		var num string
		var val int
		if len(splitted) < 2 {
			command = splitted[0]
		} else {
			command, num = splitted[0], splitted[1]
			v, _ := strconv.ParseInt(num, 10, 32)
			val = int(v)
		}

		switch command {
		case "put":
			q.put(val)
		case "get":
			v := q.get()
			if v != nil {
				fmt.Println(*v)
			} else {
				fmt.Println("error")
			}
		case "size":
			s := q.getSize()
			fmt.Println(s)
		}

		// fmt.Println(q.queue)
		// q.printQ()
	}

}

type ListNode struct {
	val int
	next *ListNode	
}

type Queue struct {
	head *ListNode
	tail *ListNode
	size int
}

func (q *Queue) get() *int {
	if q.size == 0 {
		return nil
	}

	x := &q.head.val
	q.head = q.head.next
	q.size--
	return x
}

func (q *Queue) put(x int) {
	if q.size == 0 {
		n := &ListNode{val: x}
		q.head = n
		q.tail = n
		q.size++
		return
	}
	n := &ListNode{val: x}
	q.tail.next = n
	q.tail = n
	q.size++
}

func (q *Queue) getSize() int {
	return q.size
}

func (q *Queue) printQ() {
	curr := q.head
	for curr != nil {
		fmt.Printf("%d -> ", curr.val)
		curr = curr.next
	}
	fmt.Println(nil)
}
