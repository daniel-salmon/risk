package game

type Army int

const (
	Wild      Army = 0
	Infantry  Army = 1
	Cavalry   Army = 5
	Artillery Army = 10
)

func (army Army) String() string {
	if army == 0 {
		return "Wild"
	}
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

type Game struct {
	Name          string                  `json:"name"`
	GoldenCavalry int                     `json:"goldenCavalry"`
	Territories   map[string](*Territory) `json:"territories"`
	Cards         *Cards                  `json:"cards"`
	Players       []Player                `json:"players"`
}

type Territory struct {
	Name      string       `json:"name"`
	Continent string       `json:"continent"`
	Links     []string     `json:"links"`
	OwnedBy   *Player      `json:"ownedBy"`
	Armies    map[Army]int `json:"armies"`
}

type Cards struct {
	DrawPile    []Card         `json:"drawPile"`
	DiscardPile []Card         `json:"discardPile"`
	OwnedBy     map[int][]Card `json:"ownedBy"`
}

type Card struct {
	Territory string `json:"territory"`
	ArmyType  Army   `json:"armyType"`
}

type Player struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func NewGame(name string, players []Player) (*Game, error) {
	// We are playing "World Domination Risk" which requires 3-6 players
	if len(players) < 3 || len(players) > 6 {
		return nil, &IncorrectNumberOfPlayersError{NumPlayers: len(players)}
	}

	// Initialize the draw pile of cards
	// We'll append to this as we initialize the territories, but there are two wild cards in the deck
	var card Card
	drawPile := []Card{
		Card{Territory: "Myjäss", ArmyType: Wild},
		Card{Territory: "Myjäss", ArmyType: Wild},
	}

	// Initialize the empty discard pile
	discardPile := []Card{}

	// Initialize the empty owned by pile
	// When a player obtains a card that that card will be removed from the draw pile
	// and added to this map of player ids to the cards they own
	ownedBy := make(map[int][]Card)
	for i, p := range players {
		if p.ID != i {
			return nil, &PlayerIDMustMatchIndexError{ID: p.ID, Index: i}
		}
		ownedBy[i] = []Card{}
	}

	// Initialize all 42 territories
	territories := make(map[string](*Territory))

	// North America
	card = Card{Territory: "Alaska", ArmyType: Infantry}
	drawPile = append(drawPile, card)
	territories["Alaska"] = &Territory{
		Name:      "Alaska",
		Continent: "North America",
		Links:     []string{"Alberta", "Northwest Territory", "Kamchatka"},
		OwnedBy:   nil,
		Armies:    map[Army]int{Infantry: 0, Cavalry: 0, Artillery: 0},
	}
	card = Card{Territory: "Alberta", ArmyType: Cavalry}
	drawPile = append(drawPile, card)
	territories["Alberta"] = &Territory{
		Name:      "Alberta",
		Continent: "North America",
		Links:     []string{"Alaska", "Northwest Territory", "Ontario", "Western United States"},
		OwnedBy:   nil,
		Armies:    map[Army]int{Infantry: 0, Cavalry: 0, Artillery: 0},
	}
	card = Card{Territory: "Western United States", ArmyType: Artillery}
	drawPile = append(drawPile, card)
	territories["Western United States"] = &Territory{
		Name:      "Western United States",
		Continent: "North America",
		Links:     []string{"Alberta", "Ontario", "Eastern United States", "Central America"},
		OwnedBy:   nil,
		Armies:    map[Army]int{Infantry: 0, Cavalry: 0, Artillery: 0},
	}
	card = Card{Territory: "Central America", ArmyType: Infantry}
	drawPile = append(drawPile, card)
	territories["Central America"] = &Territory{
		Name:      "Central America",
		Continent: "North America",
		Links:     []string{"Western United States", "Eastern United States", "Venezuela"},
		OwnedBy:   nil,
		Armies:    map[Army]int{Infantry: 0, Cavalry: 0, Artillery: 0},
	}
	card = Card{Territory: "Northwest Territory", ArmyType: Cavalry}
	drawPile = append(drawPile, card)
	territories["Northwest Territory"] = &Territory{
		Name:      "Northwest Territory",
		Continent: "North America",
		Links:     []string{"Alaska", "Alberta", "Ontario", "Greenland"},
		OwnedBy:   nil,
		Armies:    map[Army]int{Infantry: 0, Cavalry: 0, Artillery: 0},
	}
	card = Card{Territory: "Ontario", ArmyType: Artillery}
	drawPile = append(drawPile, card)
	territories["Ontario"] = &Territory{
		Name:      "Ontario",
		Continent: "North America",
		Links:     []string{"Northwest Territory", "Alberta", "Western United States", "Eastern United States", "Quebec", "Greenland"},
		OwnedBy:   nil,
		Armies:    map[Army]int{Infantry: 0, Cavalry: 0, Artillery: 0},
	}
	card = Card{Territory: "Eastern United States", ArmyType: Infantry}
	drawPile = append(drawPile, card)
	territories["Eastern United States"] = &Territory{
		Name:      "Eastern United States",
		Continent: "North America",
		Links:     []string{"Ontario", "Western United States", "Central America", "Quebec"},
		OwnedBy:   nil,
		Armies:    map[Army]int{Infantry: 0, Cavalry: 0, Artillery: 0},
	}
	card = Card{Territory: "Greenland", ArmyType: Cavalry}
	drawPile = append(drawPile, card)
	territories["Greenland"] = &Territory{
		Name:      "Greenland",
		Continent: "North America",
		Links:     []string{"Northwest Territory", "Ontario", "Quebec", "Iceland"},
		OwnedBy:   nil,
		Armies:    map[Army]int{Infantry: 0, Cavalry: 0, Artillery: 0},
	}
	card = Card{Territory: "Quebec", ArmyType: Artillery}
	drawPile = append(drawPile, card)
	territories["Quebec"] = &Territory{
		Name:      "Quebec",
		Continent: "North America",
		Links:     []string{"Greenland", "Ontario", "Eastern United States"},
		OwnedBy:   nil,
		Armies:    map[Army]int{Infantry: 0, Cavalry: 0, Artillery: 0},
	}

	// South America
	card = Card{Territory: "Venezuela", ArmyType: Infantry}
	drawPile = append(drawPile, card)
	territories["Venezuela"] = &Territory{
		Name:      "Venezuela",
		Continent: "South America",
		Links:     []string{"Central America", "Peru", "Brazil"},
		OwnedBy:   nil,
		Armies:    map[Army]int{Infantry: 0, Cavalry: 0, Artillery: 0},
	}
	card = Card{Territory: "Peru", ArmyType: Cavalry}
	drawPile = append(drawPile, card)
	territories["Peru"] = &Territory{
		Name:      "Peru",
		Continent: "South America",
		Links:     []string{"Venezuela", "Argentina", "Brazil"},
		OwnedBy:   nil,
		Armies:    map[Army]int{Infantry: 0, Cavalry: 0, Artillery: 0},
	}
	card = Card{Territory: "Argentina", ArmyType: Artillery}
	drawPile = append(drawPile, card)
	territories["Argentina"] = &Territory{
		Name:      "Argentina",
		Continent: "South America",
		Links:     []string{"Peru", "Brazil"},
		OwnedBy:   nil,
		Armies:    map[Army]int{Infantry: 0, Cavalry: 0, Artillery: 0},
	}
	card = Card{Territory: "Brazil", ArmyType: Infantry}
	drawPile = append(drawPile, card)
	territories["Brazil"] = &Territory{
		Name:      "Brazil",
		Continent: "South America",
		Links:     []string{"Venezuela", "Peru", "Argentina", "North Africa"},
		OwnedBy:   nil,
		Armies:    map[Army]int{Infantry: 0, Cavalry: 0, Artillery: 0},
	}

	// Africa
	card = Card{Territory: "North Africa", ArmyType: Cavalry}
	drawPile = append(drawPile, card)
	territories["North Africa"] = &Territory{
		Name:      "North Africa",
		Continent: "Africa",
		Links:     []string{"Brazil", "Congo", "East Africa", "Egypt", "Southern Europe", "Western Europe"},
		OwnedBy:   nil,
		Armies:    map[Army]int{Infantry: 0, Cavalry: 0, Artillery: 0},
	}
	card = Card{Territory: "Egypt", ArmyType: Artillery}
	drawPile = append(drawPile, card)
	territories["Egypt"] = &Territory{
		Name:      "Egypt",
		Continent: "Africa",
		Links:     []string{"Southern Europe", "North Africa", "East Africa", "Middle East"},
		OwnedBy:   nil,
		Armies:    map[Army]int{Infantry: 0, Cavalry: 0, Artillery: 0},
	}
	card = Card{Territory: "Congo", ArmyType: Infantry}
	drawPile = append(drawPile, card)
	territories["Congo"] = &Territory{
		Name:      "Congo",
		Continent: "Africa",
		Links:     []string{"North Africa", "South Africa", "East Africa"},
		OwnedBy:   nil,
		Armies:    map[Army]int{Infantry: 0, Cavalry: 0, Artillery: 0},
	}
	card = Card{Territory: "South Africa", ArmyType: Cavalry}
	drawPile = append(drawPile, card)
	territories["South Africa"] = &Territory{
		Name:      "South Africa",
		Continent: "Africa",
		Links:     []string{"Congo", "East Africa", "Madagascar"},
		OwnedBy:   nil,
		Armies:    map[Army]int{Infantry: 0, Cavalry: 0, Artillery: 0},
	}
	card = Card{Territory: "Madagascar", ArmyType: Artillery}
	drawPile = append(drawPile, card)
	territories["Madagascar"] = &Territory{
		Name:      "Madagascar",
		Continent: "Africa",
		Links:     []string{"South Africa", "East Africa"},
		OwnedBy:   nil,
		Armies:    map[Army]int{Infantry: 0, Cavalry: 0, Artillery: 0},
	}
	card = Card{Territory: "East Africa", ArmyType: Infantry}
	drawPile = append(drawPile, card)
	territories["East Africa"] = &Territory{
		Name:      "East Africa",
		Continent: "Africa",
		Links:     []string{"Egypt", "North Africa", "Congo", "South Africa", "Madagascar", "Middle East"},
		OwnedBy:   nil,
		Armies:    map[Army]int{Infantry: 0, Cavalry: 0, Artillery: 0},
	}

	// Europe
	card = Card{Territory: "Iceland", ArmyType: Cavalry}
	drawPile = append(drawPile, card)
	territories["Iceland"] = &Territory{
		Name:      "Iceland",
		Continent: "Europe",
		Links:     []string{"Greenland", "Great Britain", "Scandinavia"},
		OwnedBy:   nil,
		Armies:    map[Army]int{Infantry: 0, Cavalry: 0, Artillery: 0},
	}
	card = Card{Territory: "Great Britain", ArmyType: Artillery}
	drawPile = append(drawPile, card)
	territories["Great Britain"] = &Territory{
		Name:      "Great Britain",
		Continent: "Europe",
		Links:     []string{"Iceland", "Western Europe", "Scandinavia", "Northern Europe"},
		OwnedBy:   nil,
		Armies:    map[Army]int{Infantry: 0, Cavalry: 0, Artillery: 0},
	}
	card = Card{Territory: "Western Europe", ArmyType: Infantry}
	drawPile = append(drawPile, card)
	territories["Western Europe"] = &Territory{
		Name:      "Western Europe",
		Continent: "Europe",
		Links:     []string{"Great Britain", "North Africa", "Southern Europe", "Northern Europe"},
		OwnedBy:   nil,
		Armies:    map[Army]int{Infantry: 0, Cavalry: 0, Artillery: 0},
	}
	card = Card{Territory: "Southern Europe", ArmyType: Cavalry}
	drawPile = append(drawPile, card)
	territories["Southern Europe"] = &Territory{
		Name:      "Southern Europe",
		Continent: "Europe",
		Links:     []string{"Western Europe", "North Africa", "Egypt", "Middle East", "Ukraine", "Northern Europe"},
		OwnedBy:   nil,
		Armies:    map[Army]int{Infantry: 0, Cavalry: 0, Artillery: 0},
	}
	card = Card{Territory: "Northern Europe", ArmyType: Artillery}
	drawPile = append(drawPile, card)
	territories["Northern Europe"] = &Territory{
		Name:      "Northern Europe",
		Continent: "Europe",
		Links:     []string{"Southern Europe", "Western Europe", "Great Britain", "Scandinavia", "Ukraine"},
		OwnedBy:   nil,
		Armies:    map[Army]int{Infantry: 0, Cavalry: 0, Artillery: 0},
	}
	card = Card{Territory: "Scandinavia", ArmyType: Infantry}
	drawPile = append(drawPile, card)
	territories["Scandinavia"] = &Territory{
		Name:      "Scandinavia",
		Continent: "Europe",
		Links:     []string{"Iceland", "Great Britain", "Northern Europe", "Ukraine"},
		OwnedBy:   nil,
		Armies:    map[Army]int{Infantry: 0, Cavalry: 0, Artillery: 0},
	}
	card = Card{Territory: "Ukraine", ArmyType: Cavalry}
	drawPile = append(drawPile, card)
	territories["Ukraine"] = &Territory{
		Name:      "Ukraine",
		Continent: "Europe",
		Links:     []string{"Scandinavia", "Northern Europe", "Southern Europe", "Middle East", "Afghanistan", "Ural"},
		OwnedBy:   nil,
		Armies:    map[Army]int{Infantry: 0, Cavalry: 0, Artillery: 0},
	}

	// Asia
	card = Card{Territory: "Ural", ArmyType: Artillery}
	drawPile = append(drawPile, card)
	territories["Ural"] = &Territory{
		Name:      "Ural",
		Continent: "Asia",
		Links:     []string{"Ukraine", "Afghanistan", "China", "Siberia"},
		OwnedBy:   nil,
		Armies:    map[Army]int{Infantry: 0, Cavalry: 0, Artillery: 0},
	}
	card = Card{Territory: "Afghanistan", ArmyType: Infantry}
	drawPile = append(drawPile, card)
	territories["Afghanistan"] = &Territory{
		Name:      "Afghanistan",
		Continent: "Asia",
		Links:     []string{"Ukraine", "Middle East", "India", "China", "Ural"},
		OwnedBy:   nil,
		Armies:    map[Army]int{Infantry: 0, Cavalry: 0, Artillery: 0},
	}
	card = Card{Territory: "Middle East", ArmyType: Cavalry}
	drawPile = append(drawPile, card)
	territories["Middle East"] = &Territory{
		Name:      "Middle East",
		Continent: "Asia",
		Links:     []string{"Ukraine", "Southern Europe", "Egypt", "East Africa", "India", "Afghanistan"},
		OwnedBy:   nil,
		Armies:    map[Army]int{Infantry: 0, Cavalry: 0, Artillery: 0},
	}
	card = Card{Territory: "Siberia", ArmyType: Artillery}
	drawPile = append(drawPile, card)
	territories["Siberia"] = &Territory{
		Name:      "Siberia",
		Continent: "Asia",
		Links:     []string{"Ural", "China", "Mongolia", "Irkutsk", "Yakutsk"},
		OwnedBy:   nil,
		Armies:    map[Army]int{Infantry: 0, Cavalry: 0, Artillery: 0},
	}
	card = Card{Territory: "China", ArmyType: Infantry}
	drawPile = append(drawPile, card)
	territories["China"] = &Territory{
		Name:      "China",
		Continent: "Asia",
		Links:     []string{"Siberia", "Ural", "Afghanistan", "India", "Siam", "Mongolia"},
		OwnedBy:   nil,
		Armies:    map[Army]int{Infantry: 0, Cavalry: 0, Artillery: 0},
	}
	card = Card{Territory: "India", ArmyType: Cavalry}
	drawPile = append(drawPile, card)
	territories["India"] = &Territory{
		Name:      "India",
		Continent: "Asia",
		Links:     []string{"China", "Afghanistan", "Middle East", "Siam"},
		OwnedBy:   nil,
		Armies:    map[Army]int{Infantry: 0, Cavalry: 0, Artillery: 0},
	}
	card = Card{Territory: "Yakutsk", ArmyType: Artillery}
	drawPile = append(drawPile, card)
	territories["Yakutsk"] = &Territory{
		Name:      "Yakutsk",
		Continent: "Asia",
		Links:     []string{"Kamchatka", "Irkutsk", "Siberia"},
		OwnedBy:   nil,
		Armies:    map[Army]int{Infantry: 0, Cavalry: 0, Artillery: 0},
	}
	card = Card{Territory: "Irkutsk", ArmyType: Infantry}
	drawPile = append(drawPile, card)
	territories["Irkutsk"] = &Territory{
		Name:      "Irkutsk",
		Continent: "Asia",
		Links:     []string{"Yakutsk", "Siberia", "Mongolia", "Japan", "Kamchatka"},
		OwnedBy:   nil,
		Armies:    map[Army]int{Infantry: 0, Cavalry: 0, Artillery: 0},
	}
	card = Card{Territory: "Mongolia", ArmyType: Cavalry}
	drawPile = append(drawPile, card)
	territories["Mongolia"] = &Territory{
		Name:      "Mongolia",
		Continent: "Asia",
		Links:     []string{"Irkutsk", "Siberia", "China", "Japan"},
		OwnedBy:   nil,
		Armies:    map[Army]int{Infantry: 0, Cavalry: 0, Artillery: 0},
	}
	card = Card{Territory: "Kamchatka", ArmyType: Artillery}
	drawPile = append(drawPile, card)
	territories["Kamchatka"] = &Territory{
		Name:      "Kamchatka",
		Continent: "Asia",
		Links:     []string{"Yakutsk", "Irkutsk", "Japan", "Alaska"},
		OwnedBy:   nil,
		Armies:    map[Army]int{Infantry: 0, Cavalry: 0, Artillery: 0},
	}
	card = Card{Territory: "Japan", ArmyType: Infantry}
	drawPile = append(drawPile, card)
	territories["Japan"] = &Territory{
		Name:      "Japan",
		Continent: "Asia",
		Links:     []string{"Kamchatka", "Irkutsk", "Mongolia"},
		OwnedBy:   nil,
		Armies:    map[Army]int{Infantry: 0, Cavalry: 0, Artillery: 0},
	}
	card = Card{Territory: "Siam", ArmyType: Cavalry}
	drawPile = append(drawPile, card)
	territories["Siam"] = &Territory{
		Name:      "Siam",
		Continent: "Asia",
		Links:     []string{"China", "India", "Indonesia"},
		OwnedBy:   nil,
		Armies:    map[Army]int{Infantry: 0, Cavalry: 0, Artillery: 0},
	}

	// Australia
	card = Card{Territory: "Indonesia", ArmyType: Artillery}
	drawPile = append(drawPile, card)
	territories["Indonesia"] = &Territory{
		Name:      "Indonesia",
		Continent: "Australia",
		Links:     []string{"Siam", "New Guinea", "Western Australia"},
		OwnedBy:   nil,
		Armies:    map[Army]int{Infantry: 0, Cavalry: 0, Artillery: 0},
	}
	card = Card{Territory: "New Guinea", ArmyType: Infantry}
	drawPile = append(drawPile, card)
	territories["New Guinea"] = &Territory{
		Name:      "New Guinea",
		Continent: "Australia",
		Links:     []string{"Indonesia", "Western Australia", "Eastern Australia"},
		OwnedBy:   nil,
		Armies:    map[Army]int{Infantry: 0, Cavalry: 0, Artillery: 0},
	}
	card = Card{Territory: "Western Australia", ArmyType: Cavalry}
	drawPile = append(drawPile, card)
	territories["Western Australia"] = &Territory{
		Name:      "Western Australia",
		Continent: "Australia",
		Links:     []string{"Indonesia", "New Guinea", "Eastern Australia"},
		OwnedBy:   nil,
		Armies:    map[Army]int{Infantry: 0, Cavalry: 0, Artillery: 0},
	}
	card = Card{Territory: "Eastern Australia", ArmyType: Artillery}
	drawPile = append(drawPile, card)
	territories["Eastern Australia"] = &Territory{
		Name:      "Eastern Australia",
		Continent: "Australia",
		Links:     []string{"Western Australia", "New Guinea"},
		OwnedBy:   nil,
		Armies:    map[Army]int{Infantry: 0, Cavalry: 0, Artillery: 0},
	}

	cards := &Cards{
		DrawPile:    drawPile,
		DiscardPile: discardPile,
		OwnedBy:     ownedBy,
	}

	g := Game{
		Name:          name,
		GoldenCavalry: 4,
		Territories:   territories,
		Cards:         cards,
		Players:       players,
	}

	return &g, nil
}
