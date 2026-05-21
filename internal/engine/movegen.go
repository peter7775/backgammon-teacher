package engine

import (
	"strconv"
	"strings"

	"backgammon-go-skeleton/internal/domain"
)

func GenerateLegalMoves(pos domain.Position, dice domain.Dice) []domain.Move {
	seqs := diceOrders(dice)
	bestUsed := -1
	seen := map[string]struct{}{}
	var out []domain.Move

	for _, seq := range seqs {
		moves := generateForSequence(pos, seq)
		for _, m := range moves {
			used := len(m.Steps)
			if used > bestUsed {
				bestUsed = used
				seen = map[string]struct{}{}
				out = out[:0]
			}
			if used != bestUsed {
				continue
			}
			key := moveKey(m)
			if _, ok := seen[key]; ok {
				continue
			}
			seen[key] = struct{}{}
			out = append(out, m)
		}
	}

	if bestUsed == 1 && !dice.IsDouble() {
		high := dice.A
		if dice.B > high {
			high = dice.B
		}
		filtered := out[:0]
		for _, m := range out {
			if len(m.Steps) == 1 && m.Steps[0].Pips == high {
				filtered = append(filtered, m)
			}
		}
		if len(filtered) > 0 {
			out = filtered
		}
	}

	return out
}

func diceOrders(d domain.Dice) [][]int8 {
	if d.IsDouble() {
		return [][]int8{{d.A, d.A, d.A, d.A}}
	}
	if d.A == d.B {
		return [][]int8{{d.A, d.B}}
	}
	return [][]int8{{d.A, d.B}, {d.B, d.A}}
}

func generateForSequence(pos domain.Position, seq []int8) []domain.Move {
	results := []domain.Move{{}}
	for _, die := range seq {
		var next []domain.Move
		for _, partial := range results {
			cur := applyMoveUnchecked(pos, partial)
			candidates := legalSteps(cur, die)
			if len(candidates) == 0 {
				next = append(next, partial)
				continue
			}
			for _, step := range candidates {
				m := cloneMove(partial)
				m.Steps = append(m.Steps, step)
				next = append(next, m)
			}
		}
		results = dedupMoves(next)
	}
	return results
}

func legalSteps(pos domain.Position, die int8) []domain.Step {
	player := pos.Turn
	if pos.Bar[player] > 0 {
		if st, ok := enterFromBar(pos, die); ok {
			return []domain.Step{st}
		}
		return nil
	}

	var steps []domain.Step
	for from := int8(1); from <= 24; from++ {
		if !hasChecker(pos, player, from) {
			continue
		}
		to, bearOff, ok := destination(pos, player, from, die)
		if !ok {
			continue
		}
		if bearOff {
			steps = append(steps, domain.Step{From: from, To: 0, Pips: die})
		} else if isOpen(pos, player, to) {
			steps = append(steps, domain.Step{From: from, To: to, Pips: die})
		}
	}
	return steps
}

func enterFromBar(pos domain.Position, die int8) (domain.Step, bool) {
	player := pos.Turn
	var to int8
	if player == domain.PlayerWhite {
		to = 25 - die
	} else {
		to = die
	}
	if !isOpen(pos, player, to) {
		return domain.Step{}, false
	}
	return domain.Step{From: -1, To: to, Pips: die}, true
}

func destination(pos domain.Position, player, from, die int8) (to int8, bearOff bool, ok bool) {
	if player == domain.PlayerWhite {
		to = from - die
		if to >= 1 {
			return to, false, true
		}
		if !canBearOff(pos, player) {
			return 0, false, false
		}
		if from == die {
			return 0, true, true
		}
		if from < die && !hasHigherHomeChecker(pos, player, from) {
			return 0, true, true
		}
		return 0, false, false
	}

	to = from + die
	if to <= 24 {
		return to, false, true
	}
	if !canBearOff(pos, player) {
		return 0, false, false
	}
	need := 25 - from
	if need == die {
		return 0, true, true
	}
	if need < die && !hasHigherHomeChecker(pos, player, from) {
		return 0, true, true
	}
	return 0, false, false
}

func canBearOff(pos domain.Position, player int8) bool {
	if pos.Bar[player] > 0 {
		return false
	}
	if player == domain.PlayerWhite {
		for pt := int8(7); pt <= 24; pt++ {
			if hasChecker(pos, player, pt) {
				return false
			}
		}
		return true
	}
	for pt := int8(1); pt <= 18; pt++ {
		if hasChecker(pos, player, pt) {
			return false
		}
	}
	return true
}

func hasHigherHomeChecker(pos domain.Position, player, from int8) bool {
	if player == domain.PlayerWhite {
		for pt := int8(6); pt > from; pt-- {
			if hasChecker(pos, player, pt) {
				return true
			}
		}
		return false
	}
	for pt := int8(19); pt < from; pt++ {
		if hasChecker(pos, player, pt) {
			return true
		}
	}
	return false
}

func hasChecker(pos domain.Position, player, point int8) bool {
	v := pos.Points[point-1]
	if player == domain.PlayerWhite {
		return v > 0
	}
	return v < 0
}

func isOpen(pos domain.Position, player, point int8) bool {
	v := pos.Points[point-1]
	if player == domain.PlayerWhite {
		return v >= -1
	}
	return v <= 1
}

func applyMoveUnchecked(pos domain.Position, move domain.Move) domain.Position {
	cur := pos.Clone()
	for _, st := range move.Steps {
		cur = applyStep(cur, st)
	}
	return cur
}

func applyStep(pos domain.Position, st domain.Step) domain.Position {
	player := pos.Turn
	sign := int8(1)
	opp := domain.PlayerBlack
	if player == domain.PlayerBlack {
		sign = -1
		opp = domain.PlayerWhite
	}

	if st.From == -1 {
		pos.Bar[player]--
	} else {
		pos.Points[st.From-1] -= sign
	}

	if st.To == 0 {
		pos.Off[player]++
		return pos
	}

	idx := st.To - 1
	if player == domain.PlayerWhite {
		if pos.Points[idx] == -1 {
			pos.Points[idx] = 0
			pos.Bar[opp]++
		}
		pos.Points[idx]++
		return pos
	}

	if pos.Points[idx] == 1 {
		pos.Points[idx] = 0
		pos.Bar[opp]++
	}
	pos.Points[idx]--
	return pos
}

func dedupMoves(moves []domain.Move) []domain.Move {
	seen := map[string]struct{}{}
	out := make([]domain.Move, 0, len(moves))
	for _, m := range moves {
		k := moveKey(m)
		if _, ok := seen[k]; ok {
			continue
		}
		seen[k] = struct{}{}
		out = append(out, m)
	}
	return out
}

func cloneMove(m domain.Move) domain.Move {
	cp := domain.Move{Steps: make([]domain.Step, len(m.Steps))}
	copy(cp.Steps, m.Steps)
	return cp
}

func moveKey(m domain.Move) string {
	parts := make([]string, 0, len(m.Steps))
	for _, s := range m.Steps {
		parts = append(parts, strconv.Itoa(int(s.From))+"/"+strconv.Itoa(int(s.To))+":"+strconv.Itoa(int(s.Pips)))
	}
	return strings.Join(parts, ",")
}
