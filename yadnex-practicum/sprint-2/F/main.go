package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	s := Stack{}
	reader := bufio.NewReader(os.Stdin)
	var n int

	fmt.Scan(&n)

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

		// fmt.Println(s.items)

		switch command {
		case "get_max":
			m := s.getMax()
			if m != nil {
				fmt.Println(*m)
			} else {
				fmt.Println("None")
			}
		case "push":
			s.push(val)
		case "pop":
			s.pop()
		}
	}

}

type Stack struct {
	items []int
	max *int
}

func (s *Stack) push(item int) {
	s.items = append(s.items, item)
	if s.max == nil {
		n := 0
		s.max = &n
		*s.max = item
	} else if item > *s.max {
		*s.max = item
	}
}

func (s *Stack) pop() int {
	if s.isEmpty() {
		fmt.Println("error")
		return 0
	}
	var value int
	s.items, value = s.items[:len(s.items)-1], s.items[len(s.items)-1]
	if value == *s.max && !s.isEmpty() {
		m := slices.Max(s.items)
		*s.max = m
	} else if s.isEmpty() {
		s.max = nil
	}
	return value
}

func (s *Stack) isEmpty() bool {
	if len(s.items) > 0 {
		return false
	}

	return true
}

func (s *Stack) getMax() *int {
	if s.isEmpty() {
		return nil
	}

	return s.max
}
