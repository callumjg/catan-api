package board

import (
	"fmt"

	"github.com/callumjg/catan/internal/color"
)

type Tile struct {
	Id       string
	Number   int16
	Resource Resource
	Nodes    [6]int16
}

func NewTile(tid string, r Resource, n int16) Tile {
	t := Tile{
		Id:       tid,
		Resource: r,
	}
	if r != Desert {
		t.Number = n
	}
	switch tid {
	case "A":
		t.Nodes = [6]int16{1, 5, 9, 13, 8, 4}
		break
	case "B":
		t.Nodes = [6]int16{2, 6, 10, 14, 9, 5}
		break
	case "C":
		t.Nodes = [6]int16{3, 7, 11, 15, 10, 6}
		break
	case "D":
		t.Nodes = [6]int16{8, 13, 18, 23, 17, 12}
		break
	case "E":
		t.Nodes = [6]int16{9, 14, 19, 24, 18, 13}
		break
	case "F":
		t.Nodes = [6]int16{10, 15, 20, 25, 19, 14}
		break
	case "G":
		t.Nodes = [6]int16{11, 16, 21, 26, 20, 15}
		break
	case "H":
		t.Nodes = [6]int16{17, 23, 29, 34, 28, 22}
		break
	case "I":
		t.Nodes = [6]int16{18, 24, 30, 35, 29, 23}
		break
	case "J":
		t.Nodes = [6]int16{19, 25, 31, 36, 30, 24}
		break
	case "K":
		t.Nodes = [6]int16{20, 26, 32, 37, 31, 25}
		break
	case "L":
		t.Nodes = [6]int16{21, 27, 33, 38, 32, 26}
		break
	case "M":
		t.Nodes = [6]int16{29, 35, 40, 44, 39, 34}
		break
	case "N":
		t.Nodes = [6]int16{30, 36, 41, 45, 40, 35}
		break
	case "O":
		t.Nodes = [6]int16{31, 37, 42, 46, 41, 36}
		break
	case "P":
		t.Nodes = [6]int16{32, 38, 43, 47, 42, 37}
		break
	case "Q":
		t.Nodes = [6]int16{40, 45, 49, 52, 48, 44}
		break
	case "R":
		t.Nodes = [6]int16{41, 46, 50, 53, 49, 45}
		break
	case "S":
		t.Nodes = [6]int16{42, 47, 51, 54, 50, 46}
		break

	}
	return t
}

func (t Tile) ShortString(hasRobber bool) string {
	// t.Resource
	var col string
	switch t.Resource {
	case Wood:
		col = color.Green
		break
	case Wheat:
		col = color.Yellow
		break
	case Stone:
		col = color.Purple
		break
	case Sheep:
		col = color.Cyan
		break
	case Brick:
		col = color.Red
		break
	case Desert:
		col = color.Gray
		break
	}

	pad := ""
	if t.Number < 10 {
		pad = "0"
	}

	num := fmt.Sprintf("%s%d", pad, t.Number)

	if hasRobber {
		num = "R!"
	}
	return fmt.Sprintf("%s%s-%s%s", col, t.Resource[0:1], num, color.Reset)

}
