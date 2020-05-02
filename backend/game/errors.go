package game

import (
	"fmt"
)

type IncorrectNumberOfPlayersError struct {
	NumPlayers int
}

func (e *IncorrectNumberOfPlayersError) Error() string {
	return fmt.Sprintf("Incorrect number of players. Want between 2 and 6, got: %d", e.NumPlayers)
}

type PlayerIDMustMatchIndexError struct {
	ID    int
	Index int
}

func (e *PlayerIDMustMatchIndexError) Error() string {
	return fmt.Sprintf("Player with ID %d has index %d which doesn't equal its ID", e.ID, e.Index)
}
