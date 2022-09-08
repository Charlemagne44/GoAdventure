package Model

type location struct {
	name      string
	neighbors []location
}

func (inp *location) getName(loc *location) string {
	return loc.name
}

func (inp *location) getNeighbors(loc *location) []location {
	return loc.neighbors
}

func (inp *location) checkIfNeighbor(currLoc, newLoc *location) bool {
	for _, neighbor := range currLoc.neighbors {
		if neighbor.name == newLoc.name {
			return true
		}
	}
	return false
}
