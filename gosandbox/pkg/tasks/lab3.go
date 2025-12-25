package tasks

import (
	"fmt"
)

func LabThree() {
	age := 36.6
	temperature := float64(25)
	age, temperature = temperature, age
	fmt.Printf("Age - %v, temperature - %v", age, temperature)
}
