package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (s *Server) routes() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(s.notFoundResponse)

	router.MethodNotAllowed = http.HandlerFunc(s.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/api/read", s.getLogsHandler)
	router.HandlerFunc(http.MethodGet, "/api/read/:id", s.getLogHandler)

	return router
}
