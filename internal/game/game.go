package game

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/callumjg/catan/internal/board"
	"github.com/callumjg/catan/internal/player"
	"github.com/google/uuid"
)

const MAX_PLAYERS = 4
const MIN_PLAYERS = 2

type Game struct {
	Id               string
	Players          []player.Player
	Board            board.Board
	Host             string
	Started          bool
	Setup            bool
	CurrentPlayerIdx int
}

var Games []Game

func New() Game {
	g := Game{
		Id:               uuid.New().String(),
		Board:            board.New(),
		Started:          false,
		Setup:            false,
		CurrentPlayerIdx: 0,
	}
	return g
}

func (g *Game) AddPlayer(p player.Player) error {
	if len(g.Players) == MAX_PLAYERS {
		return errors.New("Game is full")
	}

	for _, p2 := range g.Players {
		if p2.Id == p.Id {
			return errors.New("Player already in game")
		}
	}

	var excludedColors []player.Color
	hasCol := false

	for _, p2 := range g.Players {
		if p2.Color == p.Color {
			hasCol = true
		}
		excludedColors = append(excludedColors, p2.Color)
	}
	if hasCol {
		p.GetColor(excludedColors...)
	}
	g.Players = append(g.Players, p)

	if len(g.Players) == 1 {
		g.Host = p.Id
	}

	return nil
}

func (g *Game) Start() error {
	if len(g.Players) < MIN_PLAYERS {
		return errors.New(fmt.Sprintf("Need at least %d players to start", MIN_PLAYERS))
	}

	rand.Shuffle(len(g.Players), func(i, j int) { g.Players[i], g.Players[j] = g.Players[j], g.Players[i] })

	return nil
}

func (g *Game) AddSettlement(n int16) error {
	// validate setup & in game
	// is a node
	// not occupied
	// not within 2 spaces
	// under max settlements

	// validate post setup
	// has connecting node
	// player has resources

	return nil
}

func (g *Game) AddRoad(n1 int16, n2 int16) error {
	// validate
	// is an edge
	// not occupied
	// connects to player structure (not blocked)
	// under max roads

	// validate post setup
	// has resources

	return nil
}

func (g *Game) AddCity(n int16) error {
	// validate
	// is a node
	// is occupied with player settlement
	// under max cities

	// validate post setup
	// has resources
	return nil
}

func (g *Game) TurnStart() {

	// Roll dice
	dice := [2]int{
		rand.Intn(6) + 1,
		rand.Intn(6) + 1,
	}
	sum := 0
	for _, n := range dice {
		sum += n
	}

	// Handle 7
	if sum == 7 {
		g.OnSeven()
		return
	}

	// TODO
	// Allocate resources

}

func (g *Game) TurnEnd() {

	// Check score

	isLast := g.CurrentPlayerIdx == len(g.Players)-1
	if isLast {
		g.CurrentPlayerIdx = 0
	} else {
		g.CurrentPlayerIdx++
	}

	g.TurnStart()
}

func (g *Game) OnSeven() {

}
