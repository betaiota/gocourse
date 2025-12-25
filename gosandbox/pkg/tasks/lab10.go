package tasks

import "fmt"

func LabTen() {
	fmt.Println("Введите число:")
	var num int
	if _, err := fmt.Scan(&num); err != nil {
		panic(err)
	}
	isOdd := func(num int) bool {
		return num%2 != 0
	}
	isEven := func(num int) bool {
		return num%2 == 0
	}

	fmt.Printf("Число четное? - %v\nЧисло нечетное? - %v\n", isEven(num), isOdd(num))
}
