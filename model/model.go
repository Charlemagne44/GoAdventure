package model

import "fmt"

type Model struct {
	Name  string
	World World
}

func init() {
	fmt.Println("DEBUG: Model initialized")
}

func (m Model) directControllerInput(command string) {
	//TODO
}

func (m Model) initialize() {
	//TODO
	m.World.initializeLocations("/Resources/world.json")
}

func test() {
	return
}
