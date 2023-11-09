package user

import (
	"context"
	"database/sql"
	"errors"

	"github.com/backent/go-article/models"
)

type RepositoryUserMysqlImpl struct {
}

func NewRepositoryMysqlImpl() RepositoryUserInterface {
	return &RepositoryUserMysqlImpl{}
}

func (implementation *RepositoryUserMysqlImpl) Create(ctx context.Context, tx *sql.Tx, user models.User) (models.User, error) {
	sql := "INSERT INTO users (username, name, password) values (?, ?, ?)"
	result, err := tx.ExecContext(ctx, sql, user.Username, user.Name, user.Password)
	if err != nil {
		return user, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return user, err
	}

	user.Id = int(id)
	return user, nil
}
func (implementation *RepositoryUserMysqlImpl) Update(ctx context.Context, tx *sql.Tx, user models.User) (models.User, error) {
	sql := "UPDATE users SET username = ?, name = ?, password = ?  WHERE id = ?"
	_, err := tx.ExecContext(ctx, sql, user.Username, user.Name, user.Password, user.Id)
	if err != nil {
		return user, err
	}

	return user, nil
}
func (implementation *RepositoryUserMysqlImpl) Delete(ctx context.Context, tx *sql.Tx, id int) error {
	sql := "DELETE FROM users WHERE id = ?"
	_, err := tx.ExecContext(ctx, sql, id)
	if err != nil {
		return err
	}

	return nil
}

func (implementation *RepositoryUserMysqlImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (models.User, error) {
	var user models.User
	sql := "SELECT id, username, name, password FROM users WHERE id = ?"
	rows, err := tx.QueryContext(ctx, sql, id)
	if err != nil {
		return user, err
	}
	defer rows.Close()

	if rows.Next() {
		rows.Scan(&user.Id, &user.Name, &user.Username, &user.Password)
	} else {
		return user, errors.New("user not found")
	}

	return user, nil
}
func (implementation *RepositoryUserMysqlImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]models.User, error) {
	sql := "SELECT id, username, name FROM users"
	rows, err := tx.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		user := models.User{}
		rows.Scan(&user.Id, &user.Name, &user.Username)
		users = append(users, user)
	}

	return users, nil
}

func (implementation *RepositoryUserMysqlImpl) FindByUsername(ctx context.Context, tx *sql.Tx, username string) (models.User, error) {
	var user models.User
	sql := "SELECT id, username, name, password FROM users WHERE username = ?"
	rows, err := tx.QueryContext(ctx, sql, username)
	if err != nil {
		return user, err
	}
	defer rows.Close()

	if rows.Next() {
		rows.Scan(&user.Id, &user.Name, &user.Username, &user.Password)
	} else {
		return user, errors.New("user not found")
	}

	return user, nil
}

func (implementation *RepositoryUserMysqlImpl) FindAllWithArticles(ctx context.Context, tx *sql.Tx) ([]models.User, error) {
	query := "SELECT users.id, users.username, users.name, users.password, articles.id, articles.user_id, articles.title FROM users LEFT JOIN articles ON users.id = articles.user_id"
	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*models.User
	var usersMap = make(map[int]*models.User)
	for rows.Next() {
		user := models.User{}
		var articleId sql.NullInt64
		var articleUserId sql.NullInt64
		var articleTitle sql.NullString
		err = rows.Scan(&user.Id, &user.Name, &user.Username, &user.Password, &articleId, &articleUserId, &articleTitle)
		if err != nil {
			return nil, err
		}

		item, isFound := usersMap[user.Id]
		if !isFound {
			item = &user
			usersMap[item.Id] = item
			users = append(users, item)
		}
		if articleId.Valid {
			item.Articles = append(item.Articles, models.Article{
				Id:     int(articleId.Int64),
				Title:  articleTitle.String,
				UserId: int(articleUserId.Int64),
			})
		}
	}

	var usersReturn []models.User
	for _, val := range users {
		usersReturn = append(usersReturn, *val)
	}

	return usersReturn, nil
}
