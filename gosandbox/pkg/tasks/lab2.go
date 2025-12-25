package tasks

import (
	"fmt"
	"strconv"
)

// parse int to bool
func itob(i int) bool {
	if i == 0 {
		return true
	} else {
		return false
	}
}

func LabTwo() {
	var age = "23"
	var foo = "23abc"
	ageInt, err := strconv.Atoi(age)
	if err != nil {
		fmt.Println(err)
	}
	fooInt, err := strconv.Atoi(foo) //err: parsing "23abc": invalid syntax
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v, %v\n", ageInt, fooInt) //23, 0
	fmt.Printf("%T, %T\n", ageInt, fooInt) //int, int

	age = "123abc"
	ageBool, err := strconv.ParseBool(age)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v\n", ageBool) //false
	fmt.Printf("%T\n", ageBool) //bool

	flag := 1
	//flagBool := bool(flag) - такого приведения нет в языке, равно как
	//и метода strconv.Itob. Определяем сами.
	flagBool := itob(flag)
	fmt.Printf("%v\n", flagBool) //true
	fmt.Printf("%T\n", flagBool) //bool

	str_one := "Privet"
	str_two := ""

	str1Bool, err := strconv.ParseBool(str_one)
	str2Bool, err := strconv.ParseBool(str_two)

	fmt.Printf("%v, %v\n", str1Bool, str2Bool) //false, false
	fmt.Printf("%T, %T\n", str1Bool, str2Bool) //bool, bool

	zeroToBool := itob(0)
	oneToBool := itob(1)
	fmt.Printf("%v, %v\n", zeroToBool, oneToBool) //true, false
	fmt.Printf("%T, %T\n", zeroToBool, oneToBool) //bool, bool

	falseToStr := strconv.FormatBool(false)

	fmt.Printf("%v\n", falseToStr) //false
	fmt.Printf("%T\n", falseToStr) //string
}
