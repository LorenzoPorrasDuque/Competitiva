package main

import (
	"fmt"
)

func main() {
	var nSequence int
	fmt.Scan(&nSequence)
	var numbers []int
	for {
		var input int
		fmt.Scan(&input)
		if input == 0 {
			break
		}
		numbers = append(numbers, input)
	}
	start, max, sum := calculate(numbers)
	fmt.Println(numbers)
	fmt.Println(start, max, sum)

}

func calculate(numbers []int) (a, b, c int) {
	var init, maxWinStreak, sum, localmax int
	init = numbers[0]
	for i := 0; i < len(numbers); i++ {
		sum = sum + numbers[i]
		if numbers[i] > 0 {
			localmax = localmax + numbers[i]
		} else {
			if localmax > maxWinStreak {
				maxWinStreak = localmax
			}
			localmax = sum
		}

	}
	return init, maxWinStreak, sum

}
