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
