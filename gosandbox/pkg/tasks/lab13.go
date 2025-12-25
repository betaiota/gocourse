package tasks

import (
	"bytes"
	"fmt"
)

const (
	wood  = "син"
	fire  = "красн"
	earth = "желт"
	metal = "бел"
	water = "черн"
)

const (
	rat    = "крысы"
	bull   = "быка"
	tiger  = "тигра"
	cat    = "кота"
	dragon = "дракона"
	cock   = "петуха"
	snake  = "змеи"
	horse  = "лошади"
	sheep  = "овцы"
	monke  = "обезьяны"
	dog    = "собаки"
	pig    = "свиньи"
)

const (
	endingF1 = "ей "
	endingF2 = "ой "
	endingM1 = "его "
	endingM2 = "ого "
)

func calcChineseYear(year int) string {
	var buffer bytes.Buffer
	buffer.WriteString("Год ")
	modifiedYear := year + 2397
	last := modifiedYear % 60
	fmt.Println("Last = ", last)
	llast := last % 10
	fmt.Println("Last number of last = ", llast)
	var isBlue bool = false
	if llast == 1 || llast == 2 {
		isBlue = true
		buffer.WriteString(wood)
	} else if llast == 3 || llast == 4 {
		buffer.WriteString(fire)
	} else if llast == 5 || llast == 6 {
		buffer.WriteString(earth)
	} else if llast == 7 || llast == 8 {
		buffer.WriteString(metal)
	} else if llast == 9 || llast == 0 {
		buffer.WriteString(water)
	}

	cycleLast := last % 12
	switch cycleLast {
	case 0:
		if isBlue {
			buffer.WriteString(endingF1)
		} else {
			buffer.WriteString(endingF2)
		}
		buffer.WriteString(pig)
	case 1:
		if isBlue {
			buffer.WriteString(endingF1)
		} else {
			buffer.WriteString(endingF2)
		}
		buffer.WriteString(rat)
	case 2:
		if isBlue {
			buffer.WriteString(endingM1)
		} else {
			buffer.WriteString(endingM2)
		}
		buffer.WriteString(bull)
	case 3:
		if isBlue {
			buffer.WriteString(endingM1)
		} else {
			buffer.WriteString(endingM2)
		}
		buffer.WriteString(tiger)
	case 4:
		if isBlue {
			buffer.WriteString(endingM1)
		} else {
			buffer.WriteString(endingM2)
		}
		buffer.WriteString(cat)
	case 5:
		if isBlue {
			buffer.WriteString(endingM1)
		} else {
			buffer.WriteString(endingM2)
		}
		buffer.WriteString(dragon)
	case 6:
		if isBlue {
			buffer.WriteString(endingF1)
		} else {
			buffer.WriteString(endingF2)
		}
		buffer.WriteString(snake)
	case 7:
		if isBlue {
			buffer.WriteString(endingF1)
		} else {
			buffer.WriteString(endingF2)
		}
		buffer.WriteString(horse)
	case 8:
		if isBlue {
			buffer.WriteString(endingF1)
		} else {
			buffer.WriteString(endingF2)
		}
		buffer.WriteString(sheep)
	case 9:
		if isBlue {
			buffer.WriteString(endingF1)
		} else {
			buffer.WriteString(endingF2)
		}
		buffer.WriteString(monke)
	case 10:
		if isBlue {
			buffer.WriteString(endingF1)
		} else {
			buffer.WriteString(endingF2)
		}
		buffer.WriteString(dog)
	case 11:
		if isBlue {
			buffer.WriteString(endingM1)
		} else {
			buffer.WriteString(endingM2)
		}
		buffer.WriteString(cock)
	}

	return buffer.String()
}

func LabThirteen() {
	year := 2026
	fmt.Println(calcChineseYear(year))
}
