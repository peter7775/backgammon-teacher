package ports

type PositionSource interface {
	CurrentPosition(gameID string) (string, error)
}
