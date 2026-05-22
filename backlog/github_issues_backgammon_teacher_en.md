# GitHub Issues backlog for backgammon-teacher

## Issue 1
Title: Stabilize project structure and package boundaries
Type: architecture
Priority: high
Milestone: M1 Foundation

Body:
Review and finalize package boundaries between modules, agents, api, infra, observability, and integration so that each layer has clear responsibilities and no direct shortcuts across the system.

Acceptance criteria:
- Package boundaries are documented.
- Forbidden cross-layer dependencies are removed or prevented.
- docs/context-map.md and the architecture diagram are updated.

## Issue 2
Title: Introduce unified dependency injection and composition root
Type: architecture
Priority: high
Milestone: M1 Foundation

Body:
Replace ad hoc initialization of use-cases and repository components with a central composition root for server, web, and desktop runtimes.

Acceptance criteria:
- Clear bootstrap for cmd/server.
- Configurable registries for repositories, agents, and tools.
- Testable initialization without global state.

## Issue 3
Title: Introduce configuration system for environments and feature flags
Type: backend
Priority: high
Milestone: M1 Foundation

Body:
Add configuration support for databases, embedding provider, LLM provider, local/offline mode, agentic features, and RAG mode.

Acceptance criteria:
- Typed config structures.
- .env.example or equivalent template.
- Feature flags for agentic, rag, local-knowledge, shared-knowledge.

## Issue 4
Title: Set up database migrations and schema management
Type: backend
Priority: high
Milestone: M1 Foundation

Body:
Introduce managed SQL migrations for both SQLite and PostgreSQL so the persistence layer is repeatable and deployable.

Acceptance criteria:
- Migrations for game data.
- Migrations for knowledge documents and chunks.
- Documentation for local and server-side setup.

## Issue 5
Title: Implement board position representation and game state serialization
Type: backend
Priority: high
Milestone: M2 Playable core

Body:
Refine Position, Move, Step, Cube, and position serialization so game state can be validated, persisted, and used in tests.

Acceptance criteria:
- Deterministic serialization/deserialization.
- Game state validation.
- Snapshot tests for board positions.

## Issue 6
Title: Implement legal move generator
Type: backend
Priority: critical
Milestone: M2 Playable core

Body:
Add a real legal move generator including bar entry, bearing off, doubles, and rules for mandatory dice usage.

Acceptance criteria:
- Supports standard backgammon rules.
- Tests for edge cases.
- Deterministic output for identical inputs.

## Issue 7
Title: Implement applying moves to the board
Type: backend
Priority: critical
Milestone: M2 Playable core

Body:
Connect SubmitMove to actual board state mutation instead of only storing move history.

Acceptance criteria:
- Moves update board state correctly.
- Illegal moves return validation errors.
- Saved games contain the updated post-move state.

## Issue 8
Title: Implement dice, turn handling, and game lifecycle
Type: backend
Priority: high
Milestone: M2 Playable core

Body:
Add dice rolling, turn switching, game lifecycle state, and game completion handling.

Acceptance criteria:
- Correct turn switching.
- Lifecycle validation.
- API returns game state, active player, and latest roll.

## Issue 9
Title: Implement first heuristic position evaluator
Type: AI/backend
Priority: high
Milestone: M3 Teaching loop

Body:
Replace the placeholder best move with an initial heuristic evaluator based on safety, board strength, pip count, primes, and blot exposure.

Acceptance criteria:
- Evaluator returns comparable scores for candidate moves.
- Heuristic is documented.
- Tests cover basic position types.

## Issue 10
Title: Implement move analysis against candidate moves
Type: AI/backend
Priority: high
Milestone: M3 Teaching loop

Body:
Extend AnalyzeMove so it compares the played move against multiple candidates from the move generator and evaluator instead of a single placeholder move.

Acceptance criteria:
- Top N candidates returned.
- Equity loss or equivalent scoring.
- Classifications: best, inaccuracy, mistake, blunder.

## Issue 11
Title: Implement coach explanation layer
Type: AI/product
Priority: high
Milestone: M3 Teaching loop

Body:
Create an explanation layer that translates analytical outputs into didactic feedback according to the player's level.

Acceptance criteria:
- Multiple explanation difficulty levels.
- Consistent terminology.
- Clear separation between structured explanation and natural-language rendering.

## Issue 12
Title: Add review endpoint with teaching payload
Type: API
Priority: high
Milestone: M3 Teaching loop

