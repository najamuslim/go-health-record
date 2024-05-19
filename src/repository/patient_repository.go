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

func (ur *PatientRepository) GetRecords(ctx context.Context, request dto.RequestGetRecord) ([]dto.MedicalRecords, error) {
	var records []dto.MedicalRecords

	query := "SELECT medical_records.symptoms, medical_records.medications, medical_records.created_at, users.name AS created_by_name, users.nip AS created_by_nip, users.user_id AS created_by_user_id, patients.phone_number, patients.identity_number, patients.name, patients.birth_date, patients.gender, patients.identity_card_scan_img FROM medical_records LEFT JOIN patients ON medical_records.identity_number = patients.identity_number LEFT JOIN users ON medical_records.created_by = users.user_id WHERE 1=1"

	var params []interface{}

	if request.IdentityDetail.IdentityNumber != 0 {
		query += " AND medical_records.identity_number = CAST(? AS VARCHAR)"
		params = append(params, request.IdentityDetail.IdentityNumber)
	}

	if request.CreatedBy.UserID != "" {
		query += " AND users.id = ?"
		params = append(params, request.CreatedBy.UserID)
	}

	if request.CreatedBy.Nip != "" {
		query += " AND users.nip = ?"
		params = append(params, request.CreatedBy.Nip)
	}

	if request.CreatedAt != "" {
		query += " ORDER BY medical_records.created_at " + request.CreatedAt
	}

	query += fmt.Sprintf(" LIMIT %d OFFSET %d", request.Limit, request.Offset)

	rows, err := ur.db.QueryContext(ctx, query, params...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var ptn dto.MedicalRecords
		if err := rows.Scan(
			&ptn.Symptoms,
			&ptn.Medications,
			&ptn.CreatedAt,
			&ptn.CreatedBy.Name,
			&ptn.CreatedBy.Nip,
			&ptn.CreatedBy.UserId,
			&ptn.IdentityDetail.IdentityNumber,
			&ptn.IdentityDetail.PhoneNumber,
			&ptn.IdentityDetail.Name,
			&ptn.IdentityDetail.BirthDate,
			&ptn.IdentityDetail.Gender,
			&ptn.IdentityDetail.IdentityCardScanImg,
		); err != nil {
			return nil, err
		}

		records = append(records, ptn)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return records, nil
}

func (ur *PatientRepository) CreateRecord(ctx context.Context, data database.MedicalRecord) (err error) {
	query := `
	INSERT INTO medical_records (
		identity_number,
		symptoms,
		medications,
		created_by,
		created_at)
	VALUES ($1, $2, $3, $4, $5)`

	_, err = ur.db.ExecContext(
		ctx,
		query,
		data.IdentityNumber,
		data.Symptoms,
		data.Medications,
		data.CreatedBy,
		data.CreatedAt,
	)

	return err
}
