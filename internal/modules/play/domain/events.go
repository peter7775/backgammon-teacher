package domain

type GameStarted struct { GameID string }
func (GameStarted) Name() string { return "play.game_started" }

type MoveSubmitted struct { GameID string }
func (MoveSubmitted) Name() string { return "play.move_submitted" }
