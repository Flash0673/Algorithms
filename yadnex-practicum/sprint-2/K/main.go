package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	res := fib(n+1)
	fmt.Println(res)
}

func fib(n int) int {
	if n <= 2 {
		return 1
	}

	return fib(n-1) + fib(n-2)
}