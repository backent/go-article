package user

import (
	"context"
	"database/sql"

	"github.com/backent/go-article/models"
)

type RepositoryUserInterface interface {
	Create(ctx context.Context, tx *sql.Tx, user models.User) (models.User, error)
	Update(ctx context.Context, tx *sql.Tx, user models.User) (models.User, error)
	Delete(ctx context.Context, tx *sql.Tx, id int) error
	FindById(ctx context.Context, tx *sql.Tx, id int) (models.User, error)
	FindAll(ctx context.Context, tx *sql.Tx) ([]models.User, error)
	FindAllWithArticles(ctx context.Context, tx *sql.Tx) ([]models.User, error)
	FindByUsername(ctx context.Context, tx *sql.Tx, username string) (models.User, error)
}
