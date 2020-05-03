package main

import (
	"github.com/daniel-salmon/risk/game"
)

type Success struct {
	Success bool `json:"success"`
}

// Type Error represents the object the API would return in cases where we want to return
// a descriptive error message without revealing too much about the internals of any errors
type Error struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type NewGame struct {
	Name    string        `json:"name" binding:"required"`
	Players []game.Player `json:"players" binding:"required"`
}

type GameResponse struct {
	Name          string              `json:"name"`
	GoldenCavalry int                 `json:"goldenCavalry"`
	Players       []game.Player       `json:"players"`
	Cards         CardsResponse       `json:"cards"`
	Territories   []TerritoryResponse `json:"territories"`
}

type CardsResponse struct {
	DrawPile    []game.Card         `json:"drawPile"`
	DiscardPile []game.Card         `json:"discardPile"`
	Owned       []CardOwnedResponse `json:"owned"`
}

type CardOwnedResponse struct {
	OwnedBy game.Player `json:"ownedBy"`
	Card    game.Card   `json:"card"`
}

type TerritoryResponse struct {
	Name      string         `json:"name"`
	Continent string         `json:"continent"`
	Links     []string       `json:"links"`
	OwnedBy   *game.Player   `json:"ownedBy"`
	Armies    []ArmyResponse `json:"armies"`
}

type ArmyResponse struct {
	Type  string `json:"type"`
	Value int    `json:"value"`
}
