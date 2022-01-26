package board

type Tile struct {
	Number    int16
	Resource  Resource
	Nodes     []int16
	HasRobber bool
}
