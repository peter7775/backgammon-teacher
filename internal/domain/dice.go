package domain

type Dice struct {
	A int8
	B int8
}

func (d Dice) IsDouble() bool { return d.A == d.B }
