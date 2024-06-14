package main

import "fmt"

func main()  {
	testCases := []string{"шалаш", "tot", "hello", ""}
	for _, c := range testCases {
		fmt.Println(c, solution(c))
	}
}

func solution(s string) bool {
	runes := []rune(s)
	l, r := 0, len(runes) - 1

	for l <= r {
		if runes[l] != runes[r] {
			return false
		}
		l++
		r--
	}

	return true
}