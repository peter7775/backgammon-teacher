# Backgammon Teacher

Backgammon Teacher is an open-source learning application that combines a real backgammon engine, move analysis, retrieval-augmented knowledge, and agentic AI to help players understand *why* a move is good or bad — not just what the best move is.

The goal of the project is to build a practical training companion for players who want coaching, structured review, and grounded explanations that improve their decision-making over time.

## Vision

Most backgammon tools are strong analyzers, but they are not designed as patient teachers. This project aims to bridge that gap by combining deterministic game logic with AI-powered explanation and lesson delivery.

The long-term vision is to create a system that can:
- review a move and explain the mistake clearly,
- adapt explanations to the learner’s skill level,
- retrieve grounded theory and examples from curated knowledge sources,
- work in both local and shared knowledge modes,
- support web, API, and desktop use cases.

## What the project is building

The repository is evolving toward a professional architecture with these main layers:

- **Playable core** — board representation, legal move generation, turn handling, game lifecycle, and move application.
- **Teaching loop** — position evaluation, move comparison, coaching feedback, and review endpoints.
- **Hybrid RAG** — local and shared knowledge stores, chunking, embeddings, retrieval, reranking, and grounded answers.
- **Agentic orchestration** — planning, memory, guarded tool use, evaluation, and reflection loops.
- **Product layer** — API contracts, web review UI, desktop shell, observability, analytics, and localization.

## Principles

The project follows a few core principles:

- **Deterministic domain first.** The backgammon engine must be correct before the AI layer can be trusted.
- **Grounded AI over vague AI.** Explanations should be tied to analysis and knowledge sources, not free-form guessing.
- **Teach, don’t just score.** The system should help players improve, not only classify moves.
- **Local-first where useful.** Some knowledge and tutoring flows should work in embedded or offline-friendly setups.
- **Professional engineering discipline.** Testing, observability, API stability, and architecture boundaries matter from the start.

## Roadmap themes

The current roadmap is organized around these themes:

1. Foundation — architecture cleanup, configuration, migrations, CI/CD, and tests.
2. Playable core — real game logic and a stable backgammon engine.
3. Teaching loop — move analysis, evaluator, explanations, and review APIs.
4. Hybrid RAG — document ingestion, local/shared vector stores, retrieval, and grounded answers.
5. Agentic orchestration — supervisor runtime, task graph, memory, reflection, and guardrails.
6. Production readiness — auth, logging, metrics, desktop packaging, and product analytics.

## Why contribute

This project is a good fit for contributors who enjoy one or more of these areas:

- Go backend architecture
- game engines and board game logic
- AI agents and orchestration
- retrieval-augmented generation
- embeddings, vector databases, and ranking
- API design and contract testing
- observability and production engineering
- frontend or desktop UX for learning tools
- technical writing, theory curation, and teaching content

Contributions are especially valuable because this project sits at the intersection of game theory, software architecture, and educational AI.

## Ways to contribute

Contributors are welcome across code, architecture, product, and content.

Examples of high-value contributions:
- implementing legal move generation and board-state correctness,
- improving evaluation heuristics and move review quality,
- building ingestion and retrieval for backgammon theory content,
- improving explanation quality and grounded coaching output,
- strengthening tests, tracing, and deployment workflows,
- shaping the web or desktop review experience,
- refining terminology, localization, and lesson structure.

If you want to help, pick an open issue, comment on the direction, or propose a design improvement before starting a larger change.

## Contribution style

Please prefer small, reviewable pull requests over large unstructured rewrites.

Helpful contribution habits:
- open an issue before major architectural changes,
- keep domain logic well tested,
- document non-obvious design decisions,
- keep AI behavior observable and debuggable,
- avoid adding magic without clear system boundaries.

## Project status

The project is currently in an active design and build phase. Core architecture, teaching workflows, hybrid RAG, and agentic orchestration are being shaped in parallel, with a strong focus on making the foundation solid enough for long-term evolution.

## Getting involved

A good starting path for new contributors is:

1. Read the backlog and roadmap themes.
2. Pick a focused issue in foundation, playable core, or testing.
3. Open a discussion in the issue before implementing bigger changes.
4. Submit a PR with clear scope and rationale.

The project welcomes contributors who want to help build a serious, explainable, and educational backgammon platform.
