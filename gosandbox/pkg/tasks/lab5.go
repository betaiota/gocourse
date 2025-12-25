package tasks

import (
	"fmt"
	"math"
)

func arythm(x int, y int) {
	fmt.Printf("x = %v, y = %v\n", x, y)
	fmt.Printf("x + y = %v\n", x+y)
	fmt.Printf("x - y = %v\n", x-y)
	fmt.Printf("x * y = %v\n", x*y)
	fmt.Printf("x / y = %v\n", x/y)
	fmt.Printf("x %% y = %v\n", x%y)
	fmt.Printf("x ^ y = %v\n", math.Pow(float64(x), float64(y)))
}

func LabFive() {
	x := 4
	y := 5
	arythm(x, y)
}
