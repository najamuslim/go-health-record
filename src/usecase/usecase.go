package usecase

import (
	"cats-social/model/database"
	"cats-social/model/dto"
)

type AuthUsecaseInterface interface {
	Register(request dto.RequestCreateUser) (token string, err error)
	Login(request dto.RequestAuth) (token string, user database.User, err error)
	GetUserByEmail(email string) (exists bool, err error)
}
