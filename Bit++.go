package main

import "fmt"

func main() {
	var n, y int
	y = 0

	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var s string
		fmt.Scan(&s)
		if s == "X++" || s == "++X" {
			y++
		} else {
			y--
		}
	}
	print(y)

}
