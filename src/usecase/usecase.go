package usecase

import (
	"context"
	"health-record/model/database"
	"health-record/model/dto"
)

type AuthUsecaseInterface interface {
	Register(request dto.RequestCreateUser) (token string, err error)
	Login(request dto.RequestAuth) (token string, user database.User, err error)
	GetUserByNIP(nip int64) (exists bool, err error)
}

type NurseUsecaseInterface interface {
	RegisterNurse(request dto.RequestCreateNurse) (string, error)
	UpdateNurse(ctx context.Context, userId string, nurse database.Nurse) error
	DeleteNurse(ctx context.Context, userId string) error
	GetNurses(ctx context.Context, filters map[string]interface{}) ([]database.Nurse, error)
	GetNurseByID(userId string) (bool, error)
	GetNurseByNIP(nip int64) (bool, error)
}
