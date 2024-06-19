package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	res := solution(n)

	for i := len(res) - 1; i > -1; i-- {
		fmt.Print(res[i], " ")
	}
	fmt.Println()
}

func simple_ints(n int) []int {
	nums := make([]int, n+1)
	for i := 0; i < n+1; i++ {
		nums[i] = i
	}

	nums[0], nums[1] = 0, 0

	for i := 2; i*i < n; i++ {

		if nums[i] != 0 {
			for j := i*i; j < n + 1; j+=nums[i] {
				nums[j] = 0
			}
		}
	}

	simp := make([]int, 0)
	for _, v := range nums {
		if v != 0 {
			simp = append(simp, v)
		}
	}

	return simp
}

func solution(n int) []int {
	nums := simple_ints(n)
	fmt.Println(nums)

	factors := make([]int, 0)

	i := len(nums) - 1

	for n != 1 {
		if n % nums[i] == 0 {
			factors = append(factors, nums[i])
			n /= nums[i]
		} else { 
			i--
		}
	}

	return factors
}