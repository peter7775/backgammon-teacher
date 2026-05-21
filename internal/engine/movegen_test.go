package engine

import (
	"testing"

	"backgammon-go-skeleton/internal/domain"
)

func TestGenerateLegalMovesOpening61(t *testing.T) {
	pos := domain.NewStartingPosition()
	moves := GenerateLegalMoves(pos, domain.Dice{A: 6, B: 1})
	if len(moves) == 0 {
		t.Fatal("expected legal moves for opening 6-1")
	}
}

func TestMustEnterFromBarFirst(t *testing.T) {
	var pos domain.Position
	pos.Turn = domain.PlayerWhite
	pos.Bar[domain.PlayerWhite] = 1
	pos.Points[23] = -2
	pos.Points[22] = -2
	pos.Points[21] = -2
	pos.Points[20] = -2
	pos.Points[19] = -2
	pos.Points[18] = -2

	moves := GenerateLegalMoves(pos, domain.Dice{A: 1, B: 2})
	if len(moves) != 0 {
		t.Fatalf("expected no legal moves, got %d", len(moves))
	}
}

func TestBearOffAllowed(t *testing.T) {
	var pos domain.Position
	pos.Turn = domain.PlayerWhite
	pos.Points[0] = 2
	pos.Points[1] = 2
	pos.Points[2] = 2
	pos.Points[3] = 3
	pos.Points[4] = 3
	pos.Points[5] = 3

	moves := GenerateLegalMoves(pos, domain.Dice{A: 6, B: 5})
	if len(moves) == 0 {
		t.Fatal("expected bearoff legal moves")
	}
}
