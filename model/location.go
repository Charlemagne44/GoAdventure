package model

type Location struct {
	Name      string     `json:"name"`
	Neighbors []Location `json:"locations"`
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
	// DEBUG
	// log.Printf("l name: %v\n", l.Name)
	for _, n := range l.Neighbors {
		// DEBUG
		// log.Printf("neighbor name: %v\n", n.Name)
		if n.Name == loc2.Name {
			return true
		}
	}
	return false
}
