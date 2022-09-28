package model

type Location struct {
	Name      string
	Neighbors []Location
}

func (l *Location) GetName() string {
	return l.Name
}

func (l *Location) GetNeighbors() []Location {
	return l.Neighbors
}

func (l *Location) CheckIfNeighbor(loc2 *Location) bool {
	// go through list of current loc neighbors, check to see
	// if it's in the list of neighbors
	for _, n := range l.Neighbors {
		if n.Name == l.Name {
			return true
		}
	}
	return false
}
