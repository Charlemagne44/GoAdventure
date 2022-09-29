package main

import (
	"game/model"
)

func main() {
	var m model.Model
	m.World.InitializeWorld("resources/world.json")
	// DEBUG
	// fmt.Println(m.World.GetCurrLocation().Name)
}

func determineCommand(command string) string {
	return "debug"
}
