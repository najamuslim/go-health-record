package repository

import (
	"context"
	"health-record/model/database"
	"health-record/model/dto"
)

type UserRepositoryInterface interface {
	GetUserByNIP(ctx context.Context, nip int64) (response database.User, err error)
	CreateUser(ctx context.Context, data database.User) (err error)
}

type NurseRepositoryInterface interface {
	CreateNurse(ctx context.Context, data dto.RequestCreateNurse) (string, error)
	GetUsers(ctx context.Context, param dto.RequestGetUser) ([]dto.UserDTO, error)
	UpdateNurse(ctx context.Context, userId string, nurse database.User) error
	DeleteNurse(ctx context.Context, userId string) int
	GetNurseByNIP(ctx context.Context, nip int64) (response database.User, err error)
	GetNurseByID(ctx context.Context, userId string) (response database.User, err error)
}
