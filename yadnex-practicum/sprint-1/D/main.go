package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)

	reader := bufio.NewReader(os.Stdin)

	line, _ := reader.ReadString('\n')
	splitted := strings.Split(line, " ")

	temps := make([]int, n)

	for i, v := range splitted {
		n, _ := strconv.ParseInt(v, 10, 32)
		temps[i] = int(n)
	}

	fmt.Println(temps)

	ans := solution(temps)
	fmt.Println(ans)
}


func solution(temps []int) int {
	var count int

	for i := range temps {
		switch i {
		case 0:
			if temps[i+1] < temps[i] {count++} 
		case len(temps) - 1:
			if temps[i-1] < temps[i] {count++} 
		default:
			if temps[i-1] < temps[i] && temps[i+1] < temps[i] {count++}
		}
	}

	return count
}