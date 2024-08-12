package main

import (
	"fmt"
	"sort"
)

type kirito struct {
	strength int
	combats  int
}
type dragon struct {
	strength int
	bonus    int
}

func main() {
	kirito := new(kirito)
	var dragons []dragon
	var input int
	fmt.Scan(&input)
	kirito.strength += input
	fmt.Scan(&input)
	kirito.combats += input

	for i := 0; i < kirito.combats; i++ {
		dragonInput := new(dragon)
		fmt.Scan(&input)
		dragonInput.strength += input
		fmt.Scan(&input)
		dragonInput.bonus += input
		dragons = append(dragons, *dragonInput)
	}
	sort.Slice(dragons[:], func(i, j int) bool {
		return dragons[i].strength < dragons[j].strength
	})
	result := combatProgress(kirito, dragons)
	fmt.Println(result)

}

func combatProgress(kirito *kirito, dragons []dragon) string {

	for i := 0; i < kirito.combats; i++ {
		if kirito.strength > dragons[i].strength {
			kirito.strength += dragons[i].bonus
		} else if kirito.strength == dragons[i].strength {
			break
		} else {
			break
		}
		if i <= kirito.combats {
			return "YES"
		}

	}
	return "NO"

}
