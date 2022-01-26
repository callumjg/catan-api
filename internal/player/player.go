package player

type Color string

const (
	Blue   Color = "Blue"
	Red          = "Red"
	Orange       = "Orange"
	White        = "White"
	Green        = "Green"
)

type Player struct {
	Id    string
	Name  string
	Color Color
}
