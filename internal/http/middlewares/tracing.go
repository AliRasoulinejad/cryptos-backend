package middlewares

import (
	"github.com/labstack/echo/v4"
	"go.opentelemetry.io/otel/trace"

	"github.com/AliRasoulinejad/cryptos-backend/internal/app"
)

func TracingMiddleware(tracer trace.Tracer) func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			spanCtx, span := tracer.Start(c.Request().Context(), c.Request().URL.Path)
			c.Set(app.SpanCtxName, spanCtx)
			result := next(c)
			span.End()

			return result
		}
	}
}
