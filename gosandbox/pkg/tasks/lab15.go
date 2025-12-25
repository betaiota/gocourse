package tasks

import "fmt"

func LabFifteen() {
	var arr = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i, v := range arr {
		arr[i] = v + 1
	}
	fmt.Println(arr)

}
