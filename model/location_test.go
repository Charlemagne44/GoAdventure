package model

import (
	"testing"
)

func TestGetName(t *testing.T) {
	loc := Location{Name: "test", Neighbors: make([]Location, 0)}
	loc_name := loc.getName()
	if loc_name != "test" {
		t.Fatal("Wrong name")
	}
}

func TestGetNeighbors(t *testing.T) {
	loc := Location{Name: "test", Neighbors: make([]Location, 0)}
	locTwo := Location{Name: "test2", Neighbors: make([]Location, 0)}
	loc.Neighbors = append(loc.Neighbors, locTwo)

	neighbors := loc.getNeighbors()
	if neighbors[0].getName() != "test2" {
		t.Fatal()
	}
}

func TestCheckIfNeightbor(t *testing.T) {
	loc := Location{Name: "test", Neighbors: make([]Location, 0)}
	locTwo := Location{Name: "test2", Neighbors: make([]Location, 0)}
	loc.Neighbors = append(loc.Neighbors, locTwo)
	locTwo.Neighbors = append(locTwo.Neighbors, loc)

	if !loc.checkIfNeighbor(loc, locTwo) {
		t.Fatal()
	}

}
