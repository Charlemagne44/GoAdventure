package Model

type location struct {
	name      string
	neighbors []string
}

func (location) getName(loc *location) string {
	return loc.name
}

func (location) getNeighbors(loc *location) []string {
	return loc.neighbors
}
