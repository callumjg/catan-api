package main

import (
	"fmt"

	"github.com/callumjg/catan/internal/board"
)

func main() {

	b := board.New()
	fmt.Println(b)

	// t := "   A   "
	// r := strings.Replace(t, "A", "bob", 1)
	// fmt.Println(t, r)
}
