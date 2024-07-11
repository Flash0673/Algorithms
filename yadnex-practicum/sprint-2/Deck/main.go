package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"container/list"
)

func main() {
	l := list.New()
	_ = l
	reader := bufio.NewReader(os.Stdin)
	var n int
	var maxSize int

	fmt.Scan(&n)
	fmt.Scan(&maxSize)

	q := NewDEQ(maxSize)

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
		case "pop_back":
			v := q.popBack()
			if v != nil {
				fmt.Println(*v)
			} else {
				fmt.Println("error")
			}
		case "push_back":
			q.pushBack(val)
		case "pop_front":
			v := q.popFront()
			if v != nil {
				fmt.Println(*v)
			} else {
				fmt.Println("error")
			}
		case "push_front":
			q.pushFront(val)
		}

		// fmt.Println(q.queue)
		// q.printQ()
	}

}

type MyDEQSized struct {
	maxSize int
	size    int
	head    *int
	tail    *int
	queue   []*int
}

func NewDEQ(maxSize int) MyDEQSized {
	queue := MyDEQSized{}
	head := 0
	tail := 0
	queue.maxSize = maxSize
	queue.size = 0
	queue.head = &head
	queue.tail = &tail
	queue.queue = make([]*int, maxSize)
	return queue
}

func (q *MyDEQSized) pushBack(x int) {
	if q.size == q.maxSize {
		fmt.Println("error")
		return
	}

	q.queue[*q.tail] = &x
	*q.tail = (*q.tail + 1) % q.maxSize
	q.size++
}

func (q *MyDEQSized) pushFront(x int) {
	if q.size == q.maxSize {
		fmt.Println("error")
		return
	}

	*q.head = (*q.head - 1) % q.maxSize
	if *q.head < 0 {
		*q.head = q.maxSize + *q.head
	}

	q.queue[*q.head] = &x
	q.size++
}

func (q *MyDEQSized) popBack() *int {
	if q.isEmpty() {
		return nil
	}

	*q.tail = (*q.tail - 1) % q.maxSize
	if *q.tail < 0 {
		*q.tail = q.maxSize + *q.tail
	}

	x := q.queue[*q.tail]
	q.queue[*q.tail] = nil

	q.size--

	return x
}

func (q *MyDEQSized) popFront() *int {
	if q.isEmpty() {
		return nil
	}

	x := q.queue[*q.head]
	q.queue[*q.head] = nil
	*q.head = (*q.head + 1) % q.maxSize
	q.size--

	return x
}

func (q *MyDEQSized) isEmpty() bool {
	return q.size == 0
}

func (q *MyDEQSized) printQ() {
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
