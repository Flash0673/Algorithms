package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    s, _ := reader.ReadString('\n')
    splitted := strings.Split(s, " ")

    var ans string
    m := 0
    count := make(map[string]int, 0)

    for _, word := range splitted {
        l, r := 0, 2
        for r <= len(word) {
            count[word[l:r]]++
            if count[word[l:r]] > m {
                m = count[word[l:r]]
                ans = word[l:r]
            } else if count[word[l:r]] == m {
                ans = max(ans, word[l:r])
            }
            l++
            r++
        }
    }

    fmt.Println(count, ans, m)
}
