// This file contains types that are used in the repository layer.
package repository

import "time"

// Query Struct

type GetUserByIdInput struct {
	Id int64
}

type GetUserByPhoneNumberInput struct {
	PhoneNumber string
}

type InsertUserInput struct {
	PhoneNumber string
	FullName    string
	Password    string
}

type UpdateUserInput struct {
	Id int64

	PhoneNumber       string
	FullName          string
	LoginSuccessCount int64
}

// Output struct

type GetUserByIdOutput struct {
	Id                int64
	PhoneNumber       string
	FullName          string
	LoginSuccessCount int64
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

type GetUserByPhoneNumberOutput struct {
	Id                int64
	PhoneNumber       string
	FullName          string
	LoginSuccessCount int64
	Password          string
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

type InsertUserOutput struct {
	Id int64
}

type UpdateUserOutput struct {
	IsSuccessUpdate bool
}
