package api

import (
	"auth/pkgs/logger"
	"auth/utils"
	"fmt"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
	logger *logger.Logger
	config utils.Config
}

type NewServerArgs struct {
	Logger *logger.Logger
	Config utils.Config
}

func NewServer(args *NewServerArgs) (*Server, error) {

	s := &Server{
		logger: args.Logger,
		config: args.Config,
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
