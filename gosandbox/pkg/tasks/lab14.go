package tasks

import "fmt"

func findMinimalDistances(mass []int) {
	resM := make(map[int]int)
	for i := 0; i < len(mass); i++ {
		v1 := mass[i]
		fmt.Println("Scanning element", v1)
		minDist := len(mass)
		prevI := i
		for j := i + 1; j < len(mass); j++ {
			v2 := mass[j]
			if v1 == v2 {
				fmt.Println("Index == ", j)
				minDist = min(minDist, j-prevI)
				resM[v1] = minDist
				prevI = j
			}
		}
	}
	fmt.Println(resM)
}

func LabFourteen() {
	var mass = [...]int{1, 2, 17, 54, 30, 89, 2, 1, 6, 2, 1}
	findMinimalDistances(mass[:])

}
