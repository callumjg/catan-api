package board

type Building int16

const (
	Settlement Building = 0
	City                = 1
)

type Node struct {
	Id       int16
	E        []int16
	Owner    string
	Building Building
}
