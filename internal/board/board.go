package board

import (
	"math/rand"
	"time"
)

type Building int16

const (
	Settlement Building = 0
	City                = 1
	Empty               = 2
)

type Board struct {
	Tiles []Tile
	Graph []Node
}

func New() Board {
	var b Board

	// Add desert tile
	b.Tiles = append(b.Tiles, Tile{
		HasRobber: true,
		Resource:  Desert,
	})

	numbers := [18]int16{2, 3, 3, 4, 4, 5, 5, 6, 6, 8, 8, 9, 9, 10, 10, 11, 11, 12}
	resources := [18]Resource{Wood, Wood, Wood, Wood, Wheat, Wheat, Wheat, Wheat, Sheep, Sheep, Sheep, Sheep, Stone, Stone, Stone, Brick, Brick, Brick}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(numbers), func(i, j int) { numbers[i], numbers[j] = numbers[j], numbers[i] })
	rand.Shuffle(len(resources), func(i, j int) { resources[i], resources[j] = resources[j], resources[i] })

	for i := 0; i < 18; i++ {
		t := Tile{
			Number:   numbers[i],
			Resource: resources[i],
		}

		b.Tiles = append(b.Tiles, t)
	}

	// Shuffle tiles
	rand.Shuffle(len(b.Tiles), func(i, j int) { b.Tiles[i], b.Tiles[j] = b.Tiles[j], b.Tiles[i] })

	return b
}
