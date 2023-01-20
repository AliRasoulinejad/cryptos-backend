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
	"github.com/AliRasoulinejad/cryptos-backend/internal/http/middlewares"
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
		AllowOrigins: config.C.Basic.CORSWhiteList,
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	s.e.Use(middlewares.TracingMiddleware(app.Tracer))
	s.e.Use(middlewares.EchoMiddleware)

	// Registering routes
	s.e.GET("/", handlers.Index)
	s.e.GET("/health", handlers.Health)

	categoryHandler := v1.NewCategoryHandler(app.Repositories, app.Tracer)
	categoryRoutes := s.e.Group("/api/v1/categories")
	{
		categoryRoutes.GET("", categoryHandler.All())
		categoryRoutes.GET("/:slug", categoryHandler.Get())
		categoryRoutes.GET("/top", categoryHandler.Top())
	}

	blogHandler := v1.NewBlogHandler(app.Repositories, app.Tracer)
	blogRoutes := s.e.Group("/api/v1/blogs")
	{
		blogRoutes.GET("", blogHandler.All()) // page , categorySlug
		blogRoutes.GET("/:slug", blogHandler.Get())
		blogRoutes.GET("/:slug/comments", blogHandler.GetComments())
		// blogRoutes.GET("/:slug/recommendations", blogHandler.GetRecommendations())
		blogRoutes.GET("/:slug/recommendations", blogHandler.Popular())
		blogRoutes.GET("/popular", blogHandler.Popular())
		blogRoutes.GET("/recent", blogHandler.Recent())
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
