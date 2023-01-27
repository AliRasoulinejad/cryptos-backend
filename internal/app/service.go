package app

import (
	"github.com/AliRasoulinejad/cryptos-backend/internal/services"
)

type Services struct {
	BlogService     services.BlogService
	CategoryService services.CategoryService
	CommentService  services.CommentService
}

func (application *Application) WithServices() {
	application.Services = new(Services)
	application.Services.BlogService = services.NewBlogService(application.Repositories.BlogRepo, application.Tracer)
	application.Services.CategoryService = services.NewCategoryService(application.Repositories.CategoryRepo, application.Tracer)
	application.Services.CommentService = services.NewCommentService(
		application.Repositories.CommentRepo, application.Repositories.UserRepo, application.Tracer,
	)
}
