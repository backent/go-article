package article

import (
	"context"
	"database/sql"

	"github.com/backent/go-article/helpers"
	"github.com/backent/go-article/middlewares"
	"github.com/backent/go-article/models"
	respositoriesArticle "github.com/backent/go-article/repositories/article"
	webArticle "github.com/backent/go-article/web/article"
)

type ServicesArticleImpl struct {
	DB *sql.DB
	respositoriesArticle.RepositoryArticleInterface
	*middlewares.ArticleMiddleware
}

func NewServicesArticleImpl(db *sql.DB, respositoriesArticle respositoriesArticle.RepositoryArticleInterface, articleMiddleware *middlewares.ArticleMiddleware) ServicesArticleInterface {
	return &ServicesArticleImpl{
		DB:                         db,
		RepositoryArticleInterface: respositoriesArticle,
		ArticleMiddleware:          articleMiddleware,
	}
}

func (implementation *ServicesArticleImpl) Create(ctx context.Context, request webArticle.ArticleRequestCreate) webArticle.ArticleResponse {
	tx, err := implementation.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	implementation.ArticleMiddleware.Create(ctx, &request)

	article := models.Article{
		UserId:  request.UserId,
		Title:   request.Title,
		Content: request.Content,
	}
	article, err = implementation.RepositoryArticleInterface.Create(ctx, tx, article)

	helpers.PanicIfError(err)

	return webArticle.ArticleModelToResponse(article)

}
func (implementation *ServicesArticleImpl) Update(ctx context.Context, request webArticle.ArticleRequestUpdate) webArticle.ArticleResponse {
	panic("awd")
}
func (implementation *ServicesArticleImpl) Delete(ctx context.Context, request webArticle.ArticleRequestDelete) webArticle.ArticleResponse {
	panic("awd")
}
func (implementation *ServicesArticleImpl) FindById(ctx context.Context, request webArticle.ArticleRequestFindById) webArticle.ArticleResponse {
	panic("awd")
}
func (implementation *ServicesArticleImpl) FindAll(ctx context.Context, request webArticle.ArticleRequestFindAll) []webArticle.ArticleResponse {
	tx, err := implementation.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	implementation.ArticleMiddleware.FindAll(ctx, &request)

	articles, err := implementation.RepositoryArticleInterface.FindAllWithUserDetail(ctx, tx)
	helpers.PanicIfError(err)

	return webArticle.ArticlesModelToResponses(articles)
}
