package api

import (
	"auth/pkgs/logger"
	"auth/utils"
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
	logger *logger.Logger
	config utils.Config
	wg     sync.WaitGroup
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
func (server *Server) Start(port int) error {

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      server.router,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	shutdownError := make(chan error)

	go func() {
		quit := make(chan os.Signal, 1)

		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

		s := <-quit

		server.logger.PrintInfo("Shutting out the server:", map[string]string{"Signal": s.String()})

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		// Call Shutdown() on the server like before, but now we only send on the
		// shutdownError channel if it returns an error.
		// Call Shutdown() on our server, passing in the context we just made.
		// Shutdown() will return nil if the graceful shutdown was successful, or an
		// error (which may happen because of a problem closing the listeners, or
		// because the shutdown didn't complete before the 5-second context deadline is
		// hit). We relay this return value to the shutdownError channel.
		err := srv.Shutdown(ctx)
		if err != nil {
			shutdownError <- err
		}

		// Log a message to say that we're waiting for any background goroutines to
		// complete their tasks.
		server.logger.PrintInfo("Completing background tasks", map[string]string{
			"addr": srv.Addr,
		})

		// Call Wait() to block until our WaitGroup counter is zero --- essentially
		// blocking until the background goroutines have finished. Then we return nil on
		// the shutdownError channel, to indicate that the shutdown completed without
		// any issues.
		server.wg.Wait()
		shutdownError <- nil
	}()
	server.logger.PrintInfo("Starting Server", map[string]string{
		"addr": srv.Addr,
		"env":  fmt.Sprintf("%d", server.config.PORT),
	})

	err := srv.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	// Otherwise, we wait to receive the return value from Shutdown() on the
	// shutdownError channel. If return value is an error, we know that there was a
	// problem with the graceful shutdown and we return the error.
	err = <-shutdownError
	if err != nil {
		return err
	}

	// At this point we know that the graceful shutdown completed successfully and we
	// log a "stopped server" message.
	server.logger.PrintInfo("stopped server", map[string]string{
		"addr": srv.Addr,
	})

	return nil
}
