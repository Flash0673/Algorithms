package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n)
	fmt.Scan(&m)
	
	matrix := initMatrix(n, m)
	for i := range matrix {
		for j := range matrix[i] {
			fmt.Scan(&matrix[i][j])
		}
	}

	printMatrix(matrix)
	fmt.Println()
	
	transposed := transposeMatrix(n, m, matrix)
	printMatrix(transposed)
}

func initMatrix(n, m int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = append(matrix[i], make([]int, m)...)
	}
	return matrix
}

func printMatrix(matrix [][]int)  {
	for i := range matrix {
		for j := range matrix[i] {
			fmt.Printf("%d ", matrix[i][j])
		}
		fmt.Println()
	}
}

func transposeMatrix(n, m int, matrix [][]int) [][]int {
	transposed := initMatrix(m, n)
	for i := range matrix {
		for j := range matrix[i] {
			transposed[j][i] = matrix[i][j]
		}
	}
	return transposed
}
