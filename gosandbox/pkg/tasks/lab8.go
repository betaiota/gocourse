package tasks

import "fmt"

func LabEight() bool {
	fmt.Println("Введите число:")
	var res string
	_, err := fmt.Scan(&res)
	if err != nil {
		fmt.Println(err)
	}
	rs := []rune(res)
	lrs := rs[:len(rs)/2]
	offset := 0
	if len(rs)%2 != 0 {
		offset = 1
	}
	rrs := rs[len(rs)/2+offset:]
	for i, j := 0, len(rrs)-1; i < len(lrs); {
		if lrs[i] != rrs[j] {
			return false
		}
		i++
		j--
	}
	return true
}
