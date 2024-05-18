package repository

import (
	"context"
	"database/sql"
	"fmt"
	"health-record/model/database"
	"health-record/model/dto"
	"strconv"
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

func (ur *PatientRepository) GetPatients(ctx context.Context, param dto.RequestGetPatients) ([]dto.PatientDTO, error) {
	query := "SELECT identity_number, name, phone_number, birth_date, gender, created_at FROM patients WHERE 1=1"

	var args []interface{}

	if param.IdentityNumber != nil {
		query += " AND identity_number = $" + strconv.Itoa(len(args)+1)
		args = append(args, strconv.Itoa(*param.IdentityNumber))
	}

	if param.Name != nil {
		query += " AND LOWER(name) LIKE LOWER($" + strconv.Itoa(len(args)+1) + ")"
		args = append(args, "%"+*param.Name+"%")
	}

	if param.PhoneNumber != nil {
		query += " AND phone_number LIKE $" + strconv.Itoa(len(args)+1)
		args = append(args, "%"+strconv.Itoa(*param.PhoneNumber)+"%")
	}

	if param.CreatedAt == "asc" || param.CreatedAt == "desc" {
		query += fmt.Sprintf(" ORDER BY created_at %s", param.CreatedAt)
	}

	query += fmt.Sprintf(" LIMIT %d OFFSET %d", param.Limit, param.Offset)

	rows, err := ur.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var patients []dto.PatientDTO
	for rows.Next() {
		var ptn dto.PatientDTO
		if err := rows.Scan(
			&ptn.IdentityNumber,
			&ptn.Name,
			&ptn.PhoneNumber,
			&ptn.BirthDate,
			&ptn.Gender,
			&ptn.CreatedAt); err != nil {
			return nil, err
		}

		patients = append(patients, ptn)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return patients, nil
}