package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	s, _ := reader.ReadString('\n')
	t, _ := reader.ReadString('\n')
	
	res := solution(s, t)
	fmt.Println(res)
}

func solution(s, t string) string {
	sDict := make(map[rune]int)
	tDict := make(map[rune]int)

	for _, v := range s {
		sDict[v]++
	}

	for _, v := range t {
		tDict[v]++
	}

	for k, tValue := range tDict {
		sValue, ok := sDict[k]
		if !ok || sValue != tValue {
			return string(k)
		}
	}

	return ""
}