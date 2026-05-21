package ports

type LearnerProfileReader interface {
	LevelOf(userID string) (string, error)
}
