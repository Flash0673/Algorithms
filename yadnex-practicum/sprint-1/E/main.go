package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	fmt.Scan(&n)

	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')

	s := []rune(line)

	var longest []rune
	var count int
	_, _ = longest, count

	l, r := 0, 0

	for r < len(s) + 1 {
		if r == len(s) || s[r] == ' ' {
			c := r - l 
			if c > count {
				longest = s[l:r]
				count = c
			}
			r++
			l = r
		} else {
			r++
		} 
	}

	fmt.Println(string(longest))
	fmt.Println(count)
}