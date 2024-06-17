package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	digit := make([]int, 0)

	for n != 0 {
		rem := n % 2
		n = n / 2
		digit = append(digit, rem)
	}

	for i := len(digit) - 1; i > -1; i-- {
		fmt.Print(digit[i])
	}
	fmt.Println()
}
