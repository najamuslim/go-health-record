package usecase

import (
	"health-record/model/database"
	"health-record/model/dto"
)

type AuthUsecaseInterface interface {
	Register(request dto.RequestCreateUser) (token string, userId string, err error)
	Login(request dto.RequestAuth) (token string, user database.User, err error)
	GetUserByNIP(nip int64) (exists bool, err error)
	LoginNurse(request dto.RequestAuth) (token string, user database.User, err error)
}

type NurseUsecaseInterface interface {
	RegisterNurse(request dto.RequestCreateNurse) (string, error)
	GetUsers(request dto.RequestGetUser) ([]dto.UserDTO, error)
	UpdateNurse(userId string, nurse dto.RequestUpdateNurse) int
	DeleteNurse(userId string) int
	AddAccess(userId string, password dto.RequestAddAccess) int
	GetNurseByID(userId string) (database.User, error)
	GetNurseByNIP(nip int64) (bool, error)
}

type PatientUsecaseInterface interface {
	RegisterPatient(identityNumber int) (error)
}

