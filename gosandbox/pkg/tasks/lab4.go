package tasks

import (
	"fmt"
)

func biggestOfThree(x int, y int, z int) int {
	if x > y && x > z {
		return x
	} else if y > x && y > z {
		return y
	} else {
		return z
	}
}

func LabFour() {
	x := 77
	y := 9
	z := 130
	fmt.Println("Наибольшее число", biggestOfThree(x, y, z))
}
