package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/callumjg/catan/internal/game"
	"github.com/callumjg/catan/internal/player"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	g := game.New()
	p1 := player.New("Callum")
	g.AddPlayer(p1)
	fmt.Print(g)
	// fmt.Println(len(game.Games))
	// t := "   A   "
	// r := strings.Replace(t, "A", "bob", 1)
	// fmt.Println(t, r)
}
