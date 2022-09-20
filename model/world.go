package model

import (
	"encoding/json"
	"fmt"
	"os"
)

type World struct {
	Locations    map[string]Location //string should be the location.getName
	CurrLocation Location
}

func init() {
	fmt.Println("DEBUG: world initialized")
}

func (w World) getCurrLocation() Location {
	return w.CurrLocation
}

func (w *World) setCurrLocation(newLoc *Location) {
	// check neighbors
	neighbor := w.getCurrLocation().checkIfNeighbor(w.getCurrLocation(), *newLoc)
	// if valid move to new loc
	// else reject input
	if neighbor {
		w.CurrLocation = *newLoc

	} else {
		// TODO error handling
		fmt.Printf("Loc %s is not a neighbor of current loc %s\n", newLoc.getName(), w.getCurrLocation().getName())
	}
}

func (w World) initializeLocations(filePath string) error {
	jsonFile, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	var dat map[string]*string
	if err := json.Unmarshal([]byte(jsonFile.Name()), &dat); err != nil {
		panic(err)
	}

	return err
}
