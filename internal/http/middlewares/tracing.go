package middlewares

import (
	"github.com/labstack/echo/v4"
	"go.opentelemetry.io/otel/trace"

	"github.com/AliRasoulinejad/cryptos-backend/internal/app"
)

func TracingMiddleware(tracer trace.Tracer, skipURLs []string) func(next echo.HandlerFunc) echo.HandlerFunc {
	sURLs := skipper(skipURLs)

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			if _, ok := sURLs[ctx.Request().RequestURI]; ok {
				return next(ctx)
			}

			spanCtx, span := tracer.Start(ctx.Request().Context(), ctx.Request().URL.Path)
			ctx.Set(app.SpanCtxName, spanCtx)
			result := next(ctx)
			span.End()

			return result
		}
	}
}
