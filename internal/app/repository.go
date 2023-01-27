package app

import (
	"github.com/AliRasoulinejad/cryptos-backend/internal/repositories"
)

type Repositories struct {
	CategoryRepo repositories.Category
	BlogRepo     repositories.Blog
	CommentRepo  repositories.Comment
	UserRepo     repositories.User
}

func (application *Application) WithRepositories() {
	application.Repositories = new(Repositories)
	application.Repositories.CategoryRepo = repositories.NewCategoryRepo(application.DB, application.Tracer)
	application.Repositories.BlogRepo = repositories.NewBlogRepo(application.DB, application.Tracer)
	application.Repositories.CommentRepo = repositories.NewCommentRepo(application.DB, application.Tracer)
	application.Repositories.UserRepo = repositories.NewUserRepo(application.DB, application.Tracer)
}
