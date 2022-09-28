package model

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type World struct {
	// the string that maps to Loc objects will just be the Location.Name
	Locations    map[string]Location
	CurrLocation Location
}

func init() {
	// initialize locations
	InitializeLocations("../resources/world.json")
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

func InitializeLocations(filename string) (error, []string) {
	// reads through a JSON file to initialize worlds
	// the reading goes one level up, so the world.json creates the locations contained by world
	jsonFile, err := os.Open(filename)
	if err != nil {
		return err, nil
	}
	defer jsonFile.Close()

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Users array
	var locs []string

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	err = json.Unmarshal(byteValue, &locs)
	if err != nil {
		return err, nil
	}

	return err, locs
}
