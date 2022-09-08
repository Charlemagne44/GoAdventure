package Model

import (
	"fmt"
)

type world struct {
	locations    map[string]*location
	currLocation *location
}

func (w *world) getCurrLocation() *location {
	return w.currLocation
}

func (w *world) setCurrLocation(currLoc, newLoc *location) {
	// check neighbors
	n := currLoc.checkIfNeighbor(currLoc, newLoc)
	// if valid move to new loc
	// else reject input
	if n {
		w.currLocation = newLoc
	} else {
		fmt.Printf("Loc %s is not a neighbor of current loc %s\n", newLoc.name, currLoc.name)
	}
}

func (w *world) createLocation(filename string) {
	//TODO read in JSON to create a map of locations
	//modify w.locations
}

func (w *world) initialize() {
	//create locations
	//other steps to initialize the world later on can be added here
}
