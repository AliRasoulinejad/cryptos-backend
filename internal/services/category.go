package services

import (
	"context"

	"go.opentelemetry.io/otel/trace"

	"github.com/AliRasoulinejad/cryptos-backend/internal/log"
	"github.com/AliRasoulinejad/cryptos-backend/internal/models"
	"github.com/AliRasoulinejad/cryptos-backend/internal/repositories"
)

type CategoryService interface {
	All(ctx context.Context) (*[]models.Category, error)
	TopN(ctx context.Context, n int) (*[]models.Category, error)
	GetIDBySlug(ctx context.Context, slug string) (int64, error)
	GetBySlug(ctx context.Context, slug string) (*models.Category, error)
}

type category struct {
	categoryRepo repositories.Category
	tracer       trace.Tracer
}

func NewCategoryService(repo repositories.Category, tracer trace.Tracer) CategoryService {
	return &category{categoryRepo: repo, tracer: tracer}
}

func (c *category) All(ctx context.Context) (*[]models.Category, error) {
	spanCtx, span := c.tracer.Start(ctx, "category-service: All")
	defer span.End()

	categoriesDTO, err := c.categoryRepo.SelectAll(spanCtx)
	if err != nil {
		log.Logger.WithContext(spanCtx).WithError(err).Error("error in get all categories from repo")

		return nil, err
	}

	categories := make([]models.Category, len(*categoriesDTO))
	for i, categoryDTO := range *categoriesDTO {
		categories[i] = models.Category{
			ID:          categoryDTO.ID,
			Title:       categoryDTO.Title,
			Slug:        categoryDTO.Slug,
			Image:       categoryDTO.Image,
			Description: categoryDTO.Description,
		}
	}

	return &categories, nil
}

func (c *category) TopN(ctx context.Context, n int) (*[]models.Category, error) {
	spanCtx, span := c.tracer.Start(ctx, "category-service: SelectAll")
	defer span.End()

	categoriesDTO, err := c.categoryRepo.SelectTopN(spanCtx, n)
	if err != nil {
		log.Logger.WithContext(spanCtx).WithError(err).Error("error in get top categories from repo")

		return nil, err
	}

	categories := make([]models.Category, len(*categoriesDTO))
	for i, categoryDTO := range *categoriesDTO {
		categories[i] = models.Category{
			ID:          categoryDTO.ID,
			Title:       categoryDTO.Title,
			Slug:        categoryDTO.Slug,
			Image:       categoryDTO.Image,
			Description: categoryDTO.Description,
		}
	}

	return &categories, nil
}

func (c *category) GetIDBySlug(ctx context.Context, slug string) (int64, error) {
	spanCtx, span := c.tracer.Start(ctx, "category-service: SelectAll")
	defer span.End()

	ID, err := c.categoryRepo.GetIDBySlug(spanCtx, slug)
	if err != nil {
		log.Logger.WithContext(spanCtx).WithError(err).Error("error in get category ID from repo")

		return 0, err
	}

	return ID, nil
}

func (c *category) GetBySlug(ctx context.Context, slug string) (*models.Category, error) {
	spanCtx, span := c.tracer.Start(ctx, "category-service: SelectAll")
	defer span.End()

	catDTO, err := c.categoryRepo.GetBySlug(spanCtx, slug)
	if err != nil {
		log.Logger.WithContext(spanCtx).WithError(err).Error("error in get category from repo")

		return nil, err
	}

	cat := models.Category{
		ID:          catDTO.ID,
		Title:       catDTO.Title,
		Slug:        catDTO.Slug,
		Image:       catDTO.Image,
		Description: catDTO.Description,
	}

	return &cat, nil
}
