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

func TestAddingPlayers(t *testing.T) {
	onePlayer := []Player{Player{ID: 0, Name: "Zero"}}
	if _, err := NewGame("One Player", onePlayer); err == nil {
		t.Error("Expected an error when creating a game with one player")
	}

	sevenPlayers := []Player{
		Player{ID: 0, Name: "Zero"},
		Player{ID: 1, Name: "One"},
		Player{ID: 2, Name: "Two"},
		Player{ID: 3, Name: "Three"},
		Player{ID: 4, Name: "Four"},
		Player{ID: 5, Name: "Five"},
		Player{ID: 6, Name: "Six"},
	}
	if _, err := NewGame("Seven Players", sevenPlayers); err == nil {
		t.Error("Expected an error when creating a game with seven players")
	}

	badIDPlayers := []Player{
		Player{ID: 3, Name: "Zero"},
		Player{ID: 1, Name: "One"},
		Player{ID: 2, Name: "Two"},
	}
	if _, err := NewGame("Bad ID", badIDPlayers); err == nil {
		t.Error("Expected an error when creating a game with players with gobbled indices and IDs")
	}

	passingPlayers := []Player{
		Player{ID: 0, Name: "Zero"},
		Player{ID: 1, Name: "One"},
		Player{ID: 2, Name: "Two"},
	}
	if _, err := NewGame("Passing Players", passingPlayers); err != nil {
		t.Error("Expected no errors when adding three players with correct ids, got: ", err)
	}
}

func TestCards(t *testing.T) {
	players := []Player{
		Player{ID: 0, Name: "Zero"},
		Player{ID: 1, Name: "One"},
		Player{ID: 2, Name: "Two"},
	}
	game, err := NewGame("Test Game", players)
	if err != nil {
		t.Error("Unexpected error while building new game:", err)
	}

	expectedCards := []string{"Wild", "Infantry", "Cavalry", "Artillery"}
	for _, card := range(expectedCards) {
		playerCards, ok := game.Cards[card]
		if !ok {
			t.Errorf("Expected card %q to be in the deck", card)
		}

		if len(playerCards) != len(players) {
			t.Errorf("Expected number of players within card %q to be the same as the number of players in the game. Got: %d, want: %d", card, len(playerCards), len(players))
		}

		sum := 0
		for _, count := range(playerCards) {
			sum += count
		}
		if sum != 0 {
			t.Errorf("Expected no player to have any cards at the beginning of the game. Got %d cards dealt", sum)
		}
	}
}

func TestTerritories(t *testing.T) {
	players := []Player{
		Player{ID: 0, Name: "Zero"},
		Player{ID: 1, Name: "One"},
		Player{ID: 2, Name: "Two"},
	}
	game, err := NewGame("Test Game", players)
	if err != nil {
		t.Error("Unexpected error while building new game:", err)
	}

	if len(game.Territories) != 42 {
		t.Errorf("Game does not have 42 territories. Got: %d", len(game.Territories))
	}

	// To avoid testing all 42 territories, we just spot check
	northAfrica, ok := game.Territories["North Africa"]
	if !ok {
		t.Errorf("Game does not have territory %q", "North Africa")
	}
	hasWesternEurope := false
	for _, link := range(northAfrica.Links) {
		if link == "Western Europe" {
			hasWesternEurope = true
			break
		}
	}
	if !hasWesternEurope {
		t.Errorf("Territory %q does not link to %q", "North Africa", "Western Europe")
	}
}
