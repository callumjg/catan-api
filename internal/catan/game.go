package catan

import "github.com/callumjg/catan/internal/board"

type Game struct {
	Players []Player
	Board   board.Board
}

func New() Game {
	var g Game
	return g
}
