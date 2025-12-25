package tasks

import "fmt"

func LabSix() {
	fmt.Println("Введите сторону квадрата:")
	var res int
	_, err := fmt.Scan(&res)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Площадь квадрата равна:", res*res)
}
