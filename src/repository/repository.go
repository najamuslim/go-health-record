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
	GetNurses(ctx context.Context, filters map[string]interface{}) ([]database.Nurse, error)
	UpdateNurse(ctx context.Context, userId string, nurse database.Nurse) error
	DeleteNurse(ctx context.Context, userId string) error
	GetNurseByNIP(ctx context.Context, nip int64) (response database.Nurse, err error)
	GetNurseByID(ctx context.Context, userId string) (response database.Nurse, err error)
}
