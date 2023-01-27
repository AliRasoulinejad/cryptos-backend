package services

import (
	"context"
	"errors"

	"go.opentelemetry.io/otel/trace"

	"github.com/AliRasoulinejad/cryptos-backend/internal/config"
	"github.com/AliRasoulinejad/cryptos-backend/internal/models"
	"github.com/AliRasoulinejad/cryptos-backend/internal/repositories"
)

var (
	ErrMaximumPageSize = errors.New("requested blogs count per page is out of size")
)

type BlogService interface {
	SelectAllByPaginationByCategorySlug(ctx context.Context, blogPerPage, pageNumber int, categoryID int64) (*[]models.Blog, error)
	SelectByCategoryID(ctx context.Context, categoryID int64) (*[]models.Blog, error)
	SelectTopN(ctx context.Context, n int) (*[]models.Blog, error)
	SelectLastN(ctx context.Context, n int) (*[]models.Blog, error)
	GetIDBySlug(ctx context.Context, slug string) (int64, error)
	GetBySlug(ctx context.Context, slug string) (*models.Blog, error)
}

type blogService struct {
	blogRepo repositories.Blog
	tracer   trace.Tracer
}

func NewBlogService(repo repositories.Blog, tracer trace.Tracer) BlogService {
	return &blogService{blogRepo: repo, tracer: tracer}
}

func (b *blogService) SelectAllByPaginationByCategorySlug(ctx context.Context, blogPerPage, pageNumber int, categoryID int64) (*[]models.Blog, error) {
	spanCtx, span := b.tracer.Start(ctx, "blog-service: SelectAllByPaginationByCategorySlug")
	defer span.End()

	max := config.C.Basic.Pagination.MaximumBlogPerPage
	if blogPerPage > max {
		return nil, ErrMaximumPageSize
	}

	blogsDTO, err := b.blogRepo.SelectAllByPaginationByCategorySlug(spanCtx, blogPerPage, pageNumber, categoryID)
	if err != nil {
		return nil, err
	}

	blogs := make([]models.Blog, len(*blogsDTO))
	for i, blogDTO := range *blogsDTO {
		bDTO := blogDTO.Author
		user := models.NewBasicUser(bDTO.ID, bDTO.Name, bDTO.UserName, bDTO.Image, bDTO.ImageURL)
		blogs[i] = models.Blog{
			ID:             blogDTO.ID,
			CreatedAt:      blogDTO.CreatedAt,
			UpdatedAt:      blogDTO.UpdatedAt,
			Title:          blogDTO.Title,
			Slug:           blogDTO.Slug,
			Author:         user,
			ConversationID: blogDTO.ConversationID,
			Content:        blogDTO.Content,
			TextIndex:      blogDTO.TextIndex,
			CategoryID:     blogDTO.CategoryID,
			Image:          blogDTO.Image,
			ReadingTime:    blogDTO.ReadingTime,
			Publish:        blogDTO.Publish,
			LikesCount:     blogDTO.LikesCount,
			DisLikesCount:  blogDTO.DisLikesCount,
		}
	}

	return &blogs, nil
}

func (b *blogService) SelectByCategoryID(ctx context.Context, categoryID int64) (*[]models.Blog, error) {
	spanCtx, span := b.tracer.Start(ctx, "blog-service: SelectByCategoryID")
	defer span.End()

	blogsDTO, err := b.blogRepo.SelectByCategoryID(spanCtx, categoryID)
	if err != nil {
		return nil, err
	}

	blogs := make([]models.Blog, len(*blogsDTO))
	for i, blogDTO := range *blogsDTO {
		bDTO := blogDTO.Author
		user := models.NewBasicUser(bDTO.ID, bDTO.Name, bDTO.UserName, bDTO.Image, bDTO.ImageURL)
		blogs[i] = models.Blog{
			ID:             blogDTO.ID,
			CreatedAt:      blogDTO.CreatedAt,
			UpdatedAt:      blogDTO.UpdatedAt,
			Title:          blogDTO.Title,
			Slug:           blogDTO.Slug,
			Author:         user,
			ConversationID: blogDTO.ConversationID,
			Content:        blogDTO.Content,
			TextIndex:      blogDTO.TextIndex,
			CategoryID:     blogDTO.CategoryID,
			Image:          blogDTO.Image,
			ReadingTime:    blogDTO.ReadingTime,
			Publish:        blogDTO.Publish,
			LikesCount:     blogDTO.LikesCount,
			DisLikesCount:  blogDTO.DisLikesCount,
		}
	}

	return &blogs, nil
}

