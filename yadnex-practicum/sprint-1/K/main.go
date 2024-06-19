package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)

	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')

	splitted := strings.Split(line, " ")
	lstForm := make([]int, 0)

	for _, v := range splitted {
		n, _ := strconv.ParseInt(strings.TrimSpace(v), 10, 32)
		lstForm = append(lstForm, int(n))
	}

	line, _ = reader.ReadString('\n')
	k, _ := strconv.ParseInt(strings.TrimSpace(line), 10, 32)


	fmt.Println(n, lstForm, k)
	res := solution(n, lstForm, int(k))

	fmt.Println(res)
}

func solution(n int, form []int, k int) []int {
	var x int
	for _, v := range form {
		x *= 10
		x += v
	}
	res := x + k
	newForm := make([]int, 0)
	for res != 0 {
		newForm = append(newForm, res % 10)
		res /= 10
	}

	slices.Reverse(newForm)
	return newForm
}