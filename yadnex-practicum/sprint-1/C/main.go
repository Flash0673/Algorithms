package main

import (
	"fmt"
	// "strings"
	// "bufio"
	// "strconv"
)

func main() {
	var n, m int
	fmt.Scan(&n)
	fmt.Scan(&m)

	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			matrix[i] = append(matrix[i], 0)
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Scan(&matrix[i][j])
		}
	}

	// fmt.Println(matrix)

	var row, col int
	fmt.Scan(&row)
	fmt.Scan(&col)

	neighbours := make([]int, 0)

	if col > 0 {
		neighbours = append(neighbours, matrix[row][col-1])
	}
	if col < m - 1 {
		neighbours = append(neighbours, matrix[row][col+1])
	}
	if row > 0 {
		neighbours = append(neighbours, matrix[row-1][col])
	}
	if row < n - 1 {
		neighbours = append(neighbours, matrix[row+1][col])
	}

	fmt.Println(neighbours)
}