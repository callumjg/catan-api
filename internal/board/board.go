package board

import (
	"math/rand"
	"strings"
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

type Structure int16

const (
	Settlement Structure = 0
	City                 = 1
)

type Building struct {
	Owner     string
	Structure Structure
}

type Road struct {
	Owner string
	E     [2]int16
}

type Board struct {
	Tiles      []Tile
	Graph      []Node
	RobberTile string
	Buildings  []Building
	Roads      []Road
}

func New() Board {
	var b Board

	tids := [19]string{"A", "B", "C", "G", "L", "P", "S", "R", "Q", "M", "H", "D", "E", "F", "K", "O", "N", "I", "J"} // ordered according to perimeter
	resources := [19]Resource{Wood, Wood, Wood, Wood, Wheat, Wheat, Wheat, Wheat, Sheep, Sheep, Sheep, Sheep, Rock, Rock, Rock, Brick, Brick, Brick, Desert}
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
			N: 1,
			E: []int16{4, 5},
		},
		{
			N: 2,
			E: []int16{5, 6},
		},
		{
			N: 3,
			E: []int16{6, 7},
		},
		{
			N: 4,
			E: []int16{1, 8},
		},
		{
			N: 5,
			E: []int16{1, 2, 9},
		},
		{
			N: 6,
			E: []int16{2, 3, 10},
		},
		{
			N: 7,
			E: []int16{3, 11},
		},
		{
			N: 8,
			E: []int16{4, 12, 13},
		},
		{
			N: 9,
			E: []int16{5, 13, 14},
		},
		{
			N: 10,
			E: []int16{6, 14, 15},
		},
		{
			N: 11,
			E: []int16{7, 15, 16},
		},
		{
			N: 12,
			E: []int16{8, 17},
		},
		{
			N: 13,
			E: []int16{8, 9, 18},
		},
		{
			N: 14,
			E: []int16{9, 10, 19},
		},
		{
			N: 15,
			E: []int16{10, 11, 20},
		},
		{
			N: 16,
			E: []int16{11, 21},
		},
		{
			N: 17,
			E: []int16{12, 22, 23},
		},
		{
			N: 18,
			E: []int16{13, 23, 24},
		},
		{
			N: 19,
			E: []int16{14, 24, 25},
		},
		{
			N: 20,
			E: []int16{15, 25, 26},
		},
		{
			N: 21,
			E: []int16{16, 26, 27},
		},
		{
			N: 22,
			E: []int16{17, 28},
		},
		{
			N: 23,
			E: []int16{17, 18, 29},
		},
		{
			N: 24,
			E: []int16{18, 19, 30},
		},
		{
			N: 25,
			E: []int16{19, 20, 31},
		},
		{
			N: 26,
			E: []int16{20, 21, 32},
		},
		{
			N: 27,
			E: []int16{21, 33},
		},
		{
			N: 28,
			E: []int16{22, 34},
		},
		{
			N: 29,
			E: []int16{23, 34, 35},
		},
		{
			N: 30,
			E: []int16{24, 35, 36},
		},
		{
			N: 31,
			E: []int16{25, 36, 37},
		},
		{
			N: 32,
			E: []int16{26, 37, 38},
		},
		{
			N: 33,
			E: []int16{27, 38},
		},
		{
			N: 34,
			E: []int16{28, 29, 39},
		},
		{
			N: 35,
			E: []int16{29, 30, 40},
		},
		{
			N: 36,
			E: []int16{30, 31, 41},
		},
		{
			N: 37,
			E: []int16{31, 32, 42},
		},
		{
			N: 38,
			E: []int16{32, 33, 43},
		},
		{
			N: 39,
			E: []int16{34, 44},
		},
		{
			N: 40,
			E: []int16{35, 44, 45},
		},
		{
			N: 41,
			E: []int16{36, 45, 46},
		},
		{
			N: 42,
			E: []int16{37, 46, 47},
		},
		{
			N: 43,
			E: []int16{38, 47},
		},
		{
			N: 44,
			E: []int16{39, 40, 48},
		},
		{
			N: 45,
			E: []int16{40, 41, 49},
		},
		{
			N: 46,
			E: []int16{41, 42, 50},
		},
		{
			N: 47,
			E: []int16{42, 43, 51},
		},
		{
			N: 48,
			E: []int16{44, 52},
		},
		{
			N: 49,
			E: []int16{45, 52, 53},
		},
		{
			N: 50,
			E: []int16{46, 53, 54},
		},
		{
			N: 51,
			E: []int16{47, 54},
		},
		{
			N: 52,
			E: []int16{48, 49},
		},
		{
			N: 53,
			E: []int16{49, 50},
		},
		{
			N: 54,
			E: []int16{50, 51},
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
