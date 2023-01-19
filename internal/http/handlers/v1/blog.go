package v1

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/AliRasoulinejad/cryptos-backend/internal/app"
	"github.com/AliRasoulinejad/cryptos-backend/internal/log"
	"github.com/AliRasoulinejad/cryptos-backend/internal/models"
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

type commentResponse struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Text      string    `json:"text"`
	BlogID    int64     `json:"blogID"`
	UserID    int64     `json:"userID"`
	Accepted  bool      `json:"accepted"`
}

type Blog interface {
	All() func(ctx echo.Context) error
	Get() func(ctx echo.Context) error
	GetComments() func(ctx echo.Context) error
	// GetRecommendations() func(ctx echo.Context) error
	Popular() func(ctx echo.Context) error
	Recent() func(ctx echo.Context) error
}

type blog struct {
	repositories *app.Repositories
}

func NewBlogHandler(repositories *app.Repositories) Blog {
	return blog{repositories: repositories}
}

func (b blog) All() func(ctx echo.Context) error {
	return func(ctx echo.Context) error {
		page, err := strconv.Atoi(ctx.Request().URL.Query().Get("pageNumber"))
		categorySlug := ctx.Request().URL.Query().Get("categorySlug")
		var categoryID int64
		if categorySlug == "" {
			categoryID = 0
		} else {
			categoryID, err = b.repositories.CategoryRepo.GetIDBySlug(categorySlug)
			if err != nil {
				log.Logger.WithError(err).Errorf("error in get categoryID")

				return fmt.Errorf("error in get categoryID")
			}

			if categoryID == 0 {
				return ctx.JSON(http.StatusNotFound, "Category Not Found")
			}
		}

		blogs, err := b.repositories.BlogRepo.SelectAllByPaginationByCategorySlug(10, page, categoryID)
		if err != nil {
			log.Logger.WithError(err).Errorf("error in get all blogs")

			return fmt.Errorf("error in get all blogs")
		}

		blogResponses := makeResponseFromBlogModel(*blogs)

		return ctx.JSON(http.StatusOK, blogResponses)
	}
}

func (b blog) Get() func(ctx echo.Context) error {
	return func(ctx echo.Context) error {
		slug := ctx.Param("slug")
		blg, err := b.repositories.BlogRepo.GetBySlug(slug)
		if err != nil {
			log.Logger.WithError(err).Errorf("error in get blog by slug")

			return fmt.Errorf("error in get blog by slug")
		}

		response := blogResponse{blg.ID, blg.CreatedAt, blg.UpdatedAt, blg.Title,
			blg.Slug, blg.AuthorID, blg.ConversationID, blg.Content, blg.TextIndex,
			blg.CategoryID, blg.Image, blg.ReadingTime, blg.LikesCount, blg.DisLikesCount,
		}

		return ctx.JSON(http.StatusOK, response)
	}
}

func (b blog) GetComments() func(ctx echo.Context) error {
	return func(ctx echo.Context) error {
		slug := ctx.Param("slug")
		blogID, err := b.repositories.BlogRepo.GetIDBySlug(slug)
		if err != nil {
			log.Logger.WithError(err).Errorf("error in get top blogs")

			return fmt.Errorf("error in get top blogs")
		}

		comments, err := b.repositories.CommentRepo.SelectByBlogID(blogID)
		if err != nil {
			log.Logger.WithError(err).Errorf("error in fetch blog comments")

			return fmt.Errorf("error in fetch blog comments")
		}

		commentsResponses := make([]commentResponse, len(*comments))
		for i, comment := range *comments {
			commentsResponses[i] = commentResponse{
				comment.ID, comment.CreatedAt, comment.UpdatedAt, comment.Text, comment.BlogID,
				comment.UserID, comment.Accepted,
			}
		}

		return ctx.JSON(http.StatusOK, commentsResponses)
	}
}

func (b blog) GetRecommendations() func(ctx echo.Context) error {
	// TODO: recode after establishing an stabling
	return nil
}

func (b blog) Popular() func(ctx echo.Context) error {
	return func(ctx echo.Context) error {
		cnt, err := strconv.Atoi(ctx.Request().URL.Query().Get("count"))
		if err != nil {
			log.Logger.WithError(err).Error("count number is not integer")

			return fmt.Errorf("count number is not integer")
		}

		blogs, err := b.repositories.BlogRepo.SelectTopN(cnt)
		if err != nil {
			log.Logger.WithError(err).Errorf("error in get top blogs")

			return fmt.Errorf("error in get top blogs")
		}

		blogResponses := makeResponseFromBlogModel(*blogs)

		return ctx.JSON(http.StatusOK, blogResponses)
	}
}

func (b blog) Recent() func(ctx echo.Context) error {
	return func(ctx echo.Context) error {
		cnt, err := strconv.Atoi(ctx.Request().URL.Query().Get("count"))
		if err != nil {
			log.Logger.WithError(err).Error("count number is not integer")

			return fmt.Errorf("count number is not integer")
		}

		blogs, err := b.repositories.BlogRepo.SelectLastN(cnt)
		if err != nil {
			log.Logger.WithError(err).Errorf("error in get %v recent blogs", cnt)

			return fmt.Errorf("error in get %v recent blogs", cnt)
		}

		blogResponses := makeResponseFromBlogModel(*blogs)

		return ctx.JSON(http.StatusOK, blogResponses)
	}
}

func makeResponseFromBlogModel(blogs []models.Blog) []blogResponse {
	blogResponses := make([]blogResponse, len(blogs))
	for i, blg := range blogs {
		blogResponses[i] = blogResponse{blg.ID, blg.CreatedAt, blg.UpdatedAt, blg.Title,
			blg.Slug, blg.AuthorID, blg.ConversationID, blg.Content, blg.TextIndex,
			blg.CategoryID, blg.Image, blg.ReadingTime, blg.LikesCount, blg.DisLikesCount,
		}
	}

	return blogResponses
}
