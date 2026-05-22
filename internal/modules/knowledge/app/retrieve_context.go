package app

import (
	"sort"

	knowledge "backgammon-teacher/internal/modules/knowledge/domain"
	"backgammon-teacher/internal/modules/knowledge/ports"
)

type RetrieveContext struct {
	Local  ports.LocalVectorRepository
	Shared ports.SharedVectorRepository
}

func (uc RetrieveContext) Execute(q knowledge.Query) ([]knowledge.Match, error) {
	local, err := uc.Local.Search(q.UserID, q)
	if err != nil {
		return nil, err
	}
	shared, err := uc.Shared.Search(q)
	if err != nil {
		return nil, err
	}

	merged := append(local, shared...)
	sort.SliceStable(merged, func(i, j int) bool {
		if merged[i].Score == merged[j].Score {
			return merged[i].Scope == knowledge.ScopeLocal && merged[j].Scope == knowledge.ScopeShared
		}
		return merged[i].Score > merged[j].Score
	})
	if q.TopK > 0 && len(merged) > q.TopK {
		merged = merged[:q.TopK]
	}
	return merged, nil
}
