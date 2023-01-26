package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/AliRasoulinejad/cryptos-backend/internal/app"
)

func Index() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, app.Banner())
	}
}

func Health() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		return ctx.NoContent(http.StatusNoContent)
	}
}

func Metric() echo.HandlerFunc {
	prometheusHandler := promhttp.Handler()
	return func(c echo.Context) error {
		prometheusHandler.ServeHTTP(c.Response(), c.Request())
		return nil
	}
}
