package llm

import "fmt"

type UseCase string

const (
	UseCaseReviewMove UseCase = "review_move"
	UseCaseHardReasoning UseCase = "hard_reasoning"
	UseCaseRAGAnswer UseCase = "rag_answer"
	UseCasePlanning UseCase = "planning"
	UseCaseOffline UseCase = "offline"
)

type Request struct {
	UseCase UseCase
	Prompt string
	Context map[string]any
}

type Response struct {
	Model string
	Content string
}

type Provider interface {
	Name() string
	Model() string
	Generate(Request) (Response, error)
}

type Router struct {
	providers map[UseCase]Provider
	fallbacks map[UseCase]Provider
}

func NewRouter(primary, fallback map[UseCase]Provider) *Router {
	return &Router{providers: primary, fallbacks: fallback}
}

func (r *Router) Route(req Request) (Response, error) {
	if p, ok := r.providers[req.UseCase]; ok && p != nil {
		return p.Generate(req)
	}
	if p, ok := r.fallbacks[req.UseCase]; ok && p != nil {
		return p.Generate(req)
	}
	return Response{}, fmt.Errorf("no provider configured for use case %s", req.UseCase)
}
