package domain

type Position struct {
	Points [24]int8
	Bar    [2]int8
	Off    [2]int8
	Turn   int8
}

const (
	PlayerWhite int8 = 0
	PlayerBlack int8 = 1
)

func NewStartingPosition() Position {
	var p Position
	p.Turn = PlayerWhite

	p.Points[23] = 2
	p.Points[12] = 5
	p.Points[7] = 3
	p.Points[5] = 5

	p.Points[0] = -2
	p.Points[11] = -5
	p.Points[16] = -3
	p.Points[18] = -5

	return p
}

func (p Position) Clone() Position { return p }
