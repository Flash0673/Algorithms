package main

import "fmt"

func main() {
	// s1 := []int{1,2,3,4,5}
	// s2 := s1[1:3]

	// fmt.Println("S1:", s1, len(s1), cap(s1))
	// fmt.Println("S2:", s2, len(s2), cap(s2))
	fmt.Println(removeElement([]int{1,2,3,4}, 0))
}

func removeElement(nums []int, val int) int {
	l, r := 0, len(nums) - 1
	k := len(nums)

	for l <= r {
			if nums[l] == val {
					nums[l], nums[r] = nums[r], nums[l]
					r--
					k--
			} else {
					l++
			}
	}
	return k
}