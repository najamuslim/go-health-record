package usecase

import (
	"context"
	"health-record/model/database"
	"health-record/model/dto"
)

type AuthUsecaseInterface interface {
	Register(request dto.RequestCreateUser) (token string, userId string, err error)
	Login(request dto.RequestAuth) (token string, user database.User, err error)
	GetUserByNIP(nip int64) (exists bool, err error)
}

type NurseUsecaseInterface interface {
	RegisterNurse(request dto.RequestCreateNurse) (string, error)
	GetUsers(request dto.RequestGetUser) ([]dto.UserDTO, error)
	UpdateNurse(ctx context.Context, userId string, nurse database.User) error
	DeleteNurse(userId string) int
	GetNurseByID(userId string) (bool, error)
	GetNurseByNIP(nip int64) (bool, error)
}
