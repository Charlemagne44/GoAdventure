package model

import "testing"

func TestCheckType(t *testing.T) {
	word := "move"
	filename := "../resources/directAction.json"
	answer, err := checkType(word, filename)
	if err != nil {
		t.Fatal(err)
	}
	if !answer {
		t.Fatal()
	}
}

func TestParse(t *testing.T) {
	m := Model{}
	m.Initialize()

	da, io, do := m.Parse("open door with sword")
	t.Logf("%v, %v, %v\n", da, io, do)
	if da != "open" {
		t.Fatal("Did not identify proper action")
	}
	if io != "sword" {
		t.Fatal("Did not identify proper IO")
	}
	if do != "door" {
		t.Fatal("Did not identify proper DO")
	}
}
