package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n, s int
	fmt.Scan(&n)

	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	l := strings.Split(strings.TrimSpace(line), " ")

	w := make([]int, len(l))
	for i, v := range l {
		n, _ := strconv.Atoi(v)
		w[i] = n
	}

	line, _ = reader.ReadString('\n')
	s, _ = strconv.Atoi(strings.TrimSpace(line))

	// fmt.Println(n)
	// fmt.Println(w)
	// fmt.Println(s)
	res1 := findDay(w, s, 0, n-1, -1)
	res2 := findDay(w, s*2, 0, n-1, -1)

	if res1 == -1 { fmt.Print(-1, " ") } else {fmt.Print(res1+1, " ")}
	if res2 == -1 { fmt.Println(-1)}  else {fmt.Println(res2+1)}
}

func findDay(arr []int, target, l, r, idx int) int {
	if r <= l {
		return idx
	}

	mid := (l+r)/2
	if arr[mid] >= target {
		idx = mid
		return findDay(arr, target, l, mid-1, idx)
	} else {
		return findDay(arr, target, mid+1, r, idx)
	}
}