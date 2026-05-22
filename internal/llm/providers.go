package llm

import "fmt"

type ClaudeProvider struct{ model string }
func NewClaudeProvider(model string) *ClaudeProvider { return &ClaudeProvider{model: model} }
func (c *ClaudeProvider) Name() string { return "anthropic" }
func (c *ClaudeProvider) Model() string { return c.model }
func (c *ClaudeProvider) Generate(req Request) (Response, error) { return Response{Model: c.model, Content: fmt.Sprintf("[%s] %s", c.model, req.Prompt)}, nil }

type OllamaProvider struct{ model string }
func NewOllamaProvider(model string) *OllamaProvider { return &OllamaProvider{model: model} }
func (o *OllamaProvider) Name() string { return "ollama" }
func (o *OllamaProvider) Model() string { return o.model }
func (o *OllamaProvider) Generate(req Request) (Response, error) { return Response{Model: o.model, Content: fmt.Sprintf("[%s] %s", o.model, req.Prompt)}, nil }
