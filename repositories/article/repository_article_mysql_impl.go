package article

import (
	"context"
	"database/sql"
	"errors"

	"github.com/backent/go-article/models"
)

type RepositoryArticleMysqlImpl struct {
}

func NewRepositoryArticleMysqlImpl() RepositoryArticleInterface {
	return &RepositoryArticleMysqlImpl{}
}

func (implementation *RepositoryArticleMysqlImpl) Create(ctx context.Context, tx *sql.Tx, article models.Article) (models.Article, error) {
	query := "INSERT INTO articles (user_id, title, content) VALUES(?, ?, ?)"

	result, err := tx.ExecContext(ctx, query, article.UserId, article.Title, article.Content)
	if err != nil {
		return article, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return article, err
	}

	article.Id = int(id)

	return article, nil
}
func (implementation *RepositoryArticleMysqlImpl) Update(ctx context.Context, tx *sql.Tx, article models.Article) (models.Article, error) {
	query := "UPDATE articles SET user_id = ?, title = ?, content = ? WHERE id = ?"

	result, err := tx.ExecContext(ctx, query, article.UserId, article.Title, article.Content)
	if err != nil {
		return article, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return article, err
	}

	article.Id = int(id)

	return article, nil
}
func (implementation *RepositoryArticleMysqlImpl) Delete(ctx context.Context, tx *sql.Tx, id int) error {
	query := "DELETE FROM articles WHERE id = ?"
	_, err := tx.ExecContext(ctx, query, id)
	return err
}
func (implementation *RepositoryArticleMysqlImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (models.Article, error) {

	var article models.Article

	query := "SELECT id, user_id, title, content FROM articles WHERE id = ?"
	rows, err := tx.QueryContext(ctx, query, id)
	if err != nil {
		return article, err
	}
	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(&article.Id, &article.UserId, &article.Title, &article.Content)
		if err != nil {
			return article, err
		}
	} else {
		return article, errors.New("not found article")
	}

	return article, nil
}
func (implementation *RepositoryArticleMysqlImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]models.Article, error) {

	query := "SELECT id, user_id, title, content FROM articles"
	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var articles []models.Article

	for rows.Next() {
		var article models.Article
		err = rows.Scan(&article.Id, &article.UserId, &article.Title, &article.Content)
		if err != nil {
			return nil, err
		}
		articles = append(articles, article)
	}

	return articles, nil
}

func (implementation *RepositoryArticleMysqlImpl) FindAllWithUserDetail(ctx context.Context, tx *sql.Tx) ([]models.Article, error) {

	query := "SELECT articles.id, articles.user_id, articles.title, articles.content, users.id, users.username FROM articles LEFT JOIN users ON articles.user_id = users.id"
	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var articles []*models.Article
	articleMap := make(map[int]*models.Article)

	for rows.Next() {
		var article models.Article
		var user models.User

		err = rows.Scan(&article.Id, &article.UserId, &article.Title, &article.Content, &user.Id, &user.Username)
		if err != nil {
			return nil, err
		}
		item, found := articleMap[article.Id]
		if !found {
			item = &article
			articleMap[article.Id] = item
			articles = append(articles, item)
		}
		item.UserDetails = user
	}

	var articlesReturn []models.Article
	for _, val := range articles {
		articlesReturn = append(articlesReturn, *val)
	}
	return articlesReturn, nil
}
