package domain

func PlayerSign(player int8) int8 {
	if player == PlayerWhite {
		return 1
	}
	return -1
}

func Opponent(player int8) int8 {
	if player == PlayerWhite {
		return PlayerBlack
	}
	return PlayerWhite
}
