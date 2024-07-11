package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	s := strings.Split(strings.TrimSpace(line), " ")
	n, _ := strconv.Atoi(s[0])
	k, _ := strconv.Atoi(s[1])

	res := fib(n)
	fmt.Println(res % int(math.Pow10(k)))
}

func fib(n int) int {
	l, r := 1, 1
	for i := 2; i <= n; i++ {
		l, r = r, l+r
	}

	return r
}