package Model

import "testing"

func TestGetCurrLocation(t *testing.T) {
	world := world{locations: make(map[string]*location, 0),
		currLocation: &location{name: "testing", neighbors: make([]location, 0)}}
	currLoc := world.getCurrLocation()
	if currLoc != world.currLocation {
		t.Fatal()
	}
}

//ceasing to test BASIC getters and setters after world

// this is an exception, there are neighbor logic
// checks in this particular setter
func TestSetCurrLocation(t *testing.T) {
	loc1 := location{name: "test1", neighbors: make([]location, 0)}
	loc2 := location{name: "test2", neighbors: make([]location, 0)}
	loc1.neighbors = append(loc1.neighbors, loc2)
	loc2.neighbors = append(loc2.neighbors, loc1)

	world := world{
		locations:    map[string]*location{},
		currLocation: &loc1,
	}
	world.locations[loc1.getName()] = &loc1
	world.locations[loc2.getName()] = &loc2

	world.setCurrLocation(&loc2)
	if world.getCurrLocation() != &loc2 {
		t.Fatal()
	}
}
