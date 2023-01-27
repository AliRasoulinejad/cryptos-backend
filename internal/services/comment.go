package services

import (
	"context"

	"go.opentelemetry.io/otel/trace"

	"github.com/AliRasoulinejad/cryptos-backend/internal/log"
	"github.com/AliRasoulinejad/cryptos-backend/internal/models"
	"github.com/AliRasoulinejad/cryptos-backend/internal/repositories"
)

type CommentService interface {
	FetchByBlogID(ctx context.Context, blogID int64) (*[]models.Comment, error)
}

type commentService struct {
	commentRepo repositories.Comment
	userRepo    repositories.User
	tracer      trace.Tracer
}

func NewCommentService(commentRepo repositories.Comment, userRepo repositories.User, tracer trace.Tracer) CommentService {
	return &commentService{commentRepo: commentRepo, userRepo: userRepo, tracer: tracer}
}

func (c *commentService) FetchByBlogID(ctx context.Context, blogID int64) (*[]models.Comment, error) {
	spanCtx, span := c.tracer.Start(ctx, "blog-repository: SelectAll")
	defer span.End()

	commentsDTO, err := c.commentRepo.SelectByBlogID(spanCtx, blogID)
	if err != nil {
		log.Logger.WithContext(spanCtx).WithError(err).Error("error in get blog comments from repo")

		return nil, err
	}

	comments := make([]models.Comment, len(*commentsDTO))
	for i, commentDTO := range *commentsDTO {
		uDTO := commentDTO.User
		user := models.NewBasicUser(uDTO.ID, uDTO.Name, uDTO.UserName, uDTO.Image, uDTO.ImageURL)
		comments[i] = models.Comment{
			ID:        commentDTO.ID,
			CreatedAt: commentDTO.CreatedAt,
			UpdatedAt: commentDTO.UpdatedAt,
			Text:      commentDTO.Text,
			Blog:      commentDTO.BlogID,
			User:      user,
			Accepted:  commentDTO.Accepted,
		}
	}

	return &comments, nil
}
