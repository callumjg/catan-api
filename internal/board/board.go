package board

import (
	"math/rand"
	"strings"
	"time"
)

// The board is represented by a graph of nodes (1-54) and tiles (A-S)
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
	Tiles      []Tile
	Graph      []Node
	RobberTile string
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
		t := NewTile(tid, resources[i], numbers[0])
		if resources[i] != Desert {
			numbers = numbers[1:]
		}
		b.Tiles = append(b.Tiles, t)
	}

	b.initNodes()
	return b
}

func (b *Board) initNodes() {
	b.Graph = []Node{
		{
			Id: 1,
			E:  []int16{4, 5},
		},
		{
			Id: 2,
			E:  []int16{5, 6},
		},
		{
			Id: 3,
			E:  []int16{6, 7},
		},
		{
			Id: 4,
			E:  []int16{1, 8},
		},
		{
			Id: 5,
			E:  []int16{1, 2, 9},
		},
		{
			Id: 6,
			E:  []int16{2, 3, 10},
		},
		{
			Id: 7,
			E:  []int16{3, 11},
		},
		{
			Id: 8,
			E:  []int16{4, 12, 13},
		},
		{
			Id: 9,
			E:  []int16{5, 13, 14},
		},
		{
			Id: 10,
			E:  []int16{6, 14, 15},
		},
		{
			Id: 11,
			E:  []int16{7, 15, 16},
		},
		{
			Id: 12,
			E:  []int16{8, 17},
		},
		{
			Id: 13,
			E:  []int16{8, 9, 18},
		},
		{
			Id: 14,
			E:  []int16{9, 10, 19},
		},
		{
			Id: 15,
			E:  []int16{10, 11, 20},
		},
		{
			Id: 16,
			E:  []int16{11, 21},
		},
		{
			Id: 17,
			E:  []int16{12, 22, 23},
		},
		{
			Id: 18,
			E:  []int16{13, 23, 24},
		},
		{
			Id: 19,
			E:  []int16{14, 24, 25},
		},
		{
			Id: 20,
			E:  []int16{15, 25, 26},
		},
		{
			Id: 21,
			E:  []int16{16, 26, 27},
		},
		{
			Id: 22,
			E:  []int16{17, 28},
		},
		{
			Id: 23,
			E:  []int16{17, 18, 29},
		},
		{
			Id: 24,
			E:  []int16{18, 19, 30},
		},
		{
			Id: 25,
			E:  []int16{19, 20, 31},
		},
		{
			Id: 26,
			E:  []int16{20, 21, 32},
		},
		{
			Id: 27,
			E:  []int16{21, 33},
		},
		{
			Id: 28,
			E:  []int16{22, 34},
		},
		{
			Id: 29,
			E:  []int16{23, 34, 35},
		},
		{
			Id: 30,
			E:  []int16{24, 35, 36},
		},
		{
			Id: 31,
			E:  []int16{25, 36, 37},
		},
		{
			Id: 32,
			E:  []int16{26, 37, 38},
		},
		{
			Id: 33,
			E:  []int16{27, 38},
		},
		{
			Id: 34,
			E:  []int16{28, 29, 39},
		},
		{
			Id: 35,
			E:  []int16{29, 30, 40},
		},
		{
			Id: 36,
			E:  []int16{30, 31, 41},
		},
		{
			Id: 37,
			E:  []int16{31, 32, 42},
		},
		{
			Id: 38,
			E:  []int16{32, 33, 43},
		},
		{
			Id: 39,
			E:  []int16{34, 44},
		},
		{
			Id: 40,
			E:  []int16{35, 44, 45},
		},
		{
			Id: 41,
			E:  []int16{36, 45, 46},
		},
		{
			Id: 42,
			E:  []int16{37, 46, 47},
		},
		{
			Id: 43,
			E:  []int16{38, 47},
		},
		{
			Id: 44,
			E:  []int16{39, 40, 48},
		},
		{
			Id: 45,
			E:  []int16{40, 41, 49},
		},
		{
			Id: 46,
			E:  []int16{41, 42, 50},
		},
		{
			Id: 47,
			E:  []int16{42, 43, 51},
		},
		{
			Id: 48,
			E:  []int16{44, 52},
		},
		{
			Id: 49,
			E:  []int16{45, 52, 53},
		},
		{
			Id: 50,
			E:  []int16{46, 53, 54},
		},
		{
			Id: 51,
			E:  []int16{47, 54},
		},
		{
			Id: 52,
			E:  []int16{48, 49},
		},
		{
			Id: 53,
			E:  []int16{49, 50},
		},
		{
			Id: 54,
			E:  []int16{50, 51},
		},
	}
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
		raw = strings.Replace(raw, "*"+t.Id+"**", t.ShortString(b.RobberTile == t.Id), 1)
	}
	return raw
}
