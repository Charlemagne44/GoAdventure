package model

type Location struct {
	Name      string
	Neighbors []Location
}

func (inp Location) getName() string {
	return inp.Name
}

func (inp Location) getNeighbors() []Location {
	return inp.Neighbors
}

func (inp Location) checkIfNeighbor(currLoc, newLoc Location) bool {
	for _, neighbor := range inp.Neighbors {
		if neighbor.Name == newLoc.Name {
			return true
		}
	}
	return false
}
