package middlewares

import (
	"time"

	"github.com/labstack/echo/v4"

	"github.com/AliRasoulinejad/cryptos-backend/internal/log"
)

func EchoMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		startTime := time.Now()
		result := next(c)
		endTime := time.Since(startTime).Milliseconds()
		status := c.Response().Status
		logStr := log.Logger.WithContext(c.Request().Context()).
			WithField("url", c.Request().URL.Path).
			WithField("status", status).
			WithField("time(ms)", endTime)
		switch {
		case status < 300:
			logStr.Info("success")
		case status < 400:
			logStr.Info("redirect")
		case status < 500:
			logStr.Warn("not found")
		default:
			logStr.WithError(result).Error("internal server error")
		}

		return result
	}
}