Body:
Introduce a unified endpoint that returns move analysis, explanation, and recommendation after a player move.

Acceptance criteria:
- Endpoint returns analysis + hint + next action.
- Clearly defined DTOs.
- API contract tests.

## Issue 13
Title: Design SQL schema for knowledge documents and chunks
Type: backend/data
Priority: high
Milestone: M4 Hybrid RAG

Body:
Design tables for documents, chunks, metadata, tags, scope, and embedding references for both SQLite and PostgreSQL.

Acceptance criteria:
- SQL schema for local and shared knowledge stores.
- Indexing strategy.
- Data model for citations and provenance.

## Issue 14
Title: Implement document ingestion pipeline
Type: AI/backend
Priority: high
Milestone: M4 Hybrid RAG

Body:
Add document import, chunking, metadata enrichment, embedding generation, and storage writing for the knowledge system.

Acceptance criteria:
- Pipeline works for both shared and local scope.
- Chunking is configurable.
- Import audit log is available.

## Issue 15
Title: Implement embedded local RAG store using SQLite vector layer
Type: backend/desktop
Priority: high
Milestone: M4 Hybrid RAG

Body:
Add a real local knowledge implementation for offline and personalized retrieval without requiring users to install a database server.

Acceptance criteria:
- Local embedding store in an embedded database.
- Search API for top-k similarity.
- Works in desktop mode without extra installation.

## Issue 16
Title: Implement shared knowledge store using PostgreSQL + pgvector
Type: backend/cloud
Priority: high
Milestone: M4 Hybrid RAG

Body:
Add a server-side knowledge store for the central lesson corpus and shared theory with vector retrieval in PostgreSQL.

Acceptance criteria:
- pgvector migrations and indexes.
- Search and upsert repository.
- Filtering by language, topic, and content role.

## Issue 17
Title: Implement hybrid retrieval merging and reranking
Type: AI/backend
Priority: high
Milestone: M4 Hybrid RAG

Body:
Extend retrieval so it combines local and shared search, keyword signals, metadata filters, and recency signals.

Acceptance criteria:
- Merge policy with local preference on ties.
- Reranking pipeline.
- Evaluation set for relevance quality.

## Issue 18
Title: Add grounded answer composer with source citations
Type: AI/product
Priority: high
Milestone: M4 Hybrid RAG

Body:
Create a component that builds responses from retrieved chunks with provenance and citable source metadata.

Acceptance criteria:
- Structured answer payload.
- Attached chunk IDs and source info.
- Fallback when retrieval returns nothing.

## Issue 19
Title: Connect supervisor runtime to HTTP endpoints
Type: backend/AI
Priority: high
Milestone: M5 Agentic orchestration

Body:
Integrate the supervisor layer with API endpoints so agent planning can drive move review and grounded tutoring flows.

Acceptance criteria:
- Endpoint for review_move goal.
- Endpoint for rag_answer goal.
- Unified task envelope.

## Issue 20
Title: Refine planner and introduce task graph
Type: AI/backend
Priority: medium
Milestone: M5 Agentic orchestration

Body:
Extend the rule-based planner into an explicit task graph with branching, retry policy, and conditional steps.

Acceptance criteria:
- Task execution state model.
- Retry policy and timeouts.
- Trace of the plan across individual steps.

## Issue 21
Title: Implement per-user agent memory
Type: AI/backend
Priority: high
Milestone: M5 Agentic orchestration

Body:
Extend the memory layer with short-term session memory and long-term learner memory.

Acceptance criteria:
- Session memory.
- Learner-specific memory store.
- API for reading and writing memory entries.

## Issue 22
Title: Introduce evaluator and reflection loop
Type: AI/quality
Priority: medium
Milestone: M5 Agentic orchestration

Body:
Add an evaluator for agent outputs and a critic step that checks response quality, completeness, and grounding.

Acceptance criteria:
- Minimal validation rules.
- States: accepted, retry, fallback.
- Logged rejection reasons.

## Issue 23
Title: Introduce guardrails for tool calling
Type: AI/security
Priority: high
Milestone: M5 Agentic orchestration

Body:
Define a tool whitelist, step limits, read/write restrictions, and a safe fallback mode when the AI layer fails.

Acceptance criteria:
- Tool policy registry.
- Step limit and timeout limit.
- Fallback mode without LLM.

