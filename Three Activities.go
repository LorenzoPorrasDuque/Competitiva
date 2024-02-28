package main

import "fmt"

func main() {
	var testCases int
	fmt.Scan(&testCases)
	for i := 0; i < testCases; i++ {
		var n int
		fmt.Scan(&n)
		matrix := createMatrix(n)
		fmt.Println(findMaxFriends(matrix))
	}
}

func createMatrix(n int) [][]int {
	matrix := make([][]int, 3)
	for i := 0; i < 3; i++ {
		matrix[i] = make([]int, n)
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < n; j++ {
			fmt.Scan(&matrix[i][j])
		}
	}
	return matrix
}
func findMaxFriends(matrix [][]int) int {
	var max, ubi, sum int

	return sum
}
