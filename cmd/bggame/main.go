package main

import (
	"fmt"

	"backgammon-go-skeleton/internal/domain"
	"backgammon-go-skeleton/internal/engine"
)

func main() {
	pos := domain.NewStartingPosition()
	dice := domain.Dice{A: 6, B: 1}
	moves := engine.GenerateLegalMoves(pos, dice)

	fmt.Printf("legal moves for %d-%d: %d\n", dice.A, dice.B, len(moves))
	for i, m := range moves {
		fmt.Printf("%d. ", i+1)
		for j, s := range m.Steps {
			if j > 0 {
				fmt.Print(" ")
			}
			fmt.Printf("%d/%d", s.From, s.To)
		}
		fmt.Println()
	}
}
