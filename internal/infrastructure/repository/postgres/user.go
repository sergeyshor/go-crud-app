package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"go-crud-app/internal/dto"
	"go-crud-app/internal/entity"
)

type UserRepo struct {
	*sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db}
}

func (ur *UserRepo) Create(ctx context.Context, dto dto.CreateUserDTO) error {
	query := "INSERT INTO users(name, email, password) VALUES($1, $2, $3) RETURNING id, name, email, password, created_at"
	result, err := ur.ExecContext(ctx, query, dto.Name, dto.Email, dto.Password)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows != 1 {
		return fmt.Errorf("expected to affect 1 row, affected %d", rows)
	}
	return nil
}

func (ur *UserRepo) GetByID(ctx context.Context, id int64) (*entity.User, error) {
	row := ur.QueryRowContext(ctx, "SELECT * FROM users WHERE id=$1", id)

	user := &entity.User{}
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (ur *UserRepo) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	row := ur.QueryRowContext(ctx, "SELECT * FROM users WHERE email=$1", email)

	user := &entity.User{}
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (ur *UserRepo) DeleteByID(ctx context.Context, id int64) error {
	result, err := ur.ExecContext(ctx, "DELETE FROM users WHERE id=$1", id)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows != 1 {
		return fmt.Errorf("expected to affect 1 row, affected %d", rows)
	}
	return nil
}
