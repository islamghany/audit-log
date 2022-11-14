package api

import "net/http"

func (s *Server) routes() http.Handler {

	mux := http.NewServeMux()

	mux.HandleFunc("/api/write", s.createLogHandler)
	mux.HandleFunc("*", func(w http.ResponseWriter, r *http.Request) {
		s.notFoundResponse(w, r)
	})
	return mux
}
