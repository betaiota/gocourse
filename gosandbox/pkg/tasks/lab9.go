package tasks

import "fmt"

func LabNine() {
	fmt.Println("Решение линейного ур-я Ax + B = 0")
	fmt.Println("Введите коэффициент A:")
	var A, B float64
	_, err := fmt.Scan(&A)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Введите коэффициент B:")
	_, err = fmt.Scan(&B)
	if err != nil {
		fmt.Println(err)
	}
	x := -B / A
	fmt.Printf("%vx + %v = 0; x = %v\n", A, B, x)

}
