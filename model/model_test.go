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

	da, _, _ := m.Parse("move east")
	if da != "move" {
		t.Fatal("Did not identify proper action")
	}
}
