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
	sql := "INSERT INTO users (username, name) values (?, ?)"
	result, err := tx.ExecContext(ctx, sql, user.Username, user.Name)
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
	sql := "UPDATE users SET username = ?, name = ?  WHERE id = ?"
	_, err := tx.ExecContext(ctx, sql, user.Username, user.Name, user.Id)
	if err != nil {
		return user, err
	}

	return user, nil
}
func (implementation *RepositoryUserMysqlImpl) Delete(ctx context.Context, tx *sql.Tx, id int) error {
	sql := "DELETE users WHERE id = ?"
	_, err := tx.ExecContext(ctx, sql, id)
	if err != nil {
		return err
	}

	return nil
}

func (implementation *RepositoryUserMysqlImpl) FindById(ctx context.Context, tx *sql.Tx, id int) (models.User, error) {
	var user models.User
	sql := "SELECT id, username, name FROM users WHERE id = ?"
	rows, err := tx.QueryContext(ctx, sql, id)
	if err != nil {
		return user, err
	}
	defer rows.Close()

	if rows.Next() {
		rows.Scan(&user.Id, &user.Name, &user.Username)
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
