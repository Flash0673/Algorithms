package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	res := solution(n)
	fmt.Println(res)
}

func solution(n int) bool {
	for i := 1; i <= n; i*=4 {
		if i == n {
			return true
		}
	}

	return false
}