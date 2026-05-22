package infra

type MockEmbedder struct{}

func (MockEmbedder) Embed(texts []string) ([][]float32, error) {
	out := make([][]float32, 0, len(texts))
	for range texts {
		out = append(out, []float32{0.1, 0.2, 0.3})
	}
	return out, nil
}
