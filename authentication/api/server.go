package api

import (
	"auth/pkgs/logger"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
	logger *logger.Logger
}

type NewServerArgs struct {
}

func NewServer(args *NewServerArgs) (*Server, error) {

	s := &Server{
		logger: logger.New(os.Stdout, logger.LevelInfo),
	}

	s.setupRoutes()
	return s, nil
}

func (s *Server) setupRoutes() {
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, "Success")
	})

	s.router = router
}
func (s *Server) Start(port int) {
	s.router.Run(fmt.Sprintf(":%d", port))
}
