package main

import (
	"bufio"
	"game/model"
	"game/view"
	"os"
)

func main() {
	var m model.Model
	var v view.View
	m.World.InitializeWorld("resources/world.json")

	// main game loop
	running := true
	for running {
		// get input via controller
		raw := takeRawInput()

		// parse into 3 cats via model
		directAct, indirectObj, directObj := m.Parse(raw)

		// execute actions and modify world state via model
		m.Exec(directAct, indirectObj, directObj)

		// output results via view
		v.ShowState()

		// DEBUG
		running = false
	}

}

func takeRawInput() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	return text
}
