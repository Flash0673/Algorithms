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
	var maxSize int

	fmt.Scan(&n)
	fmt.Scan(&maxSize)

	q := NewQueue(maxSize)

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
		case "peek":
			v := q.peek()
			if v != nil {
				fmt.Println(*v)
			} else {
				fmt.Println("None")
			}
		case "push":
			q.push(val)
		case "pop":
			v := q.pop()
			if v != nil {
				fmt.Println(*v)
			} else {
				fmt.Println("None")
			}
		case "size":
			s := q.getSize()
			fmt.Println(s)
		}

		// fmt.Println(q.queue)
		// q.printQ()
	}

}

type MyQueueSized struct {
	maxSize int
	size    int
	head    *int
	tail    *int
	queue   []*int
}

func NewQueue(maxSize int) MyQueueSized {
	queue := MyQueueSized{}
	head := 0
	tail := 0
	queue.maxSize = maxSize
	queue.size = 0
	queue.head = &head
	queue.tail = &tail
	queue.queue = make([]*int, maxSize)
	return queue
}

func (q *MyQueueSized) push(x int) {
	if q.size == q.maxSize {
		fmt.Println("error")
		return
	}

	q.queue[*q.tail] = &x
	*q.tail = (*q.tail + 1) % q.maxSize
	q.size++
}

func (q *MyQueueSized) pop() *int {
	if q.isEmpty() {
		return nil
	}

	x := q.queue[*q.head]
	q.queue[*q.head] = nil
	*q.head = (*q.head + 1) % q.maxSize
	q.size--

	return x
}

func (q *MyQueueSized) peek() *int {
	if q.isEmpty() {
		return nil
	}
	return q.queue[*q.head]
}

func (q *MyQueueSized) getSize() int {
	return q.size
}

func (q *MyQueueSized) isEmpty() bool {
	return q.size == 0
}

func (q *MyQueueSized) printQ() {
	fmt.Print("Q: ")
	for _, v := range q.queue {
		if v != nil {
			fmt.Printf("%d ", *v)
		} else {
			fmt.Printf("nil ")
		}
	}
	fmt.Println()
}
