package Model

import (
	"testing"
)

func TestGetName(t *testing.T) {
	loc := location{name: "test", neighbors: make([]location, 0)}
	loc_name := loc.getName()
	if loc_name != "test" {
		t.Fatal("Wrong name")
	}
}

func TestGetNeighbors(t *testing.T) {
	loc := location{name: "test", neighbors: make([]location, 0)}
	locTwo := location{name: "test2", neighbors: make([]location, 0)}
	loc.neighbors = append(loc.neighbors, locTwo)

	neighbors := loc.getNeighbors()
	if neighbors[0].getName() != "test2" {
		t.Fatal()
	}
}

func TestCheckIfNeightbor(t *testing.T) {
	loc := location{name: "test", neighbors: make([]location, 0)}
	locTwo := location{name: "test2", neighbors: make([]location, 0)}
	loc.neighbors = append(loc.neighbors, locTwo)
	locTwo.neighbors = append(locTwo.neighbors, loc)

	if !loc.checkIfNeighbor(&loc, &locTwo) {
		t.Fatal()
	}

}
