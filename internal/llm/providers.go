package llm

import "fmt"

type ClaudeProvider struct {
	providerName string
	modelName    string
}

func NewClaudeProvider(model string) *ClaudeProvider {
	return &ClaudeProvider{providerName: "anthropic", modelName: model}
}

func (c *ClaudeProvider) Name() string  { return c.providerName }
func (c *ClaudeProvider) Model() string { return c.modelName }

func (c *ClaudeProvider) Generate(req Request) (Response, error) {
	return Response{Model: c.modelName, Content: fmt.Sprintf("[%s] %s", c.modelName, req.Prompt)}, nil
}

type OllamaProvider struct {
	providerName string
	modelName    string
}

func NewOllamaProvider(model string) *OllamaProvider {
	return &OllamaProvider{providerName: "ollama", modelName: model}
}

func (o *OllamaProvider) Name() string  { return o.providerName }
func (o *OllamaProvider) Model() string { return o.modelName }

func (o *OllamaProvider) Generate(req Request) (Response, error) {
	return Response{Model: o.modelName, Content: fmt.Sprintf("[%s] %s", o.modelName, req.Prompt)}, nil
}
