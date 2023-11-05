package middlewares

import (
	"context"

	repositoriesArticle "github.com/backent/go-article/repositories/article"
	repositoriesAuth "github.com/backent/go-article/repositories/auth"
	repositoriesUser "github.com/backent/go-article/repositories/user"
	webArticle "github.com/backent/go-article/web/article"
	"github.com/go-playground/validator/v10"
)

type ArticleMiddleware struct {
	validator *validator.Validate
	repositoriesAuth.RepositoryAuthInterface
	repositoriesArticle.RepositoryArticleInterface
	repositoriesUser.RepositoryUserInterface
}

func NewArticleMiddleware(validator *validator.Validate, repositoriesAuth repositoriesAuth.RepositoryAuthInterface, repositoriesArticle repositoriesArticle.RepositoryArticleInterface, repositoriesUser repositoriesUser.RepositoryUserInterface) *ArticleMiddleware {
	return &ArticleMiddleware{
		validator:                  validator,
		RepositoryAuthInterface:    repositoriesAuth,
		RepositoryArticleInterface: repositoriesArticle,
		RepositoryUserInterface:    repositoriesUser,
	}
}

func (implementation *ArticleMiddleware) Create(ctx context.Context, request *webArticle.ArticleRequestCreate) {
	ValidateToken(ctx, implementation.RepositoryAuthInterface)

}
func (implementation *ArticleMiddleware) Update(ctx context.Context, request *webArticle.ArticleRequestUpdate) {
}
func (implementation *ArticleMiddleware) Delete(ctx context.Context, request *webArticle.ArticleRequestDelete) {
}
func (implementation *ArticleMiddleware) FindById(ctx context.Context, request *webArticle.ArticleRequestFindById) {
}
func (implementation *ArticleMiddleware) FindAll(ctx context.Context, request *webArticle.ArticleRequestFindAll) {
}
