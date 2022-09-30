package model

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Model struct {
	World World
}

func (m *Model) Parse(raw string) (directAct, indirectObj, directObj string) {
	var DA, IO, DO = "", "", ""
	// split into words
	split := strings.Split(raw, " ")

	// fmt.Printf("Raw input split: ")   // DEBUG
	// for n := 0; n < len(split); n++ { // DEBUG
	// fmt.Println(split[n]) // DEBUG
	// } // DEBUG

	// for each word check to see if it matces a DA, DO, or IO via json reading
	for n := 0; n < len(split); n++ {
		word := split[n]
		word = strings.TrimSpace(word)
		// fmt.Printf("DEBUG: word being matched in parser: %v\n", word) // DEBUG

		// check DA
		if b, err := checkType(word, "resources/directAction.json"); b {
			if err != nil {
				fmt.Printf("%v\n", err)
			}
			// fmt.Printf("DEBUG: detected DA of %v\n", word) // DEBUG
			DA = word
		} else if b, err := checkType(word, "resources/directObject.json"); b { // check DO
			if err != nil {
				fmt.Printf("%v\n", err)
			}
			// fmt.Printf("DEBUG: detected DO of %v\n", word) // DEBUG
			DO = word
		} else if b, err := checkType(word, "resources/indirectObject.json"); b { // check IO
			if err != nil {
				fmt.Printf("%v\n", err)
			}
			// fmt.Printf("DEBUG: detected IO of %v\n", word) // DEBUG
			IO = word
		}
	}

	// fmt.Printf("DEBUG: In Parse, final 3. DA: %v, IO: %v, DO: %v\n", DA, IO, DO) // DEBUG

	// if one of these isn't in the statement, returned as ""
	return DA, IO, DO
}

// LEFT OFF LOOKING AT ERROR IN CHECKTYPE -> not returning true as expected

// return command, error
func checkType(word, filename string) (bool, error) {
	// fmt.Printf("word being checked: %v\n", word) // DEBUG
	jsonFile, err := os.Open(filename)
	// fmt.Printf("Error on open: %v\n", err) // DEBUG
	if err != nil {
		return false, err
	}
	defer jsonFile.Close()

	// fmt.Printf("Opened json file \n") // DEBUG

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var potential []string

	err = json.Unmarshal(byteValue, &potential)
	if err != nil {
		return false, err
	}

	// for n := 0; n < len(potential); n++ { // DEBUG
	// fmt.Printf("potential: %v\n", potential[n]) // DEBUG
	// } // DEBUG

	// check to see if it matches any of the potentials
	for _, pot := range potential {
		if pot == word {
			return true, nil
		}
	}

	return false, nil
}

// branches based off of direct action, this function will be LARGE
// due to high amounts of piping elsewhere
func (m *Model) Exec(directAct, indirectObj, directObj string) {
	// fmt.Println("in exec") // DEBUG
	// fmt.Printf("exec direct act: %v\n", directAct) // DEBUG
	if directAct == "move" || directAct == "go" {
		// fmt.Println("In move flags") //DEBUG
		m.moveLoc(directObj)
	}
}

func (m *Model) moveLoc(directObj string) {
	// fmt.Printf("DEBUG: dir obj in moveloc: %v\n", directObj) // DEBUG
	m.World.SetCurrLocation(m.World.Locations[directObj])
}

func (m *Model) Initialize() {
	m.World.InitializeWorld("../resources/world.json")
}
