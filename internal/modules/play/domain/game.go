package domain

type Game struct {
	ID       string
	Position Position
	Cube     Cube
	Moves    []Move
}

func StartGame(id string) Game {
	return Game{ID: id, Position: NewStartingPosition(), Cube: Cube{Value: 1}, Moves: []Move{}}
}

func (g Game) ApplyMove(move Move) Game {
	g.Moves = append(g.Moves, move)
	return g
}
