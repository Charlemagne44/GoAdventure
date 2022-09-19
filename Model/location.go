package Model

type location struct {
	name      string
	neighbors []location
}

func (inp *location) getName() string {
	return inp.name
}

func (inp *location) getNeighbors() []location {
	return inp.neighbors
}

func (inp *location) checkIfNeighbor(currLoc, newLoc *location) bool {
	for _, neighbor := range inp.neighbors {
		if neighbor.name == newLoc.name {
			return true
		}
	}
	return false
}
