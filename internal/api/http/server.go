package http

import (
	"net/http"
	coachapp "backgammon-teacher/internal/modules/coach/app"
)

type Server struct {
	mux  *http.ServeMux
	hint coachapp.GenerateHintFunc
}

func NewServer(generate coachapp.GenerateHintFunc) *Server {
	if generate == nil {
		generate = coachapp.DefaultGenerateHint
	}
	s := &Server{mux: http.NewServeMux(), hint: generate}
	s.routes()
	return s
}

func (s *Server) routes() {
	s.routesPlay()
}

func (s *Server) Handler() http.Handler { return s.mux }

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}
