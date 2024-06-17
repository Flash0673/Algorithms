package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	s, _ := reader.ReadString('\n')

	ans := solution(s)

	fmt.Println(ans)
}

func solution(line string) bool {
	s := []rune(strings.ToLower(line))
	l, r := 0, len(s) - 1
	var hasLetters bool = false // Есть ли буквы

	for l < r {
		if !unicode.IsLetter(s[l]) {
			l++
			continue
		}
		if !unicode.IsLetter(s[r]) {
			r--
			continue
		}

		if s[l] == s[r] {
			l++
			r--
			hasLetters = true
		} else {
			return false
		}
	}
	if hasLetters {
		return true
	} 
	return false
}