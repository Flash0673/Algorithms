package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	line, _ := reader.ReadString('\n')
	n, _ := strconv.ParseInt(strings.TrimSpace(line), 10, 32)

	for i := 0; i < int(n); i++ {
		line, _ := reader.ReadString('\n')
		splitted := strings.Split(strings.TrimSpace(line), " ")
		s := 0
		for i := range splitted {
			v, _ := strconv.ParseInt(splitted[i], 10, 32)
			s += int(v)
		}
		fmt.Fprintf(writer, "%d\n", s)
	}

	writer.Flush()
}
