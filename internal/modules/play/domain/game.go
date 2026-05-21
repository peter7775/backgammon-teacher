package domain

type Game struct {
	ID       string
	Position Position
	Cube     Cube
}

func StartGame(id string) Game {
	return Game{ID: id, Position: NewStartingPosition(), Cube: Cube{Value: 1}}
}
