package domain

type Query struct {
	Text      string
	Language  string
	UserID    string
	SessionID string
	Tags      []string
	TopK      int
}

type Match struct {
	ChunkID    string
	DocumentID string
	Text       string
	Score      float64
	Source     string
	Scope      Scope
	Metadata   map[string]string
}
