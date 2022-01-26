package player

import (
	"math/rand"

	"github.com/google/uuid"
)

type Color string

const (
	Blue   Color = "Blue"
	Red          = "Red"
	Orange       = "Orange"
	White        = "White"
	Green        = "Green"
)

func Colors() [5]Color {
	return [5]Color{Blue, Red, Orange, White, Green}
}

type Player struct {
	Id    string
	Name  string
	Color Color
}

func New(name string) Player {
	return Player{
		Id:   uuid.New().String(),
		Name: name,
	}
}

func (p *Player) GetColor(excluded ...Color) {
	var available []Color

	for _, c := range Colors() {
		found := false
		for _, c2 := range excluded {
			if c2 == c {
				found = false
				break
			}
		}
		if !found {
			available = append(available, c)
		}
	}

	i := rand.Intn(len(available))
	p.Color = available[i]
}
