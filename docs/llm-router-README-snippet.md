## LLM Routing

This project uses a small provider interface and a use-case router. The router selects a model based on the task type, such as planning, move review, RAG answering, or offline fallback.

Primary cloud models:
- Claude Opus 4.7 for planning and hard reasoning.
- Claude Sonnet 4.6 for review, explanations, and grounded answers.

Local fallback:
- gpt-oss-20b via Ollama.
