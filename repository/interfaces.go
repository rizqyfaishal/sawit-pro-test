// This file contains the interfaces for the repository layer.
// The repository layer is responsible for interacting with the database.
// For testing purpose we will generate mock implementations of these
// interfaces using mockgen. See the Makefile for more information.
package repository

import "context"

type UserRepositoryInterface interface {
	GetById(ctx context.Context, input GetUserByIdInput) (*GetUserByIdOutput, error)
	GetByPhoneNumberIncludePassword(ctx context.Context, input GetUserByPhoneNumberInput) (*GetUserByPhoneNumberOutput, error)
	Update(ctx context.Context, input UpdateUserInput) (*UpdateUserOutput, error)
	Insert(ctx context.Context, input InsertUserInput) (*InsertUserOutput, error)
}
