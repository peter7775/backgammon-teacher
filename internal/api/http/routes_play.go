package http

import (
	stdhttp "net/http"
	"strings"

	"backgammon-teacher/internal/api/dto"
	play "backgammon-teacher/internal/modules/play/domain"
)

func (s *Server) registerPlayRoutes() {
	s.mux.HandleFunc("/api/v1/games", s.handleGames)
	s.mux.HandleFunc("/api/v1/games/", s.handleGameByID)
}

func (s *Server) handleGames(w stdhttp.ResponseWriter, r *stdhttp.Request) {
	if r.Method != stdhttp.MethodPost {
		writeError(w, stdhttp.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var req dto.StartGameRequest
	if err := decodeJSON(r, &req); err != nil {
		writeError(w, stdhttp.StatusBadRequest, "invalid json body")
		return
	}

	gameID := req.GameID
	if gameID == "" {
		gameID = req.UserID
	}
	if gameID == "" {
		writeError(w, stdhttp.StatusBadRequest, "gameId or userId is required")
		return
	}

	game, err := s.startGame.Execute(gameID)
	if err != nil {
		writeError(w, stdhttp.StatusInternalServerError, err.Error())
		return
	}

	writeJSON(w, stdhttp.StatusCreated, dto.GameResponseFromDomain(game))
}

func (s *Server) handleGameByID(w stdhttp.ResponseWriter, r *stdhttp.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/api/v1/games/")
	parts := strings.Split(strings.Trim(path, "/"), "/")
	if len(parts) == 0 || parts[0] == "" {
		writeError(w, stdhttp.StatusNotFound, "game id is required")
		return
	}

	gameID := parts[0]
	if len(parts) == 1 {
		if r.Method != stdhttp.MethodGet {
			writeError(w, stdhttp.StatusMethodNotAllowed, "method not allowed")
			return
		}
		s.handleGetGame(w, r, gameID)
		return
	}

	switch parts[1] {
	case "moves":
		if r.Method != stdhttp.MethodPost {
			writeError(w, stdhttp.StatusMethodNotAllowed, "method not allowed")
			return
		}
		s.handleSubmitMove(w, r, gameID)
	case "hint":
		if r.Method != stdhttp.MethodPost {
			writeError(w, stdhttp.StatusMethodNotAllowed, "method not allowed")
			return
		}
		s.handleHint(w, r, gameID)
	default:
		writeError(w, stdhttp.StatusNotFound, "route not found")
	}
}

func (s *Server) handleGetGame(w stdhttp.ResponseWriter, r *stdhttp.Request, gameID string) {
	game, err := s.getGame.Execute(gameID)
	if err != nil {
		writeError(w, stdhttp.StatusNotFound, err.Error())
		return
	}
	writeJSON(w, stdhttp.StatusOK, dto.GameResponseFromDomain(game))
}

func (s *Server) handleSubmitMove(w stdhttp.ResponseWriter, r *stdhttp.Request, gameID string) {
	var req dto.SubmitMoveRequest
	if err := decodeJSON(r, &req); err != nil {
		writeError(w, stdhttp.StatusBadRequest, "invalid json body")
		return
	}

	move := play.Move{Steps: make([]play.Step, 0, len(req.Steps))}
	for _, st := range req.Steps {
		move.Steps = append(move.Steps, play.Step{From: st.From, To: st.To, Pips: st.Pips})
	}

	game, err := s.submitMove.Execute(gameID, move)
	if err != nil {
		writeError(w, stdhttp.StatusNotFound, err.Error())
		return
	}

	writeJSON(w, stdhttp.StatusOK, dto.GameResponseFromDomain(game))
}

func (s *Server) handleHint(w stdhttp.ResponseWriter, r *stdhttp.Request, gameID string) {
	hint, err := s.hint.Execute(gameID)
	if err != nil {
		writeError(w, stdhttp.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, stdhttp.StatusOK, dto.HintResponse{Title: hint.Title, Message: hint.Message})
}
