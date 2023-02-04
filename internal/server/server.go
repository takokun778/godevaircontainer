package server

import (
	"context"
	"godevaircontainer/pkg/log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/brpaz/echozap"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type Server struct {
	port string
}

func New(port string) *Server {
	return &Server{
		port: port,
	}
}

func (s *Server) Run() {
	address := strings.Join([]string{":", s.port}, "")

	e := echo.New()

	log, _ := zap.NewProduction()

	e.Use(echozap.ZapLogger(log))

	e.GET("/", handle)

	go func() {
		if err := e.Start(address); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, os.Interrupt)

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}

func handle(c echo.Context) error {
	log.Log().Info("OK")
	return c.JSON(http.StatusOK, "OK")
}
