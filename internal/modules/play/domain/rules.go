package domain

func NewStartingPosition() Position {
	var p Position
	p.Turn = 0
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
