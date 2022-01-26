package catan

type Game struct {
	Players []Player
	Board   Board
}

func New() Game {
	var g Game
	return g
}
