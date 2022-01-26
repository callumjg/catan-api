package board

import (
	"math/rand"
	"strings"
	"time"
)

type Building int16

const (
	Settlement Building = 0
	City                = 1
	Empty               = 2
)

//                01        02        03
//              /    \    /    \    /    \
//           04        05        06        07
//           |    A    |    B    |    C    |
//           08        09        10        11
//         /    \    /    \    /    \    /    \
//      12        13        14        15        16
//      |    D    |    E    |    F    |    G    |
//      17        18        19        20        21
//    /    \    /    \    /    \    /    \    /    \
// 22        23        24        25        26        27
// |    H    |    I    |    J    |    K    |    L    |
// 28        29        30        31        32        33
//    \    /    \    /    \    /    \    /    \    /
//      34        35        36        37        38
//      |    M    |    N    |    O    |    P    |
//      39        40        41        42        43
//         \    /    \    /    \    /    \    /
//           44        45        46        47
//           |    Q    |    R    |    S    |
//           48        49        50        51
//              \    /    \    /    \    /
//                52        53        54

type Board struct {
	Tiles []Tile
	Graph []Node
}

func New() Board {
	var b Board
	rand.Seed(time.Now().UnixNano())

	tids := [19]string{"A", "B", "C", "G", "L", "P", "S", "R", "Q", "M", "H", "D", "E", "F", "K", "O", "N", "I", "J"} // ordered according to perimeter
	resources := [19]Resource{Wood, Wood, Wood, Wood, Wheat, Wheat, Wheat, Wheat, Sheep, Sheep, Sheep, Sheep, Stone, Stone, Stone, Brick, Brick, Brick, Desert}
	numbers := []int16{5, 2, 6, 3, 8, 10, 9, 12, 11, 4, 8, 10, 9, 4, 5, 6, 3, 11} // Numbers are ordered according to box set's A-R

	// Cut numbers at random i
	i := rand.Intn(len(numbers))
	numbers = append(numbers[i:], numbers[0:i]...)

	// Shuffle resources
	rand.Shuffle(len(resources), func(i, j int) { resources[i], resources[j] = resources[j], resources[i] })

	// Create tiles and add to board
	for i, tid := range tids {

		t := Tile{
			Resource:  resources[i],
			Id:        tid,
			HasRobber: true,
		}
		if resources[i] != Desert {
			t.Number = numbers[0]
			numbers = numbers[1:]
			t.HasRobber = false

		}
		b.Tiles = append(b.Tiles, t)
	}

	return b
}

func (b Board) String() string {
	raw := `
                  01        02        03
                /    \    /    \    /    \
             04   A    05   B    06   C    07
             |   *A**  |   *B**  |   *C**  |
             08        09        10        11
           /    \    /    \    /    \    /    \
        12   D    13   E    14   F    15   G    16
        |   *D**  |   *E**  |   *F**  |   *G**  |
        17        18        19        20        21
      /    \    /    \    /    \    /    \    /    \
   22   H    23   I    24   J    25   K    26   L    27
   |   *H**  |   *I**  |   *J**  |   *K**  |   *L**  |
   28        29        30        31        32        33
      \    /    \    /    \    /    \    /    \    /
        34   M    35   N    36   O    37   P    38
        |   *M**  |   *N**  |   *O**  |   *P**  |
        39        40        41        42        43
           \    /    \    /    \    /    \    /
             44   Q    45   R    46   S    47
             |   *Q**  |   *R**  |   *S**  |
             48        49        50        51
                \    /    \    /    \    /
                  52        53        54
`

	for _, t := range b.Tiles {
		raw = strings.Replace(raw, "*"+t.Id+"**", t.String(), 1)
	}
	return raw
}
