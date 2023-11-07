package article

import (
	"context"

	"github.com/backent/go-article/web/article"
)

type ServicesArticleInterface interface {
	Create(ctx context.Context, request article.ArticleRequestCreate) article.ArticleResponse
	Update(ctx context.Context, request article.ArticleRequestUpdate) article.ArticleResponse
	Delete(ctx context.Context, request article.ArticleRequestDelete) article.ArticleResponse
	FindById(ctx context.Context, request article.ArticleRequestFindById) article.ArticleResponse
	FindAll(ctx context.Context, request article.ArticleRequestFindAll) article.ArticleResponse
}