func (b *blogService) SelectTopN(ctx context.Context, n int) (*[]models.Blog, error) {
	spanCtx, span := b.tracer.Start(ctx, "blog-service: SelectTopN")
	defer span.End()

	blogsDTO, err := b.blogRepo.SelectTopN(spanCtx, n)
	if err != nil {
		return nil, err
	}

	blogs := make([]models.Blog, len(*blogsDTO))
	for i, blogDTO := range *blogsDTO {
		bDTO := blogDTO.Author
		user := models.NewBasicUser(bDTO.ID, bDTO.Name, bDTO.UserName, bDTO.Image, bDTO.ImageURL)
		blogs[i] = models.Blog{
			ID:             blogDTO.ID,
			CreatedAt:      blogDTO.CreatedAt,
			UpdatedAt:      blogDTO.UpdatedAt,
			Title:          blogDTO.Title,
			Slug:           blogDTO.Slug,
			Author:         user,
			ConversationID: blogDTO.ConversationID,
			Content:        blogDTO.Content,
			TextIndex:      blogDTO.TextIndex,
			CategoryID:     blogDTO.CategoryID,
			Image:          blogDTO.Image,
			ReadingTime:    blogDTO.ReadingTime,
			Publish:        blogDTO.Publish,
			LikesCount:     blogDTO.LikesCount,
			DisLikesCount:  blogDTO.DisLikesCount,
		}
	}

	return &blogs, nil
}

func (b *blogService) SelectLastN(ctx context.Context, n int) (*[]models.Blog, error) {
	spanCtx, span := b.tracer.Start(ctx, "blog-service: SelectLastN")
	defer span.End()

	blogsDTO, err := b.blogRepo.SelectLastN(spanCtx, n)
	if err != nil {
		return nil, err
	}

	blogs := make([]models.Blog, len(*blogsDTO))
	for i, blogDTO := range *blogsDTO {
		bDTO := blogDTO.Author
		user := models.NewBasicUser(bDTO.ID, bDTO.Name, bDTO.UserName, bDTO.Image, bDTO.ImageURL)
		blogs[i] = models.Blog{
			ID:             blogDTO.ID,
			CreatedAt:      blogDTO.CreatedAt,
			UpdatedAt:      blogDTO.UpdatedAt,
			Title:          blogDTO.Title,
			Slug:           blogDTO.Slug,
			Author:         user,
			ConversationID: blogDTO.ConversationID,
			Content:        blogDTO.Content,
			TextIndex:      blogDTO.TextIndex,
			CategoryID:     blogDTO.CategoryID,
			Image:          blogDTO.Image,
			ReadingTime:    blogDTO.ReadingTime,
			Publish:        blogDTO.Publish,
			LikesCount:     blogDTO.LikesCount,
			DisLikesCount:  blogDTO.DisLikesCount,
		}
	}

	return &blogs, nil
}

func (b *blogService) GetIDBySlug(ctx context.Context, slug string) (int64, error) {
	spanCtx, span := b.tracer.Start(ctx, "blog-service: GetIDBySlug")
	defer span.End()

	ID, err := b.blogRepo.GetIDBySlug(spanCtx, slug)
	if err != nil {
		return 0, err
	}

	return ID, nil
}

func (b *blogService) GetBySlug(ctx context.Context, slug string) (*models.Blog, error) {
	spanCtx, span := b.tracer.Start(ctx, "blog-service: GetBySlug")
	defer span.End()

	blogDTO, err := b.blogRepo.GetBySlug(spanCtx, slug)
	if err != nil {
		return nil, err
	}

	bDTO := blogDTO.Author
	user := models.NewBasicUser(bDTO.ID, bDTO.Name, bDTO.UserName, bDTO.Image, bDTO.ImageURL)
	blog := models.Blog{
		ID:             blogDTO.ID,
		CreatedAt:      blogDTO.CreatedAt,
		UpdatedAt:      blogDTO.UpdatedAt,
		Title:          blogDTO.Title,
		Slug:           blogDTO.Slug,
		Author:         user,
		ConversationID: blogDTO.ConversationID,
		Content:        blogDTO.Content,
		TextIndex:      blogDTO.TextIndex,
		CategoryID:     blogDTO.CategoryID,
		Image:          blogDTO.Image,
		ReadingTime:    blogDTO.ReadingTime,
		Publish:        blogDTO.Publish,
		LikesCount:     blogDTO.LikesCount,
		DisLikesCount:  blogDTO.DisLikesCount,
	}

	return &blog, nil
}
