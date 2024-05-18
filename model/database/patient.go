package database

import "time"

type Patient struct {
	IdentityNumber		int64
	PhoneNumber			string
	Name				string
	BirthDate			string
	Gender				string
	IdentityCardScanImg	string
	CreatedAt			time.Time
}