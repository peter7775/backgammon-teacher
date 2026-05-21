package ports

type ProgressReader interface {
	WeakAreas(userID string) ([]string, error)
}