## Issue 24
Title: Stabilize public API contracts
Type: API
Priority: high
Milestone: M3 Teaching loop

Body:
Define API versioning, error model, typed responses, and backward-compatible contracts for both web and desktop clients.

Acceptance criteria:
- OpenAPI or equivalent specification.
- Standard error envelope.
- Documented versioning for /api/v1.

## Issue 25
Title: Build first web review screen
Type: frontend
Priority: medium
Milestone: M3 Teaching loop

Body:
Create the initial UI for replaying a move, showing analysis, recommended move, and explanation.

Acceptance criteria:
- Board state display.
- Review panel with analysis + hint.
- Dark/light mode support.

## Issue 26
Title: Add desktop shell with embedded local knowledge
Type: desktop
Priority: medium
Milestone: M4 Hybrid RAG

Body:
Enable the desktop variant with a local knowledge database and synchronization with the shared backend.

Acceptance criteria:
- Local knowledge cache.
- Background sync.
- Offline read mode.

## Issue 27
Title: Introduce structured logging and tracing
Type: observability
Priority: high
Milestone: M6 Production readiness

Body:
Add request logs, agent step logs, retrieval logs, and trace correlation across layers.

Acceptance criteria:
- Correlation ID.
- Logs for API, agents, and RAG.
- Basic end-to-end request trace.

## Issue 28
Title: Add quality metrics for analysis and RAG
Type: observability/AI
Priority: medium
Milestone: M6 Production readiness

Body:
Measure retrieval quality, AI latency, fallback counts, hint quality, and teaching loop success.

Acceptance criteria:
- Latency and error-rate metrics.
- Retrieval relevance metrics.
- Dashboard-ready output.

## Issue 29
Title: Introduce authentication and authorization for users and devices
Type: security
Priority: high
Milestone: M6 Production readiness

Body:
Add user identity, bind sessions to devices, and define permissions for private, local, and shared data.

Acceptance criteria:
- Authentication flow.
- Role model for private vs shared knowledge.
- Protected endpoints.

## Issue 30
Title: Add CI/CD, linting, and test gates
Type: devops
Priority: high
Milestone: M1 Foundation

Body:
Set up automated builds, tests, linting, and quality gates for pull requests.

Acceptance criteria:
- CI pipeline for build and test.
- Lint and formatting checks.
- Required merge checks.

## Issue 31
Title: Introduce unit test suites for domain and use-cases
Type: testing
Priority: high
Milestone: M1 Foundation

Body:
Cover core domain rules, use-cases, and repositories with unit tests.

Acceptance criteria:
- Tests for play, analysis, knowledge, and agent runtime.
- Test fixtures.
- Runs in CI.

## Issue 32
Title: Add API contract tests and persistence integration tests
Type: testing
Priority: high
Milestone: M2 Playable core

Body:
Ensure API stability and validate SQLite/PostgreSQL repository behavior with integration tests.

Acceptance criteria:
- API contract tests.
- SQLite integration tests.
- PostgreSQL integration tests.

## Issue 33
Title: Add eval harness for AI and RAG workflows
Type: testing/AI
Priority: high
Milestone: M5 Agentic orchestration

Body:
Create an eval dataset and repeatable quality testing for move analysis, retrieval relevance, and grounded answers.

Acceptance criteria:
- Eval dataset.
- Automated runner.
- Baseline scores.

## Issue 34
Title: Create lesson content model and topic taxonomy
Type: product/content
Priority: high
Milestone: M4 Hybrid RAG

Body:
Define the lesson content model, topic tags, language variants, and difficulty levels for both RAG and tutoring.

Acceptance criteria:
- Topic taxonomy.
- Difficulty and mastery mapping.
- Content model for chunking and retrieval.

## Issue 35
Title: Add localization and terminology glossary
Type: product/localization
Priority: medium
Milestone: M5 Agentic orchestration

Body:
Introduce localization support for coach outputs and RAG answers, including a consistent backgammon terminology glossary.

Acceptance criteria:
- Terminology glossary.
- Localized response rendering.
- Support for at least CZ/EN.

## Issue 36
Title: Define product KPIs for the learning application
Type: product/analytics
Priority: medium
Milestone: M6 Production readiness

Body:
Define how tutor quality will be measured: retention, return usage, number of reviewed moves, mastery improvement, and lesson recommendation relevance.

Acceptance criteria:
- KPI document.
- Analytics event schema.
- Integration with observability and product metrics.
