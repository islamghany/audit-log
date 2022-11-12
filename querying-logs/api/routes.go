package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (s *Server) routes() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(s.notFoundResponse)

	router.MethodNotAllowed = http.HandlerFunc(s.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/", s.getLogsHandler)
	router.HandlerFunc(http.MethodGet, "/:id", s.getLogHandler)

	return router
}
