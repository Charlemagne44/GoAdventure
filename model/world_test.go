package model

import (
	"testing"
)

func TestGetCurrLocation(t *testing.T) {
	world := World{Locations: make(map[string]Location, 0),
		CurrLocation: Location{Name: "testing", Neighbors: make([]Location, 0)}}
	currLoc := world.getCurrLocation()
	_ = currLoc
	t.Log(currLoc.getName())
}

//ceasing to test BASIC getters and setters after world

// this is an exception, there are neighbor logic
// checks in this particular setter
func TestSetCurrLocation(t *testing.T) {
	loc1 := Location{Name: "test1", Neighbors: make([]Location, 0)}
	loc2 := Location{Name: "test2", Neighbors: make([]Location, 0)}
	loc1.Neighbors = append(loc1.Neighbors, loc2)
	loc2.Neighbors = append(loc2.Neighbors, loc1)

	world := World{
		Locations:    map[string]Location{},
		CurrLocation: loc1,
	}
	world.Locations[loc1.getName()] = loc1
	world.Locations[loc2.getName()] = loc2

	world.setCurrLocation(&loc2)
	if world.getCurrLocation().getName() != loc2.getName() {
		t.Fatal()
	}
}

func TestInitializeLocations(t *testing.T) {
	filePath := "/Resources/world.json"

	loc1 := Location{Name: "test1", Neighbors: make([]Location, 0)}
	loc2 := Location{Name: "test2", Neighbors: make([]Location, 0)}
	loc1.Neighbors = append(loc1.Neighbors, loc2)
	loc2.Neighbors = append(loc2.Neighbors, loc1)
	world := World{
		Locations:    map[string]Location{},
		CurrLocation: loc1,
	}

	world.initializeLocations(filePath)
}
