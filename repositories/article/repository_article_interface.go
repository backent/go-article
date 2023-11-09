package article

import (
	"context"
	"database/sql"

	"github.com/backent/go-article/models"
)

type RepositoryArticleInterface interface {
	Create(ctx context.Context, tx *sql.Tx, article models.Article) (models.Article, error)
	Update(ctx context.Context, tx *sql.Tx, article models.Article) (models.Article, error)
	Delete(ctx context.Context, tx *sql.Tx, id int) error
	FindById(ctx context.Context, tx *sql.Tx, id int) (models.Article, error)
	FindAll(ctx context.Context, tx *sql.Tx) ([]models.Article, error)
	FindAllWithUserDetail(ctx context.Context, tx *sql.Tx) ([]models.Article, error)
}
