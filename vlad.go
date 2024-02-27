package main

import (
	"fmt"
	"strconv"
)

func main() {

	var cases int

	fmt.Scan(&cases)
	numbers := make([]int, cases)

	for i := 0; i < cases; i++ {
		fmt.Scan(&numbers[i])
	}

	sums := make([]int, findMax(numbers)+1)

	sums[0] = 0
	sums[1] = 1
	sums[2] = 3
	sums[3] = 6
	sums[4] = 10
	sums[5] = 15
	sums[6] = 21
	sums[7] = 28
	sums[8] = 36
	sums[9] = 45

	for i := 10; i < numbers[len(numbers)-1]+1; i++ {
		sums[i] = sums[i-1] + calculateSumOfDigits(i)
	}
	for i := 0; i < cases; i++ {
		fmt.Println(sums[numbers[i]])
	}

	//fill sum

}

func calculateSumOfDigits(number int) int {
	strNumber := strconv.Itoa(number)
	sum := 0

	for _, digitStr := range strNumber {
		digit, _ := strconv.Atoi(string(digitStr))
		sum += digit
	}

	return sum
}

func findMax(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	max := numbers[0]
	for _, num := range numbers[1:] {
		if num > max {
			max = num
		}
	}

	return max
}
