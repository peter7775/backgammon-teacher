package ports

import knowledge "backgammon-teacher/internal/modules/knowledge/domain"

type LocalVectorRepository interface {
	UpsertChunks(userID string, chunks []knowledge.Chunk) error
	Search(userID string, q knowledge.Query) ([]knowledge.Match, error)
}

type SharedVectorRepository interface {
	UpsertChunks(chunks []knowledge.Chunk) error
	Search(q knowledge.Query) ([]knowledge.Match, error)
}
