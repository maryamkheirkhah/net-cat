package server

import (
	"reflect"
	"testing"
	"time"
)

func TestAssignColor(t *testing.T) {
	color1 := AssignColor(1)
	color15 := AssignColor(15)
	if color1 != "\033[0;31m" {
		t.Errorf("AssignColor not assigning prescribed colors correctly")
	} else if color15 != "\033[0;30m" {
		t.Errorf("AssignColor not assigning default color correctly")
	}
}

func TestGetTime(t *testing.T) {
	timeFormat := "2006-01-02 15:04:05"
	currentTime := GetTime()
	_, errTime := time.Parse(timeFormat, currentTime)
	if errTime != nil {
		t.Errorf("GetTime failed to return a time in the correct format, got: %v", currentTime)
	}
}

func TestPortAtoi(t *testing.T) {
	nbrStr := "14578"
	convertedNbr := PortAtoi(nbrStr)
	correctNbr := 14578
	if !reflect.DeepEqual(convertedNbr, correctNbr) {
		t.Errorf("PortAtoi not converting string to integer correctly \ngot: %v \nwanted: %v", convertedNbr, correctNbr)
	}
}

func TestRemoveElement(t *testing.T) {
	NumbTerminal = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	RemoveElement(4)
	correctTerminal := []int{1, 2, 3, 5, 6, 7, 8, 9}
	if !reflect.DeepEqual(NumbTerminal, correctTerminal) {
		t.Errorf("RemoveElement not working /n got: %v \nexpected: %v", NumbTerminal, correctTerminal)
	}
}

func TestFindIndex(t *testing.T) {
	// Test re-assigning integer element
	NumbTerminal = []int{1, 2, 3, 4, 89, 6, 7, 8, 9}
	index := FindIndex(89)
	correctIndex := 4
	if index != correctIndex {
		t.Errorf("FindIndex not finding correct value \ngot: %v \nexpected: %v", index, correctIndex)
	}
}
