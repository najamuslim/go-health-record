package repository

import (
	"context"
	"database/sql"
	"health-record/model/dto"
)

type PatientRepository struct {
	db *sql.DB
}

func NewPatientRepository(db *sql.DB) PatientRepositoryInterface {
	return &PatientRepository{db}
}

func (r *PatientRepository) CreatePatient(ctx context.Context, request dto.RequestCreatePatient) (err error) {
	query := `
	INSERT INTO Patients (email, name, password, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id`

	_, err = r.db.ExecContext(
		ctx,
		query,
		request.IdentityNumber,
		request.Name,
		request.BirthDate,
		request.Gender,
		request.IdentityCardScanImg,
	)

	return err
}
