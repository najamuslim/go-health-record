package usecase

import (
	"context"
	"health-record/model/dto"
	"health-record/src/repository"
)

type PatientUsecase struct {
	iPatientRepository repository.PatientRepositoryInterface	
}

func NewPatientUsecase(
	iPatientRepository repository.PatientRepositoryInterface) PatientUsecaseInterface {
	return &PatientUsecase{iPatientRepository}
}

func (uc *PatientUsecase) RegisterPatient(request dto.RequestCreatePatient) (error) {

	err := uc.iPatientRepository.CreatePatient(context.TODO(), request)
	
	return err
}

func (u *PatientUsecase) GetPatientByIdentityNumber(identityNumber int) (bool, error) {
	_, err := u.iPatientRepository.GetPatientByIdentityNumber(context.TODO(), identityNumber)
	if err != nil {
    return false, err
  }
	return true, nil
}

func (uc *PatientUsecase) GetPatients(request dto.RequestGetPatients) ([]dto.PatientDTO, error) {
	params := dto.RequestGetPatients{
		Limit:    request.Limit,
		Offset:   request.Offset,
		IdentityNumber: request.IdentityNumber,
		Name:     request.Name,
		PhoneNumber: request.PhoneNumber,
		CreatedAt: request.CreatedAt,
	}

	response, err := uc.iPatientRepository.GetPatients(context.TODO(), params)
	
	
	return response, err
}

