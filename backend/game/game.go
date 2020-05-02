package game

import (
	"fmt"
)

type Army int

const (
	Infantry Army = 1
	Cavalry Army = 5
	Artillery Army = 10

	// There is one card for every one of the 42 territories
	// And these are divided evenly among each of the three army types
	// 42 / 3 = 14
	maxCardBucket = 14
	numWildCard = 2
)

func (army Army) String() string {
	if army == 1 {
		return "Infantry"
	}
	if army == 5 {
		return "Cavalry"
	}
	if army == 10 {
		return "Artillery"
	}
	return "Uknown"
}

type Board struct {
	Name string
	Territories map[string](*Territory)
	Cards map[string]map[int]int
	Players []Player
}

type Territory struct {
	Name string
	Continent string
	Links []string
	OwnedBy *Player
}

type Player struct {
	ID int
	Name string
}

func NewBoard(name string, players []Player) (*Board, error) {
	numPlayers := len(players)
	if numPlayers < 2 || numPlayers > 6 {
		return nil, fmt.Errorf("Invalid number of players. Supports 2 to 6 players, got: %d", numPlayers)
	}

	// Initialize cards
	// The cards data structure maps the card type to a player's id
	// "Underneath" the player's id we maintain the count of cards held by that player
	cards := make(map[string]map[int]int)
	cards["Wild"] = make(map[int]int)
	cards["Infantry"] = make(map[int]int)
	cards["Cavalry"] = make(map[int]int)
	cards["Artillery"] = make(map[int]int)

	for i, p := range(players) {
		if p.ID != i {
			return nil, fmt.Errorf("Player %d's ID does not match its index. got: ID = %d", i, p.ID)
		}
		cards["Wild"][i] = 0
		cards["Infantry"][i] = 0
		cards["Cavalry"][i] = 0
		cards["Artillery"][i] = 0
	}

	// Initialize all 42 territories
	territories := make(map[string](*Territory))

	// North America
	territories["Alaska"] = &Territory{
		Name: "Alaska",
		Continent: "North America",
		Links: []string{"Alberta", "Northwest Territory", "Kamchatka"},
		OwnedBy: nil,
	}
	territories["Alberta"] = &Territory{
		Name: "Alberta",
		Continent: "North America",
		Links: []string{"Alaska", "Northwest Territory", "Ontario", "Western United States"},
		OwnedBy: nil,
	}
	territories["Western United States"] = &Territory{
		Name: "Western United States",
		Continent: "North America",
		Links: []string{"Alberta", "Ontario", "Eastern United States", "Central America"},
		OwnedBy: nil,
	}
	territories["Central America"] = &Territory{
		Name: "Central America",
		Continent: "North America",
		Links: []string{"Western United States", "Eastern United States", "Venezuela"},
		OwnedBy: nil,
	}
	territories["Northwest Territory"] = &Territory{
		Name: "Northwest Territory",
		Continent: "North America",
		Links: []string{"Alaska", "Alberta", "Ontario", "Greenland"},
		OwnedBy: nil,
	}
	territories["Ontario"] = &Territory{
		Name: "Ontario",
		Continent: "North America",
		Links: []string{"Northwest Territory", "Alberta", "Western United States", "Eastern United States", "Quebec", "Greenland"},
		OwnedBy: nil,
	}
	territories["Eastern United States"] = &Territory{
		Name: "Eastern United States",
		Continent: "North America",
		Links: []string{"Ontario", "Western United States", "Central America", "Quebec"},
		OwnedBy: nil,
	}
	territories["Greenland"] = &Territory{
		Name: "Greenland",
		Continent: "North America",
		Links: []string{"Northwest Territory", "Ontario", "Quebec", "Iceland"},
		OwnedBy: nil,
	}
	territories["Quebec"] = &Territory{
		Name: "Quebec",
		Continent: "North America",
		Links: []string{"Greenland", "Ontario", "Eastern United States"},
		OwnedBy: nil,
	}

	// South America
	territories["Venezuela"] = &Territory{
		Name: "Venezuela",
		Continent: "South America",
		Links: []string{"Central America", "Peru", "Brazil"},
		OwnedBy: nil,
	}
	territories["Peru"] = &Territory{
		Name: "Peru",
		Continent: "South America",
		Links: []string{"Venezuela", "Argentina", "Brazil"},
		OwnedBy: nil,
	}
	territories["Argentina"] = &Territory{
		Name: "Argentina",
		Continent: "South America",
		Links: []string{"Peru", "Brazil"},
		OwnedBy: nil,
	}
	territories["Brazil"] = &Territory{
		Name: "Brazil",
		Continent: "South America",
		Links: []string{"Venezuela", "Peru", "Argentina", "North Africa"},
		OwnedBy: nil,
	}

	// Africa
	territories["North Africa"] = &Territory{
		Name: "North Africa",
		Continent: "Africa",
		Links: []string{"Brazil", "Congo", "East Africa", "Egypt", "Southern Europe", "Western Europe"},
		OwnedBy: nil,
	}
	territories["Egypt"] = &Territory{
		Name: "Egypt",
		Continent: "Africa",
		Links: []string{"Southern Europe", "North Africa", "East Africa", "Middle East"},
		OwnedBy: nil,
	}
	territories["Congo"] = &Territory{
		Name: "Congo",
		Continent: "Africa",
		Links: []string{"North Africa", "South Africa", "East Africa"},
		OwnedBy: nil,
	}
	territories["South Africa"] = &Territory{
		Name: "South Africa",
		Continent: "Africa",
		Links: []string{"Congo", "East Africa", "Madagascar"},
		OwnedBy: nil,
	}
	territories["Madagascar"] = &Territory{
		Name: "Madagascar",
		Continent: "Africa",
		Links: []string{"South Africa", "East Africa"},
		OwnedBy: nil,
	}
	territories["East Africa"] = &Territory{
		Name: "East Africa",
		Continent: "Africa",
		Links: []string{"Egypt", "North Africa", "Congo", "South Africa", "Madagascar", "Middle East"},
		OwnedBy: nil,
	}

	// Europe
	territories["Iceland"] = &Territory{
		Name: "Iceland",
		Continent: "Europe",
		Links: []string{"Greenland", "Great Britain", "Scandinavia"},
		OwnedBy: nil,
	}
	territories["Great Britain"] = &Territory{
		Name: "Great Britain",
		Continent: "Europe",
		Links: []string{"Iceland", "Western Europe", "Scandinavia", "Northern Europe"},
		OwnedBy: nil,
	}
	territories["Western Europe"] = &Territory{
		Name: "Western Europe",
		Continent: "Europe",
		Links: []string{"Great Britain", "North Africa", "Southern Europe", "Northern Europe"},
		OwnedBy: nil,
	}
	territories["Southern Europe"] = &Territory{
		Name: "Southern Europe",
		Continent: "Europe",
		Links: []string{"Western Europe", "North Africa", "Egypt", "Middle East", "Ukraine", "Northern Europe"},
		OwnedBy: nil,
	}
	territories["Northern Europe"] = &Territory{
		Name: "Northern Europe",
		Continent: "Europe",
		Links: []string{"Southern Europe", "Western Europe", "Great Britain", "Scandinavia", "Ukraine"},
		OwnedBy: nil,
	}
	territories["Scandinavia"] = &Territory{
		Name: "Scandinavia",
		Continent: "Europe",
		Links: []string{"Iceland", "Great Britain", "Northern Europe", "Ukraine"},
		OwnedBy: nil,
	}
	territories["Ukraine"] = &Territory{
		Name: "Ukraine",
		Continent: "Europe",
		Links: []string{"Scandinavia", "Northern Europe", "Southern Europe", "Middle East", "Afghanistan", "Ural"},
		OwnedBy: nil,
	}

	// Asia
	territories["Ural"] = &Territory{
		Name: "Ural",
		Continent: "Asia",
		Links: []string{"Ukraine", "Afghanistan", "China", "Siberia"},
		OwnedBy: nil,
	}
	territories["Afghanistan"] = &Territory{
		Name: "Afghanistan",
		Continent: "Asia",
		Links: []string{"Ukraine", "Middle East", "India", "China", "Ural"},
		OwnedBy: nil,
	}
	territories["Middle East"] = &Territory{
		Name: "Middle East",
		Continent: "Asia",
		Links: []string{"Ukraine", "Southern Europe", "Egypt", "East Africa", "India", "Afghanistan"},
		OwnedBy: nil,
	}
	territories["Siberia"] = &Territory{
		Name: "Siberia",
		Continent: "Asia",
		Links: []string{"Ural", "China", "Mongolia", "Irkutsk", "Yakutsk"},
		OwnedBy: nil,
	}
	territories["China"] = &Territory{
		Name: "China",
		Continent: "Asia",
		Links: []string{"Siberia", "Ural", "Afghanistan", "India", "Siam", "Mongolia"},
		OwnedBy: nil,
	}
	territories["India"] = &Territory{
		Name: "India",
		Continent: "Asia",
		Links: []string{"China", "Afghanistan", "Middle East", "Siam"},
		OwnedBy: nil,
	}
	territories["Yakutsk"] = &Territory{
		Name: "Yakutsk",
		Continent: "Asia",
		Links: []string{"Kamchatka", "Irkutsk", "Siberia"},
		OwnedBy: nil,
	}
	territories["Irkutsk"] = &Territory{
		Name: "Irkutsk",
		Continent: "Asia",
		Links: []string{"Yakutsk", "Siberia", "Mongolia", "Japan", "Kamchatka"},
		OwnedBy: nil,
	}
	territories["Mongolia"] = &Territory{
		Name: "Mongolia",
		Continent: "Asia",
		Links: []string{"Irkutsk", "Siberia", "China", "Japan"},
		OwnedBy: nil,
	}
	territories["Kamchatka"] = &Territory{
		Name: "Kamchatka",
		Continent: "Asia",
		Links: []string{"Yakutsk", "Irkutsk", "Japan", "Alaska"},
		OwnedBy: nil,
	}
	territories["Japan"] = &Territory{
		Name: "Japan",
		Continent: "Asia",
		Links: []string{"Kamchatka", "Irkutsk", "Mongolia"},
		OwnedBy: nil,
	}
	territories["Siam"] = &Territory{
		Name: "Siam",
		Continent: "Asia",
		Links: []string{"China", "India", "Indonesia"},
		OwnedBy: nil,
	}

	// Australia
	territories["Indonesia"] = &Territory{
		Name: "Indonesia",
		Continent: "Australia",
		Links: []string{"Siam", "New Guinea", "Western Australia"},
		OwnedBy: nil,
	}
	territories["New Guinea"] = &Territory{
		Name: "New Guinea",
		Continent: "Australia",
		Links: []string{"Indonesia", "Western Australia", "Eastern Australia"},
		OwnedBy: nil,
	}
	territories["Western Australia"] = &Territory{
		Name: "Western Australia",
		Continent: "Australia",
		Links: []string{"Indonesia", "New Guinea", "Eastern Australia"},
		OwnedBy: nil,
	}
	territories["Eastern Australia"] = &Territory{
		Name: "Eastern Australia",
		Continent: "Australia",
		Links: []string{"Western Australia", "New Guinea"},
		OwnedBy: nil,
	}

	b := Board{
		Name: name,
		Territories: territories,
		Cards: cards,
		Players: players,
	}

	return &b, nil
}
