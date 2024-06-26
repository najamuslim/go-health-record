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
	UpdateNurse(ctx context.Context, userId string, nurse dto.RequestUpdateNurse) int
	DeleteNurse(ctx context.Context, userId string) int
	AddAccess(ctx context.Context, userId string, password dto.RequestAddAccess) int
	GetNurseByNIP(ctx context.Context, nip int64) (response database.User, err error)
	GetNurseByID(ctx context.Context, userId string) (response database.User, err error)
}

type PatientRepositoryInterface interface {
	CreatePatient(ctx context.Context, request dto.RequestCreatePatient) (err error)
	GetPatientByIdentityNumber(ctx context.Context, identityNumber int) (database.Patient, error)
	GetPatients(ctx context.Context, params dto.RequestGetPatients) ([]dto.PatientDTO, error)
	GetRecords(ctx context.Context, request dto.RequestGetRecord) ([]dto.MedicalRecords, error)
	CreateRecord(ctx context.Context, data database.MedicalRecord) (err error)
}
