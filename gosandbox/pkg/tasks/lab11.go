package tasks

import "fmt"

func LabEleven() {
	monthNum := 11
	switch monthNum {
	case 1:
		fmt.Println("Январь")
	case 2:
		fmt.Println("Февраль")
	case 3:
		fmt.Println("Март")
	case 4:
		fmt.Println("Апрель")
	case 5:
		fmt.Println("Май")
	case 6:
		fmt.Println("Июнь")
	case 7:
		fmt.Println("Июль")
	case 8:
		fmt.Println("Август")
	case 9:
		fmt.Println("Сентябрь")
	case 10:
		fmt.Println("Октябрь")
	case 11:
		fmt.Println("Ноябрь")
	case 12:
		fmt.Println("Декабрь")
	default:
		fmt.Println("Такого месяца нет в году!")
	}
}
