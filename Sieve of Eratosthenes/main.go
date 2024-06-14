package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	
	// fmt.Println(arr)
	solution(n)
}

func createEmptyArr(n int) []int {
	arr := make([]int, n+1)
	for i := 0; i <= n; i++ {
		arr[i] = i
	}

	return arr
}

func solution(n int) {
	arr := createEmptyArr(n)

	if n < 2 {
		return
	}

	arr[0], arr[1] = 0, 0

	i := 2

	for i := i; i < n; i++ {
		if i*i > n {
			break
		}

		if arr[i] != 0 {
			for j := i*i; j < n + 1; j += i {
				arr[j] = 0
			}

		} 
	}

	fmt.Println(arr)
}