package main

import "fmt"

func main() {
	arr := []int{1,2,3,4,5}
	t := 6

	res := binSearch(arr, t, 0, len(arr))
	fmt.Println(res)
	
}

func binSearch(arr []int, target, l, r int) int {
	// recursion base
	if r <= l {
		return -1
	}

	mid := (l + r) / 2
	el := arr[mid]
	switch {
	case el == target:
		return mid
	case el < target:
		return binSearch(arr, target, l+1, r)
	case el > target:
		return binSearch(arr, target, l, r-1)
	}

	return -1
}