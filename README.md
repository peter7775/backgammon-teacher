# backgammon-teacher

Headless DDD skeleton for a backgammon teacher application with web and desktop clients.

## Run server

```bash
go run ./cmd/server
```

Default address is `:8080` and can be changed with `HTTP_ADDR`.

## Endpoints

- `GET /health`
- `POST /api/v1/games`
- `GET /api/v1/games/{id}`
- `POST /api/v1/games/{id}/moves`
- `POST /api/v1/games/{id}/analysis`
- `POST /api/v1/games/{id}/hint`

## Current teaching flow

1. Client submits a move.
2. Backend stores the move.
3. Analysis module compares it with a placeholder best move.
4. Coach module returns teaching feedback, classification and recommended best move.
