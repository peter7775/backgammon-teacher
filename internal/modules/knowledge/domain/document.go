package domain

type Document struct {
	ID        string
	Source    string
	Title     string
	Language  string
	Tags      []string
	Chunks    []Chunk
	Scope     Scope
}

type Scope string

const (
	ScopeLocal  Scope = "local"
	ScopeShared Scope = "shared"
)

type Chunk struct {
	ID         string
	DocumentID string
	Text       string
	Ordinal    int
	Embedding  []float32
	Keywords   []string
	Metadata   map[string]string
}
