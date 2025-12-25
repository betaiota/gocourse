package tasks

import "fmt"

func convertToKgs(num int, M float64) float64 {
	switch num {
	case 1:
		return M
	case 2:
		return M * 0.00001
	case 3:
		return M * 0.001
	case 4:
		return M * 1000
	case 5:
		return M * 100
	default:
		return -1.0
	}
}

func LabTwelve() {
	var mass float64
	var num int
	fmt.Println("Введите единицу измерения массы (1 - килограмм, 2 - миллиграмм, 3 - грамм, 4 - тонна, 5 - центнер):")
	_, err := fmt.Scan(&num)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Введите массу:")
	_, err = fmt.Scan(&mass)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	if mass < 0 || num > 5 || num < 1 {
		fmt.Println("Вы ввели неверную массу или номер единицы измерения!")
		return
	} else {
		fmt.Printf("Масса - %f кг.\n", convertToKgs(num, mass))
	}
}
