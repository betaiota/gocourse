package tasks

import (
	"fmt"
	"math"
)

type Point struct {
	X int
	Y int
}

func (p *Point) calculateDistance(p1 *Point) float64 {
	return math.Sqrt((math.Pow(float64(p.X-p1.X), 2)) + (math.Pow(float64(p.Y-p1.Y), 2)))
}

func LabSeven() {
	pointA := new(Point)
	pointB := new(Point)
	pointC := new(Point)
	fmt.Println("Введите координаты точки A")
	fmt.Println("\tA.x:")
	var res int
	_, err := fmt.Scan(&res)
	if err != nil {
		fmt.Println(err)
	}
	pointA.X = res
	fmt.Println("\tA.y:")
	_, err = fmt.Scan(&res)
	if err != nil {
		fmt.Println(err)
	}
	pointA.Y = res
	fmt.Println("Введите координаты точки B")
	fmt.Println("\tB.x:")
	_, err = fmt.Scan(&res)
	if err != nil {
		fmt.Println(err)
	}
	pointB.X = res
	fmt.Println("\tB.y:")
	_, err = fmt.Scan(&res)
	if err != nil {
		fmt.Println(err)
	}
	pointB.Y = res
	fmt.Println("Введите координаты точки C")
	fmt.Println("\tC.x:")
	_, err = fmt.Scan(&res)
	if err != nil {
		fmt.Println(err)
	}
	pointC.X = res
	fmt.Println("\tC.y:")
	_, err = fmt.Scan(&res)
	if err != nil {
		fmt.Println(err)
	}
	pointC.Y = res

	fmt.Println(pointA, pointB, pointC)

	ac := pointA.calculateDistance(pointC)
	bc := pointB.calculateDistance(pointC)
	fmt.Println("Длина AC:", ac)
	fmt.Println("Длина BC:", bc)
	fmt.Println("Сумма отрезков:", ac+bc)
}
