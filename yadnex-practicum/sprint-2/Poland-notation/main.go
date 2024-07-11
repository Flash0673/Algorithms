package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	s := strings.Split(strings.TrimSpace(line), " ")

	st := Stack{}

	for _, c := range s {
		if unicode.IsDigit([]rune(c)[0]) {
			n, _ := strconv.Atoi(c)
			st.append(n)
		} else {
			fir, sec := st.pop(), st.pop()
			var r int
			switch c {
			case "+":
				r = fir + sec
			case "-":
				r = fir + sec
			case "*":
				r = fir * sec
			case "/":
				r = fir / sec
			}
			st.append(r)
		}
	}

	res := st.pop()
	fmt.Println(res)
}

type Stack struct {
	stack []int
}

func (st *Stack) append(a int) {
	st.stack = append(st.stack, a)
}

func (st *Stack) pop() int {
	if !st.isEmpty() {
		el := st.stack[len(st.stack)-1]
		st.stack = st.stack[:len(st.stack)-1]
		return el
	}
	return 0
}

func (st *Stack) isEmpty() bool {
	return len(st.stack) == 0
}