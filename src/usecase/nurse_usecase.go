package usecase

import (
	"context"
	"errors"
	"health-record/model/database"
	"health-record/model/dto"
	"health-record/src/repository"
)

type NurseUsecase struct {
	iNurseRepository repository.NurseRepositoryInterface
}

func NewNurseUsecase(
	iNurseRepository repository.NurseRepositoryInterface) NurseUsecaseInterface {
	return &NurseUsecase{iNurseRepository}
}

func (uc *NurseUsecase) RegisterNurse(request dto.RequestCreateNurse) (string, error) {
	// Check if the nurse NIP already exists in the database
	existingNurse, err := uc.iNurseRepository.GetNurseByNIP(context.TODO(), request.Nip)
	if err == nil && existingNurse.Id != "" {
			return "", errors.New("a nurse with this NIP already exists")
	}

	// If the nurse does not exist, proceed to create a new nurse
	userId, err := uc.iNurseRepository.CreateNurse(context.TODO(), request)
	if err != nil {
			return "", err
	}
	return userId, nil
}

// UpdateNurse handles the updating of an existing nurse's information.
func (uc *NurseUsecase) UpdateNurse(ctx context.Context, userId string, nurse database.Nurse) error {
	// Ensure the nurse exists before attempting to update
	_, err := uc.iNurseRepository.GetNurseByID(ctx, userId)
	if err != nil {
			return errors.New("nurse not found")
	}

	// Proceed with updating the nurse
	return uc.iNurseRepository.UpdateNurse(ctx, userId, nurse)
}

// DeleteNurse handles the deletion of a nurse.
func (uc *NurseUsecase) DeleteNurse(ctx context.Context, userId string) error {
	return uc.iNurseRepository.DeleteNurse(ctx, userId)
}

// GetNurses retrieves nurses based on optional filters.
func (uc *NurseUsecase) GetNurses(ctx context.Context, filters map[string]interface{}) ([]database.Nurse, error) {
	return uc.iNurseRepository.GetNurses(ctx, filters)
}

func (u *NurseUsecase) GetNurseByNIP(nip int64) (bool, error) {
	_, err := u.iNurseRepository.GetNurseByNIP(context.TODO(), nip)
	if err != nil {
    return false, err
  }
	return true, nil
}

func (u *NurseUsecase) GetNurseByID(id string) (bool, error) {
	_, err := u.iNurseRepository.GetNurseByID(context.TODO(), id)
  if err != nil {
    return false, err
  }
  return true, nil
}