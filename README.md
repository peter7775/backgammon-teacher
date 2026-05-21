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
- `POST /api/v1/games/{id}/hint`

## Examples

Create game:

```json
{
  "gameId": "game-001"
}
```

Submit move:

```json
{
  "steps": [
    {"from": 13, "to": 7, "pips": 6},
    {"from": 8, "to": 7, "pips": 1}
  ]
}
```
