package repository

import (
	"cats-social/model/database"
	"context"
)

type UserRepositoryInterface interface {
	GetUserByEmail(ctx context.Context, email string) (response database.User, err error)
	CreateUser(ctx context.Context, data database.User) (err error)
}
