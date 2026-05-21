package dto

import play "backgammon-teacher/internal/modules/play/domain"

type StartGameRequest struct {
	UserID string `json:"userId"`
	GameID string `json:"gameId"`
}

type StepDTO struct {
	From int8 `json:"from"`
	To   int8 `json:"to"`
	Pips int8 `json:"pips"`
}

type SubmitMoveRequest struct {
	Steps []StepDTO `json:"steps"`
}

type PositionDTO struct {
	Points [24]int8 `json:"points"`
	Bar    [2]int8  `json:"bar"`
	Off    [2]int8  `json:"off"`
	Turn   int8     `json:"turn"`
}

type MoveDTO struct {
	Steps []StepDTO `json:"steps"`
}

type GameResponse struct {
	GameID   string      `json:"gameId"`
	Cube     int         `json:"cube"`
	Turn     int8        `json:"turn"`
	Position PositionDTO `json:"position"`
	Moves    []MoveDTO   `json:"moves"`
}

func PositionDTOFromDomain(p play.Position) PositionDTO {
	return PositionDTO{Points: p.Points, Bar: p.Bar, Off: p.Off, Turn: p.Turn}
}

func MoveDTOFromDomain(m play.Move) MoveDTO {
	out := MoveDTO{Steps: make([]StepDTO, 0, len(m.Steps))}
	for _, s := range m.Steps {
		out.Steps = append(out.Steps, StepDTO{From: s.From, To: s.To, Pips: s.Pips})
	}
	return out
}

func GameResponseFromDomain(g play.Game) GameResponse {
	moves := make([]MoveDTO, 0, len(g.Moves))
	for _, m := range g.Moves {
		moves = append(moves, MoveDTOFromDomain(m))
	}
	return GameResponse{
		GameID:   g.ID,
		Cube:     g.Cube.Value,
		Turn:     g.Position.Turn,
		Position: PositionDTOFromDomain(g.Position),
		Moves:    moves,
	}
}
