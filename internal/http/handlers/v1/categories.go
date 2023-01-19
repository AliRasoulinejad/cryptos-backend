package v1

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

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
	All() func(ctx echo.Context) error
	Get() func(ctx echo.Context) error
	Top() func(ctx echo.Context) error
}

type category struct {
	repositories *app.Repositories
}

func NewCategoryHandler(repositories *app.Repositories) Category {
	return category{repositories: repositories}
}

func (c category) All() func(ctx echo.Context) error {
	return func(ctx echo.Context) error {
		categories, err := c.repositories.CategoryRepo.SelectAll()
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

func (c category) Get() func(ctx echo.Context) error {
	return func(ctx echo.Context) error {
		slug := ctx.Param("slug")
		cat, err := c.repositories.CategoryRepo.GetBySlug(slug)
		if err != nil {
			log.Logger.WithError(err).Errorf("error in get category by slug")

			return fmt.Errorf("error in get category by slug")
		}

		response := categoryResponse{cat.ID, cat.Title, cat.Slug, cat.Image, cat.Description}

		return ctx.JSON(http.StatusOK, response)
	}
}

func (c category) Top() func(ctx echo.Context) error {
	return func(ctx echo.Context) error {
		cnt, err := strconv.Atoi(ctx.Request().URL.Query().Get("count"))
		if err != nil {
			log.Logger.WithError(err).Error("count number is not integer")

			return fmt.Errorf("count number is not integer")
		}

		categories, err := c.repositories.CategoryRepo.SelectTopN(cnt)
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
