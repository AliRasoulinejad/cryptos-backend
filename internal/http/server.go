package http

import (
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	echomw "github.com/labstack/echo/v4/middleware"

	"github.com/AliRasoulinejad/cryptos-backend/internal/app"
	"github.com/AliRasoulinejad/cryptos-backend/internal/config"
	"github.com/AliRasoulinejad/cryptos-backend/internal/http/handlers"
	v1 "github.com/AliRasoulinejad/cryptos-backend/internal/http/handlers/v1"
)

type server struct {
	e *echo.Echo
}

func NewServer() *server {
	e := echo.New()

	e.HideBanner = true
	e.Server.ReadTimeout = config.C.HTTPServer.ReadTimeout
	e.Server.WriteTimeout = config.C.HTTPServer.WriteTimeout
	e.Server.ReadHeaderTimeout = config.C.HTTPServer.ReadHeaderTimeout
	e.Server.IdleTimeout = config.C.HTTPServer.IdleTimeout

	return &server{
		e: e,
	}
}

func (s *server) Serve(app *app.Application) *server {
	s.e.Pre(echomw.RemoveTrailingSlash())
	s.e.Use(echomw.CORSWithConfig(echomw.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	// Registering routes
	s.e.GET("/", handlers.Index)
	s.e.GET("/health", handlers.Health)

	categoryHandler := v1.NewCategoryHandler(app.Repositories)
	categoryRoutes := s.e.Group("/api/v1/categories")
	{
		categoryRoutes.GET("", categoryHandler.All())
		categoryRoutes.GET("/:slug", categoryHandler.Get())
		categoryRoutes.GET("/top", categoryHandler.Top())
	}

	// Starting the server
	go func() {
		if err := s.e.Start(config.C.HTTPServer.Listen); err != nil && err != http.ErrServerClosed {
			s.e.Logger.Fatal("shutting down the server")
		}
	}()

	return s
}

func (s *server) WaitForSignals(shutdownRequest chan struct{}) (shutdownReady chan struct{}) {
	shutdownReady = make(chan struct{})
	go func() {
		<-shutdownRequest
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := s.e.Shutdown(ctx); err != nil {
			s.e.Logger.Fatal(err)
		}
		shutdownReady <- struct{}{}
	}()
	return
}
