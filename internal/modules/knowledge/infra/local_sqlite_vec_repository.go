package infra

import knowledge "backgammon-teacher/internal/modules/knowledge/domain"

type SQLiteVecRepository struct{}

func NewSQLiteVecRepository() *SQLiteVecRepository { return &SQLiteVecRepository{} }

func (r *SQLiteVecRepository) UpsertChunks(userID string, chunks []knowledge.Chunk) error {
	_ = userID
	_ = chunks
	return nil
}

func (r *SQLiteVecRepository) Search(userID string, q knowledge.Query) ([]knowledge.Match, error) {
	_ = userID
	_ = q
	return []knowledge.Match{
		{ChunkID: "local-1", DocumentID: "local-doc", Text: "Local lesson about safe play and reducing direct shots.", Score: 0.91, Source: "local-knowledge.db", Scope: knowledge.ScopeLocal},
	}, nil
}
