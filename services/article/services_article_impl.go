package article

import (
	"context"
	"database/sql"

	"github.com/backent/go-article/helpers"
	"github.com/backent/go-article/middlewares"
	"github.com/backent/go-article/models"
	repositoriesUser "github.com/backent/go-article/repositories/user"
	"github.com/backent/go-article/web/article"
	webArticle "github.com/backent/go-article/web/article"
)

type ServicesArticleImpl struct {
	DB *sql.DB
	repositoriesUser.RepositoryUserInterface
	*middlewares.ArticleMiddleware
}

func NewServicesArticleImpl(db *sql.DB, repositoriesUser repositoriesUser.RepositoryUserInterface, articleMiddleware *middlewares.ArticleMiddleware) ServicesArticleInterface {
	return &ServicesArticleImpl{
		DB:                      db,
		RepositoryUserInterface: repositoriesUser,
		ArticleMiddleware:       articleMiddleware,
	}
}

func (implementation *ServicesArticleImpl) Create(ctx context.Context, request article.ArticleRequestCreate) webArticle.ArticleResponse {
	tx, err := implementation.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	implementation.ArticleMiddleware.Create(ctx, &request)

	article := models.Article{
		UserId:  request.UserId,
		Title:   request.Title,
		Content: request.Content,
	}
	implementation.
		article, err = implementation.RepositoryArticleInterface.Create(ctx, tx, article)

	user
	helpers.PanicIfError(err)

	return webArticle.ArticleModelToArticleResponse(article)

}
func (implementation *ServicesArticleImpl) Update(ctx context.Context, request article.ArticleRequestUpdate) article.ArticleResponse {
	panic("awd")
}
func (implementation *ServicesArticleImpl) Delete(ctx context.Context, request article.ArticleRequestDelete) article.ArticleResponse {
	panic("awd")
}
func (implementation *ServicesArticleImpl) FindById(ctx context.Context, request article.ArticleRequestFindById) article.ArticleResponse {
	panic("awd")
}
func (implementation *ServicesArticleImpl) FindAll(ctx context.Context, request article.ArticleRequestFindAll) article.ArticleResponse {
	panic("awd")
}
