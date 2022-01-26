package board

import (
	"fmt"

	"github.com/callumjg/catan/internal/color"
)

type Tile struct {
	Id        string
	Number    int16
	Resource  Resource
	HasRobber bool
}

func (t Tile) String() string {
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

	num := "!!"
	if !t.HasRobber {
		pad := ""
		if t.Number < 10 {
			pad = "0"
		}
		num = fmt.Sprintf("%s%d", pad, t.Number)
	}
	return fmt.Sprintf("%s%s-%s%s", col, t.Resource[0:1], num, color.Reset)

}
