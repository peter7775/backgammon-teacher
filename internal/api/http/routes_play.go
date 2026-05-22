package http

import "net/http"

func (s *Server) routesPlay() {
	if s.mux == nil {
		s.mux = http.NewServeMux()
	}
	s.mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
	})
}
