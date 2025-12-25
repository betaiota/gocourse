package tasks

import "fmt"

func LabOne() {
	var i int = 42
	var f float64 = 3.15
	var b bool = true
	var s string = "test"
	var a any

	fmt.Println(i, f, b, s, a)

	var pi *int = &i
	var pf *float64 = &f
	var pb *bool = &b
	var ps *string = &s
	var pa *any = &a

	fmt.Println(*pi, *pf, *pb, *ps, *pa)

	fmt.Printf("%T, %T, %T, %T, %T", *pi, *pf, *pb, *ps, *pa)
}
