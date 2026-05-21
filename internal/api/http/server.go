package http

import (
	stdhttp "net/http"

	coachapp "backgammon-teacher/internal/modules/coach/app"
	playapp "backgammon-teacher/internal/modules/play/app"
	playinfra "backgammon-teacher/internal/modules/play/infra"
)

type Server struct {
	mux        *stdhttp.ServeMux
	startGame  playapp.StartGame
	getGame    playapp.GetGame
	submitMove playapp.SubmitMove
	hint       coachapp.GenerateHint
}

func NewServer() *Server {
	gameRepo := playinfra.NewSQLiteGameRepository()

	s := &Server{
		mux:        stdhttp.NewServeMux(),
		startGame:  playapp.StartGame{Games: gameRepo},
		getGame:    playapp.GetGame{Games: gameRepo},
		submitMove: playapp.SubmitMove{Games: gameRepo},
		hint:       coachapp.GenerateHint{},
	}

	s.registerHealthRoutes()
	s.registerPlayRoutes()
	return s
}

func (s *Server) ServeHTTP(w stdhttp.ResponseWriter, r *stdhttp.Request) {
	s.mux.ServeHTTP(w, r)
}

func (s *Server) registerHealthRoutes() {
	s.mux.HandleFunc("/health", func(w stdhttp.ResponseWriter, r *stdhttp.Request) {
		writeJSON(w, stdhttp.StatusOK, map[string]any{"status": "ok"})
	})
}
