package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/AliRasoulinejad/cryptos-backend/internal/app"
)

func Index(ctx echo.Context) error {
	return ctx.String(http.StatusOK, app.Banner())
}

func Health(ctx echo.Context) error {
	return ctx.NoContent(http.StatusNoContent)
}
