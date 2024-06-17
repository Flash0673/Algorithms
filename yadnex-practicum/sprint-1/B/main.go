package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	s, _ := reader.ReadString('\n')

	splitted := strings.Split(strings.TrimSpace(s), " ")

	res := solution(splitted)

	if res {
		fmt.Println("WIN")
	} else {
		fmt.Println("FAIL")
	}

}

func solution(splitted []string) bool {
	num, _ := strconv.ParseInt(splitted[0], 10, 32)
	f := num % 2

	for _, v := range splitted[1:] {
		num, _ := strconv.ParseInt(v, 10, 32)
		if num%2 != f {
			return false
		}
	}

	return true
}
