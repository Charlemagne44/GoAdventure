package view

import (
	"fmt"
	"game/model"
)

type View struct{}

func (v *View) ShowState(m *model.Model) {
	printDesc(m)
}

func printDesc(m *model.Model) {
	fmt.Println(m.World.CurrLocation.Desc)
}
