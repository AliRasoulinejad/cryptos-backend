package v1

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"go.opentelemetry.io/otel/trace"

	"github.com/AliRasoulinejad/cryptos-backend/internal/app"
	"github.com/AliRasoulinejad/cryptos-backend/internal/log"
)

type blogResponse struct {
	ID             uint      `json:"id"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
	Title          string    `json:"title"`
	Slug           string    `json:"slug"`
	AuthorID       int64     `json:"authorID"`
	ConversationID int64     `json:"conversationID"`
	Content        string    `json:"content"`
	TextIndex      string    `json:"textIndex"`
	CategoryID     int64     `json:"categoryID"`
	Image          *string   `json:"image"`
	ReadingTime    int       `json:"readingTime"`
	LikesCount     uint64    `json:"likesCount"`
	DisLikesCount  uint64    `json:"disLikesCount"`
}

type Blog interface {
	All() echo.HandlerFunc
	Get() echo.HandlerFunc
	GetComments() echo.HandlerFunc
	// GetRecommendations() echo.HandlerFunc)
	Popular() echo.HandlerFunc
	Recent() echo.HandlerFunc
}

type blog struct {
	services *app.Services
	tracer   trace.Tracer
}

func NewBlogHandler(services *app.Services, tracer trace.Tracer) Blog {
	return blog{services: services, tracer: tracer}
}

func (b blog) All() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		baseCtx := ctx.Get(app.SpanCtxName).(context.Context)
		spanCtx, span := b.tracer.Start(baseCtx, "all-handler")
		defer span.End()

		page, err := strconv.Atoi(ctx.Request().URL.Query().Get("pageNumber"))
		categorySlug := ctx.Request().URL.Query().Get("categorySlug")
		var categoryID int64
		if categorySlug == "" {
			categoryID = 0
		} else {
			categoryID, err = b.services.CategoryService.GetIDBySlug(spanCtx, categorySlug)
			if err != nil {
				log.Logger.WithContext(spanCtx).WithError(err).Errorf("error in get categoryID")

				return fmt.Errorf("error in get categoryID")
			}

			if categoryID == 0 {
				return ctx.JSON(http.StatusNotFound, "Category Not Found")
			}
		}

		blogs, err := b.services.BlogService.SelectAllByPaginationByCategorySlug(spanCtx, 10, page, categoryID)
		if err != nil {
			log.Logger.WithContext(spanCtx).WithError(err).Errorf("error in get all blogs")

			return fmt.Errorf("error in get all blogs")
		}

		return ctx.JSON(http.StatusOK, blogs)
	}
}

func (b blog) Get() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		baseCtx := ctx.Get(app.SpanCtxName).(context.Context)
		spanCtx, span := b.tracer.Start(baseCtx, "get-handler")
		defer span.End()

		slug := ctx.Param("slug")
		blog, err := b.services.BlogService.GetBySlug(spanCtx, slug)
		if err != nil {
			log.Logger.WithContext(spanCtx).WithError(err).Errorf("error in get blog by slug")

			return fmt.Errorf("error in get blog by slug")
		}

		return ctx.JSON(http.StatusOK, blog)
	}
}

func (b blog) GetComments() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		baseCtx := ctx.Get(app.SpanCtxName).(context.Context)
		spanCtx, span := b.tracer.Start(baseCtx, "get-comments-handler")
		defer span.End()

		slug := ctx.Param("slug")
		blogID, err := b.services.BlogService.GetIDBySlug(spanCtx, slug)
		if err != nil {
			log.Logger.WithContext(spanCtx).WithError(err).Errorf("error in get top blogs")

			return fmt.Errorf("error in get top blogs")
		}

		comments, err := b.services.CommentService.FetchByBlogID(spanCtx, blogID)
		if err != nil {
			log.Logger.WithContext(spanCtx).WithError(err).Errorf("error in fetch blog comments")

			return fmt.Errorf("error in fetch blog comments")
		}

		return ctx.JSON(http.StatusOK, comments)
	}
}

func (b blog) GetRecommendations() echo.HandlerFunc {
	// TODO: recode after establishing an stabling
	return nil
}

func (b blog) Popular() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		baseCtx := ctx.Get(app.SpanCtxName).(context.Context)
		spanCtx, span := b.tracer.Start(baseCtx, "popular-handler")
		defer span.End()

		cnt, err := strconv.Atoi(ctx.Request().URL.Query().Get("count"))
		if err != nil {
			log.Logger.WithContext(spanCtx).WithError(err).Error("count number is not integer")

			return fmt.Errorf("count number is not integer")
		}

		blogs, err := b.services.BlogService.SelectTopN(spanCtx, cnt)
		if err != nil {
			log.Logger.WithContext(spanCtx).WithError(err).Errorf("error in get top blogs")

			return fmt.Errorf("error in get top blogs")
		}

		return ctx.JSON(http.StatusOK, blogs)
	}
}

func (b blog) Recent() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		baseCtx := ctx.Get(app.SpanCtxName).(context.Context)
		spanCtx, span := b.tracer.Start(baseCtx, "recent-handler")
		defer span.End()

		cnt, err := strconv.Atoi(ctx.Request().URL.Query().Get("count"))
		if err != nil {
			log.Logger.WithContext(spanCtx).WithError(err).Error("count number is not integer")

			return fmt.Errorf("count number is not integer")
		}

		blogs, err := b.services.BlogService.SelectLastN(spanCtx, cnt)
		if err != nil {
			log.Logger.WithContext(spanCtx).WithError(err).Errorf("error in get %v recent blogs", cnt)

			return fmt.Errorf("error in get %v recent blogs", cnt)
		}

		return ctx.JSON(http.StatusOK, blogs)
	}
}
