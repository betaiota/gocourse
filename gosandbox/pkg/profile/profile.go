package profile

import (
	"fmt"

	"github.com/google/uuid"
)

type IBase interface {
	ChangeColor(color string)
	PrintColor()
}

type base struct {
	color string
}

func (b *base) ChangeColor(color string) {
	b.color = color
}

func (b base) PrintColor() {
	fmt.Println(b.color)
}

type derived struct {
	base
	size int
}

func Test() {
	d := derived{base{"black"}, 15}
	d.PrintColor()
	d.ChangeColor("violet")
	d.PrintColor()
	DoWithBase(&d)
	d.PrintColor()
}

type profile struct {
	uuid      uuid.UUID
	login     string
	passwd    string
	isBlocked bool
}

func CreateProfile(login string, passwd string, isBlocked bool) *profile {
	return &profile{uuid.New(), login, passwd, isBlocked}
}

func DoWithBase(base IBase) {
	base.ChangeColor("Green")
}

func (p profile) GetUUIDAsString() string {
	return p.uuid.String()
}

func (p *profile) ChangeBlockStatus(newStatus bool) {
	p.isBlocked = newStatus
}
