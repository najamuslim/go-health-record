package database

import "time"

type Nurse struct {
	Id			string
	Nip		string
	Password	string
	Name		string
	IdentityCardScanImg string
	CreatedAt 	time.Time
	UpdatedAt 	time.Time
}