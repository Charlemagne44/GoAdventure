package model

import "log"

type World struct {
	// the string that maps to Loc objects will just be the Location.Name
	Locations    map[string]Location
	CurrLocation Location
}

func (w *World) GetCurrLocation() Location {
	return w.CurrLocation
}

func (w *World) SetCurrLocation(newLoc Location) {
	// should make call to check neighbors to make sure this move is valid
	// must be a neighbor in order to move the location
	if w.CurrLocation.CheckIfNeighbor(&newLoc) {
		w.CurrLocation = newLoc
	} else {
		log.Printf("Could not switch to location %v, not a neighbor of %v\n", newLoc.Name, w.CurrLocation.Name)
	}
}

func (w *World) InitializeLocations() {
	// reads through a JSON file to initialize worlds
}
