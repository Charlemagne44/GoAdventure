package Model

type model struct {
	world *world
}

func (m *model) directControllerInput(command string) {
	//TODO
}

func (m *model) initialize() {
	//TODO
	m.world.initialize()
}
