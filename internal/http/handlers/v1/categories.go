package v1

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"go.opentelemetry.io/otel/trace"

	"github.com/AliRasoulinejad/cryptos-backend/internal/app"
	"github.com/AliRasoulinejad/cryptos-backend/internal/log"
)

type CategoryHandler interface {
	All() echo.HandlerFunc
	Get() echo.HandlerFunc
	Top() echo.HandlerFunc
}

type category struct {
	services *app.Services
	tracer   trace.Tracer
}

func NewCategoryHandler(services *app.Services, trace trace.Tracer) CategoryHandler {
	return category{services: services, tracer: trace}
}

func (c category) All() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		baseCtx := ctx.Get(app.SpanCtxName).(context.Context)
		spanCtx, span := c.tracer.Start(baseCtx, "all-handler")
		defer span.End()

		categories, err := c.services.CategoryService.All(spanCtx)
		if err != nil {
			log.Logger.WithContext(spanCtx).WithError(err).Errorf("error in get all categories")

			return fmt.Errorf("error in get all categories")
		}

		return ctx.JSON(http.StatusOK, categories)
	}
}

func (c category) Get() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		baseCtx := ctx.Get(app.SpanCtxName).(context.Context)
		spanCtx, span := c.tracer.Start(baseCtx, "get-handler")
		defer span.End()

		slug := ctx.Param("slug")
		cat, err := c.services.CategoryService.GetBySlug(spanCtx, slug)
		if err != nil {
			log.Logger.WithContext(spanCtx).WithError(err).Errorf("error in get category by slug")

			return fmt.Errorf("error in get category by slug")
		}

		return ctx.JSON(http.StatusOK, cat)
	}
}

func (c category) Top() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		baseCtx := ctx.Get(app.SpanCtxName).(context.Context)
		spanCtx, span := c.tracer.Start(baseCtx, "top-handler")
		defer span.End()

		cnt, err := strconv.Atoi(ctx.Request().URL.Query().Get("count"))
		if err != nil {
			log.Logger.WithContext(spanCtx).WithError(err).Error("count number is not integer")

			return fmt.Errorf("count number is not integer")
		}

		categories, err := c.services.CategoryService.TopN(spanCtx, cnt)
		if err != nil {
			log.Logger.WithContext(spanCtx).WithError(err).Errorf("error in get top categories")

			return fmt.Errorf("error in get top categories")
		}

		return ctx.JSON(http.StatusOK, categories)
	}
}
