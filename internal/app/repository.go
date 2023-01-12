package app

import (
	"github.com/AliRasoulinejad/cryptos-backend/internal/http/repositories"
)

type Repositories struct {
	CategoryRepo repositories.Category
	BlogRepo     repositories.Blog
}

func (application *Application) WithRepositories() {
	application.Repositories = new(Repositories)
	application.Repositories.CategoryRepo = repositories.NewCategoryRepo(application.DB)
	application.Repositories.BlogRepo = repositories.NewBlogRepo(application.DB)
}
