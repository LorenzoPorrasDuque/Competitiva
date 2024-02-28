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
	var sum int

	for i1, ints1 := range matrix[0] {
		for i2, ints2 := range matrix[1] {
			for i3, ints3 := range matrix[2] {
				if i2 != i1 && i3 != i2 && ints1+ints2+ints3 > sum {
					sum = ints1 + ints2 + ints3
				}
			}
		}
	}
	return sum
}

/*
func findMaxFriends(matrix [][]int) int {
	var sum int
	for i := 0; i < len(matrix[0]); i++ {
		for j := 0; j < len(matrix[1]); j++ {
			for k := 0; k < len(matrix[2]); k++ {

			}
		}
	}

*/
