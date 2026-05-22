package app

import (
	knowledge "backgammon-teacher/internal/modules/knowledge/domain"
	"backgammon-teacher/internal/modules/knowledge/ports"
)

type Embedder interface {
	Embed(texts []string) ([][]float32, error)
}

type IngestDocument struct {
	Local  ports.LocalVectorRepository
	Shared ports.SharedVectorRepository
	Embed  Embedder
}

func (uc IngestDocument) Execute(userID string, doc knowledge.Document) error {
	texts := make([]string, 0, len(doc.Chunks))
	for _, c := range doc.Chunks {
		texts = append(texts, c.Text)
	}
	vectors, err := uc.Embed.Embed(texts)
	if err != nil {
		return err
	}
	for i := range doc.Chunks {
		if i < len(vectors) {
			doc.Chunks[i].Embedding = vectors[i]
		}
	}
	if doc.Scope == knowledge.ScopeLocal {
		return uc.Local.UpsertChunks(userID, doc.Chunks)
	}
	return uc.Shared.UpsertChunks(doc.Chunks)
}
