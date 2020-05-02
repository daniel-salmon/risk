package game

import (
	"testing"
)

func TestArmy(t *testing.T) {
	if Wild != 0 {
		t.Errorf("Wild is not worth 0!. Got: %d", Wild)
	}
	if Infantry != 1 {
		t.Errorf("Infantry is not worth 1!. Got: %d", Infantry)
	}
	if Cavalry != 5 {
		t.Errorf("Cavalry is not worth 5!. Got: %d", Cavalry)
	}
	if Artillery != 10 {
		t.Errorf("Artillery is not worth 10!. Got: %d", Artillery)
	}

	if Wild.String() != "Wild" {
		t.Errorf("Wild string is wrong. Got: %s", Wild.String())
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

	// NOTE: The game contains two wild cards + one card per territory
	if (len(game.Cards.DrawPile) - 2) != len(game.Territories) {
		t.Errorf("Expected the number of Territory cards to equal the number of Territories. Got: %d, want: %d", len(game.Cards.DrawPile)-2, len(game.Territories))
	}

	if len(game.Cards.DiscardPile) != 0 {
		t.Error("Discard pile is non-empty at game initialization")
	}

	if len(game.Cards.OwnedBy) != len(players) {
		t.Errorf("Unexpected difference between the number of players (%d) and the number of people who can own cards (%d)", len(players), len(game.Cards.OwnedBy))
	}

	// Confirm each player has an empty deck of cards
	for _, cards := range game.Cards.OwnedBy {
		if len(cards) != 0 {
			t.Errorf("Some player is starting the game with a non-zero number of cards: %d", len(cards))
		}
	}

	// Aside from the wild cards, there should be an equal distribution of cards between each army type
	// Since there are 42 territories and 3 army types, there should be 14 (=42/3) cards per type
	armyDist := make(map[Army]int)
	for _, card := range game.Cards.DrawPile {
		armyDist[card.ArmyType]++
	}

	if armyDist[Wild] != 2 {
		t.Errorf("Number of Wild cards is not 2. Got: %d, want: %d", armyDist[Wild], 2)
	}
	if armyDist[Infantry] != 14 {
		t.Errorf("Number of Infantry cards is not 14. Got: %d, want: %d", armyDist[Infantry], 14)
	}
	if armyDist[Cavalry] != 14 {
		t.Errorf("Number of Cavalry cards is not 14. Got: %d, want: %d", armyDist[Cavalry], 14)
	}
	if armyDist[Artillery] != 14 {
		t.Errorf("Number of Artillery cards is not 14. Got: %d, want: %d", armyDist[Artillery], 14)
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
	for _, link := range northAfrica.Links {
		if link == "Western Europe" {
			hasWesternEurope = true
			break
		}
	}
	if !hasWesternEurope {
		t.Errorf("Territory %q does not link to %q", "North Africa", "Western Europe")
	}
}
