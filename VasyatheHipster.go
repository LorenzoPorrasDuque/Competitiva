package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	switch a >= b {
	case true:
		a -= b
		a = a / 2
		fmt.Printf("%d %d\n", a, b)
	case false:
		b -= a
		b = b / 2
		fmt.Printf("%d %d\n", a, b)
	}
}
