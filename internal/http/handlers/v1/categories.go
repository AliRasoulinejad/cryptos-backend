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

type categoryResponse struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Slug        string `json:"slug"`
	Image       string `json:"image"`
	Description string `json:"description"`
}

type Category interface {
	All() echo.HandlerFunc
	Get() echo.HandlerFunc
	Top() echo.HandlerFunc
}

type category struct {
	repositories *app.Repositories
	tracer       trace.Tracer
}

func NewCategoryHandler(repositories *app.Repositories, trace trace.Tracer) Category {
	return category{repositories: repositories, tracer: trace}
}

func (c category) All() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		baseCtx := ctx.Get(app.SpanCtxName).(context.Context)
		spanCtx, span := c.tracer.Start(baseCtx, "all-handler")
		defer span.End()

		categories, err := c.repositories.CategoryRepo.SelectAll(spanCtx)
		if err != nil {

			log.Logger.WithError(err).Errorf("error in get all categories")

			return fmt.Errorf("error in get all categories")
		}

		categoryResponses := make([]categoryResponse, len(*categories))
		for i, cat := range *categories {
			categoryResponses[i] = categoryResponse{cat.ID, cat.Title, cat.Slug, cat.Image, cat.Description}
		}

		return ctx.JSON(http.StatusOK, categoryResponses)
	}
}

func (c category) Get() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		baseCtx := ctx.Get(app.SpanCtxName).(context.Context)
		spanCtx, span := c.tracer.Start(baseCtx, "get-handler")
		defer span.End()

		slug := ctx.Param("slug")
		cat, err := c.repositories.CategoryRepo.GetBySlug(spanCtx, slug)
		if err != nil {
			log.Logger.WithError(err).Errorf("error in get category by slug")

			return fmt.Errorf("error in get category by slug")
		}

		response := categoryResponse{cat.ID, cat.Title, cat.Slug, cat.Image, cat.Description}

		return ctx.JSON(http.StatusOK, response)
	}
}

func (c category) Top() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		baseCtx := ctx.Get(app.SpanCtxName).(context.Context)
		spanCtx, span := c.tracer.Start(baseCtx, "top-handler")
		defer span.End()

		cnt, err := strconv.Atoi(ctx.Request().URL.Query().Get("count"))
		if err != nil {
			log.Logger.WithError(err).Error("count number is not integer")

			return fmt.Errorf("count number is not integer")
		}

		categories, err := c.repositories.CategoryRepo.SelectTopN(spanCtx, cnt)
		if err != nil {
			log.Logger.WithError(err).Errorf("error in get top categories")

			return fmt.Errorf("error in get top categories")
		}

		categoryResponses := make([]categoryResponse, cnt)
		for i, cat := range *categories {
			categoryResponses[i] = categoryResponse{cat.ID, cat.Title, cat.Slug, cat.Image, cat.Description}
		}

		return ctx.JSON(http.StatusOK, categoryResponses)
	}
}
