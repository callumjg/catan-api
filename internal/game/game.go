package game

import (
	"github.com/callumjg/catan/internal/board"
	"github.com/callumjg/catan/internal/player"
)

type Game struct {
	Players []player.Player
	Board   board.Board
}

func New() Game {
	var g Game
	return g
}
