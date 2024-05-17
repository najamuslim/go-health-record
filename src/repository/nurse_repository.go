package repository

import (
	"context"
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"health-record/model/database"
	"health-record/model/dto"
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type NurseRepository struct {
	db *sql.DB
}

func NewNurseRepository(db *sql.DB) NurseRepositoryInterface {
	return &NurseRepository{db}
}

// CreateNurse inserts a new nurse into the database.

func (repo *NurseRepository) CreateNurse(ctx context.Context, nurse dto.RequestCreateNurse) (string, error) {
	// Generate a new password for the nurse
	password, err := GeneratePassword(12) // For example, generate a 12 character long password
	if err != nil {
			return "", err
	}

	// Hash the generated password before storing it
	hashedPassword, err := HashPassword(password)
	if err != nil {
			return "", err
	}

	// Prepare the SQL query to insert the new nurse with the hashed password
	const query = `INSERT INTO users (user_id, nip, name, role, identity_card_scan_img, password, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING user_id`
	var userId string
	err = repo.db.QueryRowContext(ctx, query, time.Now().UTC().Format("2006-01-02 15:04:05") + strconv.Itoa(randomInt(1, 100000)), nurse.Nip, nurse.Name, "nurse", nurse.IdentityCardScanImg, hashedPassword, time.Now()).Scan(&userId)
	if err != nil {
			return "", err
	}

	// Optionally, you might want to handle sending the password to the nurse or displaying it as needed
	// For security reasons, do not log or display the raw password
	return userId, nil
}


// UpdateNurse updates an existing nurse's information in the database.
func (repo *NurseRepository) UpdateNurse(ctx context.Context, userId string, nurse database.Nurse) error {
	const query = `UPDATE nurses SET nip = $1, name = $2, identity_card_scan_img = $3 WHERE id = $4`
	_, err := repo.db.ExecContext(ctx, query, nurse.Nip, nurse.Name, nurse.IdentityCardScanImg, userId)
	if err != nil {
			return err
	}
	return nil
}

// DeleteNurse removes a nurse from the database.
func (repo *NurseRepository) DeleteNurse(ctx context.Context, userId string) error {
	const query = `DELETE FROM nurses WHERE id = $1`
	_, err := repo.db.ExecContext(ctx, query, userId)
	if err != nil {
			return err
	}
	return nil
}

// GetUsers retrieves nurses from the database based on various filters.
func (repo *NurseRepository) GetNurses(ctx context.Context, filters map[string]interface{}) ([]database.Nurse, error) {
	var users []database.Nurse
	// This is a simplified query that should be adjusted to handle filters correctly.
	query := `SELECT id, nip, name, identity_card_scan_img FROM nurses`
	rows, err := repo.db.QueryContext(ctx, query)
	if err != nil {
			return nil, err
	}
	defer rows.Close()
	
	for rows.Next() {
			var user database.Nurse
			if err := rows.Scan(&user.Id, &user.Nip, &user.Name, &user.IdentityCardScanImg); err != nil {
					return nil, err
			}
			users = append(users, user)
	}

	if err := rows.Err(); err != nil {
			return nil, err
	}
	return users, nil
}

func GeneratePassword(length int) (string, error) {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
			return "", err
	}
	password := base64.URLEncoding.EncodeToString(bytes)
	return password[:length], nil // Trimming the password in case base64 encoding exceeds the desired length
}

// HashPassword hashes the given password using bcrypt.
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
			return "", err
	}
	return string(hashedPassword), nil
}

func (r *NurseRepository) GetNurseByNIP(ctx context.Context, nip int64) (response database.Nurse, err error) {
	err = r.db.QueryRowContext(ctx, "SELECT user_id, name, nip, password FROM users WHERE nip = $1", nip).Scan(&response.Id, &response.Name, &response.Nip, &response.Password)
	if err != nil {
		return
	}
	return
}

func (r *NurseRepository) GetNurseByID(ctx context.Context, userId string) (response database.Nurse, err error) {
	err = r.db.QueryRowContext(ctx, "SELECT id, name, nip, password FROM users WHERE id = $1", userId).Scan(&response.Id, &response.Name, &response.Nip, &response.Password)
	if err != nil {
		return
	}
	return
}