package stores

import (
	"fmt"
	"github.com/daniel-salmon/risk/game"
)

type Store struct {
	game *game.Game
}

func NewStore() (*Store, error) {
	return &Store{}, nil
}

func (s *Store) CreateGame(name string, players []game.Player) error {
	g, err := game.NewGame(name, players)
	if err != nil {
		return err
	}
	s.game = g
	fmt.Println(s.game)
	return nil
}

func (s *Store) Close() {
	return
}
