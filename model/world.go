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

func (w *World) InitializeWorld(worldFile string) {
	// initialize location list
	locs, err := InitializeLocations(worldFile)
	if err != nil {
		log.Printf("Loc list init: %v\n", err)
	}

	// DEBUG show the first loc
	// log.Printf("first loc: %v\n", locs[0].Name)

	// create Locations map out of list
	w.Locations = CreateLocationsMap(locs)

	// initialize current location
	// TODO decide a way to determine first lock
	// if not the first one in JSON
	w.CurrLocation = locs[0]
}

func CreateLocationsMap(locs []Location) map[string]Location {
	locMap := make(map[string]Location, 0)
	for _, loc := range locs {
		locMap[loc.Name] = loc
	}
	return locMap
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

func InitializeLocations(filename string) ([]Location, error) {
	// reads through a JSON file to initialize worlds
	// the reading goes one level up, so the world.json creates the locations contained by world
	jsonFile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	// read opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// initialize Users array
	var locs []Location

	// unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	err = json.Unmarshal(byteValue, &locs)
	if err != nil {
		return nil, err
	}

	return locs, err
}
