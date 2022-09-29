package model

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Model struct {
	World World
}

// for reading through JSON
type DirectAction struct {
	Actions []string
}

type DirectObject struct {
	Objects []string
}

type IndirectObject struct {
	Objects []string
}

func (m *Model) Parse(raw string) (directAct, indirectObj, directObj string) {
	var DA, IO, DO = "", "", ""
	// split into words
	split := strings.Split(raw, " ")

	// for each word check to see if it matces a DA, DO, or IO via json reading
	for _, word := range split {
		// check DA
		if b, err := checkType(word, "../resources/directAction.json"); b {
			if err != nil {
				log.Printf("%v\n", err)
			}
			DA = word
		} else if b, err := checkType(word, "../resources/directObject.json"); b { // check DO
			if err != nil {
				log.Printf("%v\n", err)
			}
			DO = word
		} else if b, err := checkType(word, "../resources/indirectObject.json"); b { // check IO
			if err != nil {
				log.Printf("%v\n", err)
			}
			IO = word
		}
	}

	// if one of these isn't in the statement, returned as ""
	return DA, IO, DO
}

// return command, error
func checkType(word, filename string) (bool, error) {
	jsonFile, err := os.Open(filename)
	if err != nil {
		return false, err
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var potential []string

	err = json.Unmarshal(byteValue, &potential)
	if err != nil {
		return false, err
	}

	// check to see if it matches any of the potentials
	for _, pot := range potential {
		if pot == word {
			return true, nil
		}
	}

	return false, nil
}

func (m *Model) Exec(directAct, indirectObj, directObj string) {

}

func (m *Model) Initialize() {

}
