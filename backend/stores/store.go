package stores

import (
	"github.com/daniel-salmon/risk/game"
)

type Store struct {
	game *game.Game
}

func NewStore() (*Store, error) {
	return &Store{}, nil
}

func (s *Store) CreateGame(name string, players []game.Player) (*game.Game, error) {
	g, err := game.NewGame(name, players)
	if err != nil {
		return nil, err
	}
	s.game = g
	return s.game, nil
}

func (s *Store) Close() {
	return
}
