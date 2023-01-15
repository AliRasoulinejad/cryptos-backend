package app

import (
	repositories2 "github.com/AliRasoulinejad/cryptos-backend/internal/repositories"
)

type Repositories struct {
	CategoryRepo repositories2.Category
	BlogRepo     repositories2.Blog
	CommentRepo  repositories2.Comment
}

func (application *Application) WithRepositories() {
	application.Repositories = new(Repositories)
	application.Repositories.CategoryRepo = repositories2.NewCategoryRepo(application.DB)
	application.Repositories.BlogRepo = repositories2.NewBlogRepo(application.DB)
	application.Repositories.CommentRepo = repositories2.NewCommentRepo(application.DB)
}
