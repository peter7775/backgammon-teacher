package infra

import knowledge "backgammon-teacher/internal/modules/knowledge/domain"

type PGVectorRepository struct{}

func NewPGVectorRepository() *PGVectorRepository { return &PGVectorRepository{} }

func (r *PGVectorRepository) UpsertChunks(chunks []knowledge.Chunk) error {
	_ = chunks
	return nil
}

func (r *PGVectorRepository) Search(q knowledge.Query) ([]knowledge.Match, error) {
	_ = q
	return []knowledge.Match{
		{ChunkID: "shared-1", DocumentID: "shared-doc", Text: "Shared theory note about making points in front of anchors.", Score: 0.84, Source: "postgres:knowledge_chunks", Scope: knowledge.ScopeShared},
	}, nil
}
