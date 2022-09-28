package model

import "testing"

func TestGetName(t *testing.T) {
	l := Location{}
	l.Name = "name"
	name := l.GetName()
	if name != "name" {
		t.Fatal()
	}
}

func TestCheckIfNeightbor(t *testing.T) {
	l := Location{}
	l2 := Location{}

	l.Neighbors = append(l.Neighbors, l2)

	if !l.CheckIfNeighbor(&l2) {
		t.Fatal()
	}
}
