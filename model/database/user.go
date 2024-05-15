package database

import "time"

type User struct {
	Id			int
	Nip		string
	Password	string
	Name		string
	CreatedAt 	time.Time
	UpdatedAt 	time.Time
}

type Nurse struct {
	Id			int
	Nip		string
	Password	string
	Name		string
	IdentityCardScanImg string
	CreatedAt 	time.Time
	UpdatedAt 	time.Time
}