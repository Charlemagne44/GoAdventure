package model

import "testing"

func TestSetCurrLocation(t *testing.T) {
	loc1 := Location{Name: "loc1"}
	loc2 := Location{Name: "loc2"}
	loc1.Neighbors = append(loc1.Neighbors, loc2)
	loc2.Neighbors = append(loc2.Neighbors, loc1)
	w := World{
		CurrLocation: loc1,
		Locations:    make(map[string]Location),
	}
	w.Locations[loc1.Name] = loc1
	w.Locations[loc2.Name] = loc2

	t.Logf("Currloc name: %v\n", w.CurrLocation.Name)

	w.SetCurrLocation(loc2)

	t.Logf("Currloc name: %v\n", w.CurrLocation.Name)

	if w.GetCurrLocation().Name != loc2.Name {
		t.Fatal()
	}
}

func TestInitializeLocations(t *testing.T) {
	filename := "../resources/world.json"
	err, locs := InitializeLocations(filename)
	if err != nil {
		t.Fatal(err)
	}
	for _, l := range locs {
		t.Logf("%v\n", l)
	}

}
