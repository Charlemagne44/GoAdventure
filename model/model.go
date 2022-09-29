package model

type Model struct {
	World World
}

func (m *Model) Parse(raw string) (directAct, indirectObj, directObj string) {
	return "", "", ""
}

func (m *Model) Exec(directAct, indirectObj, directObj string) {

}

func (m *Model) Initialize() {

}
