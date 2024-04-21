// This file contains the repository implementation layer.
package repository

import (
	"context"
	"database/sql"
	"errors"
	_ "github.com/lib/pq"
)

const TableUser = "users"

type Repository struct {
	Conn *sql.Conn
}

type NewRepositoryOptions struct {
	Conn *sql.Conn
}

func (r Repository) GetById(ctx context.Context, input GetUserByIdInput) (*GetUserByIdOutput, error) {

	query := `SELECT id, phone_number, full_name, login_success_count, created_at, updated_at FROM users WHERE id = $1;`

	queryStatement, err := r.Conn.PrepareContext(ctx, query)

	if err != nil {
		return nil, err
	}

	result := GetUserByIdOutput{}

	err = queryStatement.QueryRowContext(ctx, input.Id).
		Scan(&result.Id, &result.PhoneNumber, &result.FullName, &result.LoginSuccessCount, &result.CreatedAt, &result.UpdatedAt)

	if err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return &result, nil
}

func (r Repository) Update(ctx context.Context, input UpdateUserInput) (*UpdateUserOutput, error) {
	query := `UPDATE users SET phone_number = $1, full_name = $2, login_success_count = $3 WHERE id = $4;`

	queryStatement, err := r.Conn.PrepareContext(ctx, query)

	if err != nil {
		return nil, err
	}

	_, err = queryStatement.ExecContext(ctx, input.PhoneNumber, input.FullName, input.LoginSuccessCount, input.Id)

	if err != nil {
		return nil, err
	}

	output := &UpdateUserOutput{
		IsSuccessUpdate: true,
	}

	return output, nil
}

func (r Repository) Insert(ctx context.Context, input InsertUserInput) (output *InsertUserOutput, err error) {

	var lastInsertId int64

	query := `INSERT INTO users (phone_number, full_name, password) VALUES ($1, $2, $3) RETURNING id;`

	queryStatement, err := r.Conn.PrepareContext(ctx, query)

	if err != nil {
		return nil, err
	}

	err = queryStatement.QueryRowContext(ctx, input.PhoneNumber, input.FullName, input.Password).Scan(&lastInsertId)

	if err != nil {
		return nil, err
	}

	output = &InsertUserOutput{
		Id: lastInsertId,
	}

	return output, nil
}

func (r Repository) GetByPhoneNumberIncludePassword(ctx context.Context, input GetUserByPhoneNumberInput) (*GetUserByPhoneNumberOutput, error) {

	query := `SELECT id, phone_number, full_name, password, login_success_count, created_at, updated_at FROM users WHERE phone_number = $1;`

	queryStatement, err := r.Conn.PrepareContext(ctx, query)

	if err != nil {
		return nil, err
	}

	result := GetUserByPhoneNumberOutput{}

	err = queryStatement.QueryRowContext(ctx, input.PhoneNumber).
		Scan(&result.Id, &result.PhoneNumber, &result.FullName, &result.Password, &result.LoginSuccessCount, &result.CreatedAt, &result.UpdatedAt)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return &result, nil
}

func NewRepository(opts NewRepositoryOptions) *Repository {

	return &Repository{
		Conn: opts.Conn,
	}
}
