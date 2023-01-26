package middlewares

import (
	"time"

	"github.com/labstack/echo/v4"

	"github.com/AliRasoulinejad/cryptos-backend/internal/log"
)

func LoggingMiddleware(skipURLs []string) echo.MiddlewareFunc {
	sURLs := skipper(skipURLs)

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			startTime := time.Now()
			err := next(ctx)
			if err != nil {
				ctx.Error(err)
			}

			req := ctx.Request()
			res := ctx.Response()

			if _, ok := sURLs[req.RequestURI]; ok {
				return err
			}

			endTime := time.Since(startTime).Milliseconds()
			status := res.Status
			logStr := log.Logger.WithContext(req.Context()).
				WithField("real_ip", ctx.RealIP()).
				WithField("host", req.Host).
				WithField("method", req.Method).
				WithField("url", req.URL.Path).
				WithField("status", status).
				WithField("time(ms)", endTime)

			var message string
			if err != nil {
				if he, ok := err.(*echo.HTTPError); ok {
					if hs, ok := he.Message.(echo.Map); ok {
						message = hs["message"].(string)
					} else {
						message = he.Message.(string)
					}
				} else {
					message = err.Error()
				}
			}

			switch {
			case status >= 500:
				logStr.Error(message)
			case status >= 400:
				logStr.Warn(message)
			case status >= 300:
				logStr.Info("Redirection")
			default:
				logStr.Info("Success")
			}

			return err
		}
	}
}
