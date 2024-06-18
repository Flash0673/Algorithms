package main

import (
	"fmt"
	"strconv"
)

func main() {
	var first, second string
	fmt.Scan(&first)
	fmt.Scan(&second)

	f := make([]int, 0)
	s := make([]int, 0)

	for _, v := range first {
		n, _ := strconv.ParseInt(string(v), 10, 32)
		f = append(f, int(n))
	}

	for _, v := range second {
		n, _ := strconv.ParseInt(string(v), 10, 32)
		s = append(s, int(n))
	}

	res := solution1(f, s)
	for i := len(res) - 1; i > -1; i-- {
		fmt.Print(res[i])
	}
	fmt.Println()

	res = solution2(first, second)
	for i := len(res) - 1; i > -1; i-- {
		fmt.Print(res[i])
	}
	fmt.Println()
}

func solution1(f, s []int) []int {
	rem := 0

	// Определим какое число меньше
	var m bool
	if len(f) < len(s) {
		m = true
	} else {
		m = false
	}

	res := make([]int, 0)
	if m {
		for i := 0; i < len(f); i++ {
			tmp := f[len(f)-1-i] + s[len(s)-1-i] + rem
			res = append(res, tmp % 2)
			rem = tmp / 2
		}

		if rem > 0 {
			for i := len(f); i < len(s); i++ {
				tmp := s[len(s)-1-i] + rem
				res = append(res, tmp % 2)
				rem = tmp / 2
			}
		}

		if rem > 0 {
			res = append(res, 1)
		}
	} else {
		for i := 0; i < len(s); i++ {
			tmp := f[len(f)-1-i] + s[len(s)-1-i] + rem
			res = append(res, tmp % 2)
			rem = tmp / 2
		}

		if rem > 0 {
			for i := len(s); i < len(f); i++ {
				tmp := f[len(f)-1-i] + rem
				res = append(res, tmp % 2)
				rem = tmp / 2
			}
		}

		if rem > 0 {
			res = append(res, 1)
		}
	}

	return res
}

func solution2(first, second string) []int {
	var f,s int
	n, _ := strconv.ParseInt(first, 10, 32)
	f = int(n)
	n, _ = strconv.ParseInt(second, 10, 32)
	s = int(n)

	res := make([]int, 0)

	rem := 0
	for f != 0 || s != 0 || rem != 0 {
		lastF := f % 10
		lastS := s % 10
		tmp := lastF + lastS + rem
		res = append(res, tmp % 2)
		rem = tmp / 2
		f /= 10
		s /= 10
	}

	return res
}
