package game

import (
	"testing"
)

func TestArmy(t *testing.T) {
	if Infantry != 1 {
		t.Errorf("Infantry is not worth 1!. Got: %d", Infantry)
	}
	if Cavalry != 5 {
		t.Errorf("Cavalry is not worth 5!. Got: %d", Cavalry)
	}
	if Artillery != 10 {
		t.Errorf("Artillery is not worth 10!. Got: %d", Artillery)
	}

	if Infantry.String() != "Infantry" {
		t.Errorf("Infantry string is wrong. Got: %s", Infantry.String())
	}
	if Cavalry.String() != "Cavalry" {
		t.Errorf("Cavalry string is wrong. Got: %s", Cavalry.String())
	}
	if Artillery.String() != "Artillery" {
		t.Errorf("Artillery string is wrong. Got: %s", Artillery.String())
	}
}

func TestBoard(t *testing.T) {
	board, err := NewBoard("Test Board", []Player{Player{ID: 0, Name: "Player Zero"}, Player{ID: 1, Name: "Player One"}})
	if err != nil {
		t.Error("Building new board:", err)
	}
	if len(board.Territories) != 42 {
		t.Errorf("Board does not have 42 territories. Got: %d", len(board.Territories))
	}
}
