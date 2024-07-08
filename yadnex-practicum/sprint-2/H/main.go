package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	res := isValid(s)
	fmt.Println(res)
}

func isValid(s string) bool {
	stack := make([]rune, 0, len(s))
	r := []rune(s)
	for _, v := range r {
		if len(stack) != 0 {
			t := stack[len(stack)-1]
			if v == '}' && t == '{' || v == ']' && t == '[' || v == ')' && t == '(' {
				stack = stack[:len(stack)-1]
			} else {
				stack = append(stack, v)
			}

		} else {
			stack = append(stack, v)
		}
	}

	if len(stack) == 0 {
		return true
	} else {
		return false
	}
}
