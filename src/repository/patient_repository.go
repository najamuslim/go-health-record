package repository

import (
	"context"
	"database/sql"
	"health-record/model/database"
	"health-record/model/dto"
	"time"
)

type PatientRepository struct {
	db *sql.DB
}

func NewPatientRepository(db *sql.DB) PatientRepositoryInterface {
	return &PatientRepository{db}
}

func (r *PatientRepository) CreatePatient(ctx context.Context, request dto.RequestCreatePatient) (err error) {
	query := `
	INSERT INTO Patients (
		identity_number,
		phone_number,
		name,
		birth_date,
		gender,
		identity_card_scan_img,
		created_at)
	VALUES ($1, $2, $3, $4, $5, $6, $7)`

	_, err = r.db.ExecContext(
		ctx,
		query,
		request.IdentityNumber,
		request.PhoneNumber,
		request.Name,
		request.BirthDate,
		request.Gender,
		request.IdentityCardScanImg,
		time.Now(),
	)

	return err
}

func (r *PatientRepository) GetPatientByIdentityNumber(ctx context.Context, identityNumber int) (response database.Patient, err error) {
	err = r.db.QueryRowContext(ctx, "SELECT identity_number FROM patients WHERE identity_number = $1", identityNumber).Scan(&response.IdentityNumber)
	if err != nil {
		
		return database.Patient{}, err
	}
	return response, nil
}
